package listener

import (
	"encoding/json"
	"fmt"
	"helia/listener/models"
	"helia/sql"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//HandleRegister Handles a user registering
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	//enable cors
	enableCors(&w)

	//get services
	userSvc := sql.GetUserService()
	shipSvc := sql.GetShipService()

	//read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//parse model
	m := models.APIRegisterModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, "parsemodel: "+err.Error(), 500)
		return
	}

	//validation
	if len(m.Username) == 0 {
		http.Error(w, "Username must not be empty.", 500)
		return
	}

	if len(m.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters.", 500)
		return
	}

	//create user
	u, err := userSvc.NewUser(m.Username, m.Password)

	if err != nil {
		http.Error(w, "createuser: "+err.Error(), 500)
		return
	}

	//create starter ship
	systemID, err := uuid.Parse("1d4e0a339f674f248b7b1af4d5aa2ef1")

	if err != nil {
		http.Error(w, "castsystemid: "+err.Error(), 500)
		return
	}

	t := sql.Ship{
		SystemID: systemID,
		UserID:   u.ID,
		ShipName: fmt.Sprintf("%s's Starter Ship", m.Username),
		Texture:  "Mass Testing Brick",
		Theta:    0,
	}

	starterShip, err := shipSvc.NewShip(t)

	if err != nil {
		http.Error(w, "createstartership: "+err.Error(), 500)
		return
	}

	//put user in starter ship
	err = userSvc.SetCurrentShipID(u.ID, &starterShip.ID)

	if err != nil {
		http.Error(w, "putuserinstartership: "+err.Error(), 500)
		return
	}

	//success!
	w.WriteHeader(200)
}

//HandleLogin Handles a user signing in
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//enable cors
	enableCors(&w)

	//get services
	userSvc := sql.GetUserService()
	sessionSvc := sql.GetSessionService()

	//read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//parse model
	m := models.APILoginModel{}
	err = json.Unmarshal(b, &m)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//attempt login
	res := models.APILoginResponseModel{}
	user, err := userSvc.GetUserByLogin(m.Username, m.Password)

	if err != nil {
		res.Success = false
		res.Message = err.Error()
	} else {
		res.UID = user.ID.String()

		//delete old session
		err := sessionSvc.DeleteSession(user.ID)

		if err != nil {
			res.Success = false
			res.Message = err.Error()
		} else {
			//create session
			s, err := sessionSvc.NewSession(user.ID)

			if err != nil {
				res.Success = false
				res.Message = err.Error()
			} else {
				//store sessionid in result
				res.SessionID = (&s.ID).String()
				res.Success = true
			}
		}
	}

	//return result
	o, _ := json.Marshal(res)
	w.Write(o)
}
