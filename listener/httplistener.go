package listener

import (
	"encoding/json"
	"errors"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"io/ioutil"
	"net/http"
	"net/mail"
	"runtime/pprof"
	"strings"
)

// Listener for handling and dispatching incoming http requests
type HTTPListener struct {
	Engine        *engine.HeliaEngine
	Configuration *listenerConfig
}

// Loads the listener configuration and initializes the listener
func (l *HTTPListener) Initialize() {
	// load listener configuration
	config, err := loadConfiguration()

	if err != nil {
		panic("unable to load listener configuration")
	}

	l.Configuration = &config
}

// Sets the CORS policy (currently all, needs securing)
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Returns the port the game engine is configured to listen on
func (l *HTTPListener) GetPort() int {
	return l.Configuration.Port
}

// Responds to HTTP pings
func (l *HTTPListener) HandlePing(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// write pingback
	fmt.Fprintf(w, "alive!")
}

// Handles a user registering
func (l *HTTPListener) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// get services
	userSvc := sql.GetUserService()
	startSvc := sql.GetStartService()
	containerSvc := sql.GetContainerService()
	factionSvc := sql.GetFactionService()

	// read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// parse model
	m := models.APIRegisterModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, "parsemodel: "+err.Error(), 500)
		return
	}

	// trim strings
	m.CharacterName = strings.Trim(m.CharacterName, " ")
	m.EmailAddress = strings.Trim(m.EmailAddress, " ")

	// validation
	if len(m.CharacterName) == 0 {
		http.Error(w, "Character Name must not be empty.", 500)
		return
	}

	if len(m.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters.", 500)
		return
	}

	if m.Password != m.ConfirmPassword {
		http.Error(w, "Passwords must match.", 500)
		return
	}

	email, v := isValidEmailAddress(m.EmailAddress)

	if !v {
		http.Error(w, "Email address is invalid.", 500)
		return
	}

	m.EmailAddress = email

	// validate start
	startID := m.StartID

	start, err := startSvc.GetStartByID(startID)

	if err != nil {
		http.Error(w, "getstart: "+err.Error(), 500)
		return
	}

	if !start.Available {
		http.Error(w, "You must select an available start.", 500)
		return
	}

	// create escrow container for user
	ec, err := containerSvc.NewContainer(sql.Container{
		Meta: sql.Meta{},
	})

	if err != nil {
		http.Error(w, "createscrowcontainer: "+err.Error(), 500)
		return
	}

	// get user's ip
	ip := readUserIP(r)

	// create user
	u, err := userSvc.NewUser(m.CharacterName, m.Password, startID, ec.ID, start.FactionID, m.EmailAddress, ip)

	if err != nil {
		http.Error(w, "createuser: "+err.Error(), 500)
		return
	}

	// create their noob ship
	_, err = engine.CreateNoobShipForPlayer(start, u.ID)

	if err != nil {
		http.Error(w, "createnoobshipforplayer: "+err.Error(), 500)
		return
	}

	// load faction
	f, err := factionSvc.GetFactionByID(start.FactionID)

	if err != nil {
		http.Error(w, "loadfaction: "+err.Error(), 500)
		return
	}

	// build initial reputation sheet
	u.ReputationSheet.FactionEntries = make(map[string]sql.PlayerReputationSheetFactionEntry)

	for _, r := range f.ReputationSheet.Entries {
		u.ReputationSheet.FactionEntries[r.TargetFactionID.String()] = sql.PlayerReputationSheetFactionEntry{
			FactionID:        r.TargetFactionID,
			StandingValue:    r.StandingValue,
			AreOpenlyHostile: r.AreOpenlyHostile,
		}
	}

	// save initial reputation sheet
	err = userSvc.SaveReputationSheet(u.ID, u.ReputationSheet)

	if err != nil {
		http.Error(w, "saveinitialrepsheet: "+err.Error(), 500)
		return
	}

	// success!
	w.WriteHeader(200)
}

// Handles a user signing in
func (l *HTTPListener) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// get services
	userSvc := sql.GetUserService()
	sessionSvc := sql.GetSessionService()

	// read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// parse model
	m := models.APILoginModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// attempt login
	res := models.APILoginResponseModel{}
	user, err := userSvc.GetUserByLogin(m.EmailAddress, m.Password)

	if err == nil {
		// verify not an NPC account
		if user.IsNPC {
			err = errors.New("logins not allowed for NPC accounts")
		}

		// verify not banned
		if user.Banned {
			err = errors.New("you have been banned")
		}

		// log dev login attempts
		if user.IsDev {
			shared.TeeLog(fmt.Sprintf("A developer is trying to log in! %v", *user.EmailAddress))
		}
	}

	if err != nil {
		res.Success = false
		res.Message = err.Error()
	} else {
		res.UID = user.ID.String()

		// delete old session
		err := sessionSvc.DeleteSession(user.ID)

		if err != nil {
			res.Success = false
			res.Message = err.Error()
		} else {
			// get user's ip
			ip := readUserIP(r)

			// create session
			s, err := sessionSvc.NewSession(user.ID, ip)

			if err != nil {
				res.Success = false
				res.Message = err.Error()
			} else {
				// store sessionid in result
				res.SessionID = (&s.ID).String()
				res.Success = true
			}
		}
	}

	// log result
	if res.Success {
		shared.TeeLog(fmt.Sprintf("Login success: %v | %v >> %v", m.EmailAddress, readUserIP(r), res.SessionID))
	} else {
		shared.TeeLog(fmt.Sprintf("Login failed: %v | %v", m.EmailAddress, readUserIP(r)))
	}

	// return result
	o, _ := json.Marshal(res)
	w.Write(o)
}

// Shuts down the server after saving the game state
func (l *HTTPListener) HandleShutdown(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// get auth token
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		return
	}

	key := keys[0]

	// validate auth token
	if l.Configuration.ShutdownToken == key {
		// initiate shutdown
		l.Engine.Shutdown()

		// stop profiling
		pprof.StopCPUProfile()

		// write response
		fmt.Fprintln(w, "shutdown complete")
	}
}

// This only exists to allow a user to accept the self-signed cert
func (l *HTTPListener) HandleAcceptCert(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// write response
	fmt.Fprintln(w, "you've accepted the high quality, self-signed, certificate - woohoo!")
}

// Determines whether a given string is a valid email address
func isValidEmailAddress(e string) (string, bool) {
	addr, err := mail.ParseAddress(e)

	if err != nil {
		return "", false
	}

	return addr.Address, true
}

// Determines a user's real ip to as best as possible from within a container
func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	return IPAddress
}
