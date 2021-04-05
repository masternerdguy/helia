package listener

import (
	"encoding/json"
	"helia/engine"
	"helia/listener/models"
	"helia/sql"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

// Listener for handling and dispatching incoming http requests
type HTTPListener struct {
	Engine *engine.HeliaEngine
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Handles a user registering
func (l *HTTPListener) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// enable cors
	enableCors(&w)

	// get services
	userSvc := sql.GetUserService()
	startSvc := sql.GetStartService()
	containerSvc := sql.GetContainerService()

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

	// validation
	if len(m.Username) == 0 {
		http.Error(w, "Username must not be empty.", 500)
		return
	}

	if len(m.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters.", 500)
		return
	}

	// get start - todo: make this player selectable from a list of available starts
	startID, err := uuid.Parse("49f06e8929fb4a02a0344b5d0702adac")

	if err != nil {
		http.Error(w, "caststartid: "+err.Error(), 500)
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

	// create user
	u, err := userSvc.NewUser(m.Username, m.Password, startID, ec.ID)

	if err != nil {
		http.Error(w, "createuser: "+err.Error(), 500)
		return
	}

	start, err := startSvc.GetStartByID(startID)

	if err != nil {
		http.Error(w, "getstart: "+err.Error(), 500)
		return
	}

	// create their noob ship
	_, err = engine.CreateNoobShipForPlayer(start, u.ID, true)

	if err != nil {
		http.Error(w, "createnoobshipforplayer: "+err.Error(), 500)
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
	user, err := userSvc.GetUserByLogin(m.Username, m.Password)

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
			// create session
			s, err := sessionSvc.NewSession(user.ID)

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

	// load listener configuration
	config, err := loadConfiguration()
	if err != nil {
		return
	}

	// validate auth token
	if config.ShutdownToken == key {
		// initiate shutdown
		l.Engine.Shutdown()
	}
}
