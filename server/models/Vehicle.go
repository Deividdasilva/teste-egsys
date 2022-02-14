package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// Vehicle_ID
// Fleet_ID
// Name
// Max_Speed

type Vehicle struct {
	Vehicle_ID uint64  `gorm:"primary_key;auto_increment;" json:"vehicle_id"`
	Name       string  `gorm:"size:100;not null;" json:"name"`
	Max_Speed  float32 `gorm:"not null;" json:"max_speed"`
	Fleet      Fleet   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"fleet"`
	Fleet_ID   uint64  `sql:"type:int REFERENCES fleets(fleet_id)" json:"fleet_id"`
}

func (v *Vehicle) PrepareVehicle() {
	v.Vehicle_ID = 0
	v.Name = html.EscapeString(strings.TrimSpace(v.Name))
	v.Fleet = Fleet{}
}

func (v *Vehicle) ValidateVehicle() error {
	if v.Name == "" {
		return errors.New("Nome da frota deve ser informado")
	}
	if v.Fleet_ID == 0 {
		return errors.New("É necessario adicionar uma frota")
	}
	if v.Max_Speed < 0 && v.Fleet_ID == 0 {
		return errors.New("Velocidade invalida")
	}
	return nil
}

func (v *Vehicle) SaveVehicle(db *gorm.DB) (*Vehicle, error) {
	if v.Max_Speed == 0 {
		err := db.Debug().Model(&Fleet{}).Where("fleet_id = ?", v.Fleet_ID).Take(&v.Fleet).Error
		if err != nil {
			return &Vehicle{}, err
		}
		v.Max_Speed = v.Fleet.Max_Speed
	}
	err := db.Debug().Model(&Vehicle{}).Create(&v).Error
	if err != nil {
		return &Vehicle{}, err
	}
	if v.Vehicle_ID != 0 {
		err = db.Debug().Model(&Fleet{}).Where("fleet_id = ?", v.Fleet_ID).Take(&v.Fleet).Error
		if err != nil {
			return &Vehicle{}, err
		}
	}
	return v, nil
}

func (v *Vehicle) FindAllVehicle(db *gorm.DB) (*[]Vehicle, error) {
	vehicles := []Vehicle{}
	err := db.Debug().Model(&Vehicle{}).Find(&vehicles).Error
	if err != nil {
		return &[]Vehicle{}, err
	}
	if len(vehicles) > 0 {
		for i := range vehicles {
			err := db.Debug().Model(&Vehicle{}).Where("fleet_id = ?", vehicles[i].Fleet_ID).Take(&vehicles[i].Fleet).Error
			if err != nil {
				return &[]Vehicle{}, err
			}
		}
	}
	return &vehicles, nil
}

func (v *Vehicle) DeleteAllVehicle(db *gorm.DB) (*Vehicle, error) {
	err := db.Debug().Model(&Vehicle{}).Delete(&Vehicle{}).Error
	if err != nil {
		return &Vehicle{}, err
	}
	return v, nil
}
