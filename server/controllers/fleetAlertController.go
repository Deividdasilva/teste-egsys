package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"teste-egsys/server/models"
	"teste-egsys/server/responses"

	"github.com/gorilla/mux"
)

func (server *Server) CreateFleetAlert(w http.ResponseWriter, r *http.Request) {

	Fleet_Alert := models.Fleet_Alert{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &Fleet_Alert)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	Fleet_Alert.Fleet_ID = pid
	Fleet_Alert.PrepareFleetAlert()
	err = Fleet_Alert.ValidateFleetAlert()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	Fleet_AlertCreated, err := Fleet_Alert.SaveFleetAlert(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, Fleet_AlertCreated))
	responses.JSON(w, http.StatusCreated, Fleet_AlertCreated)
}

func (server *Server) GetAllFleetAlerts(w http.ResponseWriter, r *http.Request) {

	Fleet_Alert := models.Fleet_Alert{}

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	Fleet_Alerts, err := Fleet_Alert.FindAllFleetAlert(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Fleet_Alerts)
}

func (server *Server) DeleteAllFleetAlert(w http.ResponseWriter, r *http.Request) {
	Fleet_Alert := models.Fleet_Alert{}

	Fleet_Alerts, err := Fleet_Alert.DeleteAllFleetAlert(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Fleet_Alerts)
}
