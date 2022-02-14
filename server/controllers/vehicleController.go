package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"teste-egsys/server/models"
	"teste-egsys/server/responses"
)

func (server *Server) CreateVehicle(w http.ResponseWriter, r *http.Request) {

	Vehicle := models.Vehicle{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &Vehicle)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Vehicle.PrepareVehicle()
	err = Vehicle.ValidateVehicle()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	VehicleCreated, err := Vehicle.SaveVehicle(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, VehicleCreated))
	responses.JSON(w, http.StatusCreated, VehicleCreated)
}

func (server *Server) GetAllVehicles(w http.ResponseWriter, r *http.Request) {

	Vehicle := models.Vehicle{}

	Vehicles, err := Vehicle.FindAllVehicle(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Vehicles)
}

func (server *Server) DeleteAllVehicle(w http.ResponseWriter, r *http.Request) {
	Vehicle := models.Vehicle{}

	Vehicles, err := Vehicle.DeleteAllVehicle(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Vehicles)
}
