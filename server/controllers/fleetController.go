package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"teste-egsys/server/models"
	"teste-egsys/server/responses"
)

func (server *Server) CreateFleet(w http.ResponseWriter, r *http.Request) {

	fleet := models.Fleet{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &fleet)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fleet.Prepare()
	err = fleet.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fleetCreated, err := fleet.SaveFleet(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, fleetCreated))
	responses.JSON(w, http.StatusCreated, fleetCreated)
}

func (server *Server) GetAllFleets(w http.ResponseWriter, r *http.Request) {

	fleet := models.Fleet{}

	fleets, err := fleet.FindAllFleet(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fleets)
}

func (server *Server) DeleteAllFleet(w http.ResponseWriter, r *http.Request) {
	fleet := models.Fleet{}

	fleets, err := fleet.DeleteAllFleet(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fleets)
}
