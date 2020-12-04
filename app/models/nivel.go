package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Nivel struct {
	Id         string `gorm:"primaryKey;"`
	Nom  	   string
	Desc	   string
	Grados []Grado
}

func (tab Nivel) ToString() string {
	return tab.Nom
}

func (tab *Nivel) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (nivel Nivel) FindAll(conn *gorm.DB) ([]Nivel, error) {
	var niveles []Nivel
	if err := conn.Preload("Grados").Find(&niveles).Error; err != nil {
		return nil, err
	}
	return niveles, nil
}

func (nivel Nivel) GetAll(conn *gorm.DB) ([]Nivel, error) {
	var niveles []Nivel
	if err := conn.Find(&niveles).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return niveles, nil
}
