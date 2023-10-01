package listener

import (
	"encoding/json"
	"errors"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/shared"
	"helia/sql"
	"io"
	"math/rand"
	"net/http"
	"net/mail"
	"runtime/pprof"
	"strings"

	"github.com/google/uuid"
)

// Lookup table of braille cells
var BRAILLE_CELLS = [...]string{
	"⠮", "⠐", "⠼", "⠫", "⠩", "⠯", "⠄", "⠷",
	"⠾", "⠡", "⠬", "⠠", "⠤", "⠨", "⠌", "⠴",
	"⠂", "⠆", "⠒", "⠲", "⠢", "⠖", "⠶", "⠦",
	"⠔", "⠱", "⠰", "⠣", "⠿", "⠜", "⠹", "⠈",
	"⠁", "⠃", "⠉", "⠙", "⠑", "⠋", "⠛", "⠓",
	"⠊", "⠚", "⠅", "⠇", "⠍", "⠝", "⠕", "⠏",
	"⠟", "⠗", "⠎", "⠞", "⠥", "⠧", "⠺", "⠭",
	"⠽", "⠵", "⠪", "⠳", "⠻", "⠘", "⠸",
}

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

	// get health
	ph, hm := shared.GetServerHealth()

	// get 2 random braille cells because it looks cool
	q1 := randomBraille()
	q2 := randomBraille()

	// build pingback
	pb := fmt.Sprintf("%v %v %v %v", q1, ph, q2, hm)

	// write pingback
	fmt.Fprint(w, pb)
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
	b, err := io.ReadAll(r.Body)
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
	err = userSvc.UpdateReputationSheet(u.ID, u.ReputationSheet)

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
	b, err := io.ReadAll(r.Body)
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

// Password reset requests for existing users (forgot stage)
func (l *HTTPListener) HandleForgot(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// shared error message to hide details
	const se = "Something went wrong requesting a password reset."

	// get services
	userSvc := sql.GetUserService()

	// get user's ip
	ip := readUserIP(r)

	// read body
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, se, 500)
		return
	}

	// parse model
	m := models.APIForgotModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, se, 500)
		return
	}

	// get user id for email address
	res := models.APIForgotResponseModel{}
	uid, err := userSvc.GetUserIdByEmailAddress(m.EmailAddress)

	if err != nil || uid == nil {
		http.Error(w, se, 500)
		return
	}

	// get user
	user, err := userSvc.GetUserByID(*uid)

	if err != nil {
		res.Success = false
		res.Message = se
	} else {
		// log dev reset attempts
		if user.IsDev {
			shared.TeeLog(fmt.Sprintf("A developer is trying to request a password reset token! %v << %v", *user.EmailAddress, ip))
		} else {
			// log regular reset attempts
			shared.TeeLog(fmt.Sprintf("A user is trying to request a password reset token: %v << %v", *user.EmailAddress, ip))
		}

		// generate reset token
		token := uuid.New()

		// save reset token
		err := userSvc.UpdateResetToken(user.ID, &token)

		if err != nil {
			res.Success = false
			res.Message = se
		} else {
			// generate email body
			b := shared.FillPasswordResetTokenEmailBody(r.Referer(), user.ID, token)

			// send email
			err := shared.SendEmail(*user.EmailAddress, "Password Reset", b)
			res.Success = err == nil
		}
	}

	// log success
	if res.Success {
		shared.TeeLog(fmt.Sprintf("Reset token sent for %v << %v", m.EmailAddress, ip))
		res.Message = "A link to reset your password has been emailed to you."
	} else {
		res.Message = se
	}

	// return result
	o, _ := json.Marshal(res)
	w.Write(o)
}

// Password reset requests for existing users (reset stage)
func (l *HTTPListener) HandleReset(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// shared error message to hide details
	const se = "Something went wrong resetting your password."

	// get services
	userSvc := sql.GetUserService()

	// get user's ip
	ip := readUserIP(r)

	// read body
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, se, 500)
		return
	}

	// parse model
	m := models.APIResetModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, se, 500)
		return
	}

	// get user id for email address
	res := models.APIResetResponseModel{
		Success: false,
		Message: se,
	}

	// get user
	user, err := userSvc.GetUserByID(m.UserID)

	if err != nil {
		// not found
		res.Success = false
		res.Message = se
	} else {
		// log dev reset attempts
		if user.IsDev {
			shared.TeeLog(fmt.Sprintf("A developer is trying to reset their password! %v << %v", *user.EmailAddress, ip))
		} else {
			// log regular reset attempts
			shared.TeeLog(fmt.Sprintf("A user is trying to reset their password: %v << %v", *user.EmailAddress, ip))
		}

		// validate reset token
		t, _ := userSvc.GetPasswordResetToken(m.UserID)

		// for empty uuid check
		var emptyUUID uuid.UUID

		if t == nil || *t == emptyUUID || *t != m.Token || m.Token == emptyUUID {
			// validation failed
			res.Success = false
			res.Message = se
		} else {
			valid := true

			// validate password
			if len(m.Password) < 8 {
				res.Message = "Password must be at least 8 characters."
				valid = false
			}

			if m.Password != m.ConfirmPassword {
				res.Message = "Passwords must match."
				valid = false
			}

			if valid {
				// do reset
				err = userSvc.UpdatePassword(m.UserID, m.Token, m.Password)

				if err != nil {
					// reset failed
					res.Success = false
					res.Message = se
				} else {
					// mark as success
					res.Success = true

					// send notification email
					b := shared.FillPasswordResetSuccessEmailBody(ip)
					err := shared.SendEmail(*user.EmailAddress, "Password Reset Success", b)

					if err != nil {
						// log failure to send email
						shared.TeeLog(
							fmt.Sprintf("Failed to send password reset success email for %v: %v", *user.EmailAddress, err),
						)
					}
				}
			}
		}
	}

	// log success
	if res.Success {
		shared.TeeLog(fmt.Sprintf("Password reset for %v << %v", *user.EmailAddress, ip))
		res.Message = "You can now log in using your new password."
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

// Helper function to return a random 6-dot braille cell
func randomBraille() string {
	// get random cell
	i := rand.Intn(len(BRAILLE_CELLS))
	c := BRAILLE_CELLS[i]

	// return result
	return c
}
