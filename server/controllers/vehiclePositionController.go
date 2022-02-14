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

func (server *Server) CreateVehiclePosition(w http.ResponseWriter, r *http.Request) {

	VehiclePosition := models.Vehicle_Position{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &VehiclePosition)
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
	VehiclePosition.Vehicle_ID = pid
	VehiclePosition.PrepareVehiclePosition()
	err = VehiclePosition.ValidateVehiclePosition()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	VehiclePositionCreated, err := VehiclePosition.SaveVehiclePosition(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, VehiclePositionCreated))
	responses.JSON(w, http.StatusCreated, VehiclePositionCreated)
}

func (server *Server) GetAllVehiclePositions(w http.ResponseWriter, r *http.Request) {

	VehiclePosition := models.Vehicle_Position{}

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	VehiclePositions, err := VehiclePosition.FindAllVehiclePosition(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, VehiclePositions)
}

func (server *Server) DeleteAllVehiclePosition(w http.ResponseWriter, r *http.Request) {
	VehiclePosition := models.Vehicle_Position{}

	VehiclePositions, err := VehiclePosition.DeleteAllVehiclePosition(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, VehiclePositions)
}
