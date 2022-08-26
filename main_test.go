package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teste-egsys/server/models"
	"testing"
)

func TestFleetGet(t *testing.T) {
	req, err := http.Get("http://localhost:8081/fleet")
	if err != nil {
		t.Fatal(err)
	}
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFleetCreate(t *testing.T) {
	fleetTest := models.Fleet{
		Fleet_ID:  1029,
		Name:      "Teste fleet",
		Max_Speed: 25.5,
	}
	responseBody, err := json.Marshal(fleetTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.Post("http://localhost:8081/fleet", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}
	status := req.StatusCode
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVehicleGet(t *testing.T) {
	req, err := http.Get("http://localhost:8081/vehicle")
	if err != nil {
		t.Fatal(err)
	}
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVehicleCreate(t *testing.T) {
	vehicleTest := models.Vehicle{
		Name:      "teste vehicle",
		Max_Speed: 100,
		Fleet_ID:  1,
	}
	responseBody, err := json.Marshal(vehicleTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.Post("http://localhost:8081/vehicle", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}
	status := req.StatusCode
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFleetAlertGet(t *testing.T) {
	req, err := http.Get("http://localhost:8081/fleet/1/alert")
	if err != nil {
		t.Fatal(err)
	}
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFleetAlertCreate(t *testing.T) {
	fleetAlertTest := models.Fleet_Alert{
		WebHook: "http://localhost:8081/fleet/alert",
	}
	responseBody, err := json.Marshal(fleetAlertTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.Post("http://localhost:8081/fleet/1/alert", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}
	status := req.StatusCode
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVehiclePositionGet(t *testing.T) {
	req, err := http.Get("http://localhost:8081/vehicle/3/positon")
	if err != nil {
		t.Fatal(err)
	}
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVehiclePositionCreate(t *testing.T) {
	fleetVehiclePosition := models.Vehicle_Position{
		Latitude:      -32,
		Longitude:     -15,
		Current_Speed: 25,
	}
	responseBody, err := json.Marshal(fleetVehiclePosition)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.Post("http://localhost:8081/vehicle/3/positon", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}
	status := req.StatusCode
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestReset(t *testing.T) {
	// Create request
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8081/reset", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
