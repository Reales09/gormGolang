package handlers

import (
	"gorm/db"
	"gorm/models"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

/*
func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendFound(rw)
	} else {
		models.SendData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnProcessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendFound(rw)
	} else {
		userId = user.Id
	}

	//Actualizar un registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnProcessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {

	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err

	} else {
		return *user, nil
	}
}
*/
