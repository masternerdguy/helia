package listener

import (
	"encoding/json"
	"helia/listener/models"
	"helia/sql"
	"io/ioutil"
	"net/http"
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
		http.Error(w, err.Error(), 500)
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
	_, err = userSvc.NewUser(m.Username, m.Password)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//success!
	w.WriteHeader(200)
}
