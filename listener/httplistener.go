package listener

import (
	"encoding/json"
	"fmt"
	"helia/engine"
	"helia/listener/models"
	"helia/physics"
	"helia/sql"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

//HTTPListener Listener for handling and dispatching incoming http requests
type HTTPListener struct {
	Engine *engine.HeliaEngine
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//HandleRegister Handles a user registering
func (l *HTTPListener) HandleRegister(w http.ResponseWriter, r *http.Request) {
	//enable cors
	enableCors(&w)

	//get services
	userSvc := sql.GetUserService()
	shipSvc := sql.GetShipService()
	shipTmpSvc := sql.GetShipTemplateService()
	startSvc := sql.GetStartService()
	itemSvc := sql.GetItemService()

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

	//get start - todo: make this player selectable from a list of available starts
	startID, err := uuid.Parse("49f06e8929fb4a02a0344b5d0702adac")

	if err != nil {
		http.Error(w, "caststartid: "+err.Error(), 500)
		return
	}

	//create user
	u, err := userSvc.NewUser(m.Username, m.Password, startID)

	if err != nil {
		http.Error(w, "createuser: "+err.Error(), 500)
		return
	}

	start, err := startSvc.GetStartByID(startID)

	if err != nil {
		http.Error(w, "getstart: "+err.Error(), 500)
		return
	}

	//get ship template from start
	temp, err := shipTmpSvc.GetShipTemplateByID(start.ShipTemplateID)

	if err != nil {
		http.Error(w, "getstartertemplate: "+err.Error(), 500)
		return
	}

	//create fitting from start
	const moduleCreationReason = "Module for starter ship of new player"

	fitting := sql.Fitting{
		ARack: []sql.FittedSlot{},
		BRack: []sql.FittedSlot{},
		CRack: []sql.FittedSlot{},
	}

	//create rack a modules
	for _, l := range start.ShipFitting.ARack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
		})

		if err != nil {
			http.Error(w, "createitemforstartership(a): "+err.Error(), 500)
			return
		}

		//link item to slot
		fitting.ARack = append(fitting.ARack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create rack b modules
	for _, l := range start.ShipFitting.BRack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
		})

		if err != nil {
			http.Error(w, "createitemforstartership(b): "+err.Error(), 500)
			return
		}

		//link item to slot
		fitting.BRack = append(fitting.BRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create rack c modules
	for _, l := range start.ShipFitting.CRack {
		//create item for slot
		i, err := itemSvc.NewItem(sql.Item{
			ItemTypeID:    l.ItemTypeID,
			Meta:          sql.Meta{},
			CreatedBy:     &u.ID,
			CreatedReason: moduleCreationReason,
		})

		if err != nil {
			http.Error(w, "createitemforstartership(c): "+err.Error(), 500)
			return
		}

		//link item to slot
		fitting.CRack = append(fitting.CRack, sql.FittedSlot{
			ItemID:     i.ID,
			ItemTypeID: l.ItemTypeID,
		})
	}

	//create starter ship
	t := sql.Ship{
		SystemID:       start.SystemID,
		UserID:         u.ID,
		ShipName:       fmt.Sprintf("%s's Starter Ship", m.Username),
		Texture:        temp.Texture,
		Theta:          0,
		Shield:         temp.BaseShield,
		Armor:          temp.BaseArmor,
		Hull:           temp.BaseHull,
		Fuel:           temp.BaseFuel,
		Heat:           0,
		Energy:         temp.BaseEnergy,
		ShipTemplateID: temp.ID,
		PosX:           float64(physics.RandInRange(-50000, 50000)),
		PosY:           float64(physics.RandInRange(-50000, 50000)),
		Fitting:        fitting,
		Destroyed:      false,
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
func (l *HTTPListener) HandleLogin(w http.ResponseWriter, r *http.Request) {
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

//HandleShutdown Shuts down the server after saving the game state
func (l *HTTPListener) HandleShutdown(w http.ResponseWriter, r *http.Request) {
	//enable cors
	enableCors(&w)

	//get auth token
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		return
	}

	key := keys[0]

	//load listener configuration
	config, err := loadConfiguration()
	if err != nil {
		return
	}

	//validate auth token
	if config.ShutdownToken == key {
		//initiate shutdown
		l.Engine.Shutdown()
	}
}
