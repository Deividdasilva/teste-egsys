package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

//Vehicle_Position_ID Vehicle_ID Timestamp Latitude Longitude Current_Speed Max_Speed

type Vehicle_Position struct {
	Vehicle_Position_ID uint64    `gorm:"primary_key;auto_increment;" json:"vehicle_position_id"`
	Timestamp           time.Time `json:"timestamp"`
	Latitude            float32   `gorm:"not null;" json:"latitude"`
	Longitude           float32   `gorm:"not null;" json:"longitude"`
	Current_Speed       float32   `gorm:"not null;" json:"current_speed"`
	Max_Speed           float32   `gorm:"not null;" json:"max_speed"`
	Vehicle             Vehicle   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vehicle"`
	Vehicle_ID          uint64    `sql:"type:int REFERENCES vehicles(vehicle_id)" json:"vehicle_id"`
}

func (vp *Vehicle_Position) PrepareVehiclePosition() {
	vp.Vehicle_Position_ID = 0
	vp.Timestamp = time.Now()
	vp.Vehicle = Vehicle{}
}

func (vp *Vehicle_Position) ValidateVehiclePosition() error {
	if vp.Latitude == 0 {
		return errors.New("Latitude deve ser informada")
	}
	if vp.Longitude == 0 {
		return errors.New("Longitude deve ser informada")
	}
	if vp.Current_Speed == 0 {
		return errors.New("CurrentSpeed deve ser informada")
	}
	if vp.Vehicle_ID == 0 {
		return errors.New("vehicle_ID deve ser informado")
	}
	return nil
}

func (vp *Vehicle_Position) SaveVehiclePosition(db *gorm.DB) (*Vehicle_Position, error) {
	err := db.Debug().Model(&Vehicle{}).Where("vehicle_id = ?", vp.Vehicle_ID).Find(&vp.Vehicle).Error
	vp.Max_Speed = vp.Vehicle.Max_Speed
	if vp.Current_Speed > vp.Vehicle.Max_Speed {
		data, err := json.Marshal(vp)
		if err != nil {
			fmt.Println(err.Error())
		}

		fleetAlert := Fleet_Alert{}
		err = db.Debug().Model(&fleetAlert).Where("fleet_id = ?", vp.Vehicle.Fleet_ID).Find(&fleetAlert).Error
		_, err = http.Post(fleetAlert.WebHook, "application/json",
			bytes.NewBuffer(data))

		if err != nil {
			fmt.Println(err.Error())
		}

	}

	err = db.Debug().Model(&Vehicle_Position{}).Create(&vp).Error
	if err != nil {
		return &Vehicle_Position{}, err
	}
	if vp.Vehicle_Position_ID != 0 {
		err = db.Debug().Model(&Vehicle{}).Where("vehicle_id = ?", vp.Vehicle_ID).Take(&vp.Vehicle).Error
		if err != nil {
			return &Vehicle_Position{}, err
		}
	}
	return vp, nil
}

func (vp *Vehicle_Position) FindAllVehiclePosition(db *gorm.DB, pid uint64) (*[]Vehicle_Position, error) {
	vehicles := []Vehicle_Position{}
	err := db.Debug().Model(&Vehicle_Position{}).Where("vehicle_id = ?", pid).Find(&vehicles).Error
	if err != nil {
		return &[]Vehicle_Position{}, err
	}
	if len(vehicles) > 0 {
		for i := range vehicles {
			err := db.Debug().Model(&Vehicle_Position{}).Where("vehicle_id = ?", vehicles[i].Vehicle_ID).Take(&vehicles[i].Vehicle).Error
			if err != nil {
				return &[]Vehicle_Position{}, err
			}
		}
	}
	return &vehicles, nil
}

func (vp *Vehicle_Position) DeleteAllVehiclePosition(db *gorm.DB) (*Vehicle_Position, error) {
	err := db.Debug().Model(&Vehicle_Position{}).Delete(&Vehicle_Position{}).Error
	if err != nil {
		return &Vehicle_Position{}, err
	}
	return vp, nil
}
