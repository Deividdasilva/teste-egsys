package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// Fleet_ID
// Name
// Max_Speed

type Fleet struct {
	Fleet_ID  uint64  `gorm:"primary_key;auto_increment" json:"fleet_id"`
	Name      string  `gorm:"size:100;not null;" json:"name"`
	Max_Speed float32 `gorm:"not null;" json:"max_speed"`
}

func (f *Fleet) Prepare() {
	f.Fleet_ID = 0
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
}

func (f *Fleet) Validate() error {
	if f.Name == "" {
		return errors.New("Nome da frota deve ser informado")
	}
	if f.Max_Speed < 0 {
		return errors.New("Velocidade invalida")
	}
	return nil
}

func (f *Fleet) SaveFleet(db *gorm.DB) (*Fleet, error) {
	err := db.Debug().Model(&Fleet{}).Create(&f).Error
	if err != nil {
		return &Fleet{}, err
	}
	return f, nil
}

func (f *Fleet) FindAllFleet(db *gorm.DB) (*[]Fleet, error) {
	fleets := []Fleet{}
	err := db.Debug().Model(&Fleet{}).Find(&fleets).Error
	if err != nil {
		return &[]Fleet{}, err
	}
	return &fleets, nil
}

func (f *Fleet) DeleteAllFleet(db *gorm.DB) (*Fleet, error) {
	err := db.Debug().Model(&Fleet{}).Delete(&Fleet{}).Error
	if err != nil {
		return &Fleet{}, err
	}
	return f, nil
}
