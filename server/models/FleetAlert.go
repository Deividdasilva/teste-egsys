package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Fleet_Alert struct {
	Fleet_Alert_ID uint64 `gorm:"primary_key;auto_increment;" json:"fleet_alert_id"`
	WebHook        string `gorm:"size:100;not null;" json:"webhook"`
	Fleet          Fleet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"fleet"`
	Fleet_ID       uint64 `sql:"type:int REFERENCES fleets(fleet_id)" json:"fleet_id"`
}

func (a *Fleet_Alert) PrepareFleetAlert() {
	a.Fleet_Alert_ID = 0
	a.WebHook = html.EscapeString(strings.TrimSpace(a.WebHook))
	a.Fleet = Fleet{}
}

func (a *Fleet_Alert) ValidateFleetAlert() error {
	if a.WebHook == "" {
		return errors.New("Url do WebHook deve ser informado")
	}
	if a.Fleet_ID == 0 {
		return errors.New("É necessario adicionar uma frota")
	}
	return nil
}

func (a *Fleet_Alert) SaveFleetAlert(db *gorm.DB) (*Fleet_Alert, error) {
	err := db.Debug().Model(&Vehicle{}).Create(&a).Error
	if err != nil {
		return &Fleet_Alert{}, err
	}
	if a.Fleet_Alert_ID != 0 {
		err = db.Debug().Model(&Fleet{}).Where("fleet_id = ?", a.Fleet_ID).Take(&a.Fleet).Error
		if err != nil {
			return &Fleet_Alert{}, err
		}
	}
	return a, nil
}

func (a *Fleet_Alert) FindAllFleetAlert(db *gorm.DB, pid uint64) (*[]Fleet_Alert, error) {
	alerts := []Fleet_Alert{}
	err := db.Debug().Model(&Fleet_Alert{}).Where("fleet_id = ?", pid).Find(&alerts).Error
	if err != nil {
		return &[]Fleet_Alert{}, err
	}
	if len(alerts) > 0 {
		for i := range alerts {
			err := db.Debug().Model(&Vehicle{}).Where("fleet_id = ?", alerts[i].Fleet_ID).Take(&alerts[i].Fleet).Error
			if err != nil {
				return &[]Fleet_Alert{}, err
			}
		}
	}
	return &alerts, nil
}

func (a *Fleet_Alert) DeleteAllFleetAlert(db *gorm.DB) (*Fleet_Alert, error) {
	err := db.Debug().Model(&Fleet_Alert{}).Delete(&Fleet_Alert{}).Error
	if err != nil {
		return &Fleet_Alert{}, err
	}
	return a, nil
}
