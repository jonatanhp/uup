package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Grado struct {
	Id string `gorm:"primary_key;"`
	//Fecha  string
	Nomg string
	NivelId string `gorm:"size:191"`
	Nivel   Nivel //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Grado) ToString() string {
	return fmt.Sprintf("id: %d\nNomg: %s", tab.Id, tab.Nomg)
}

func (tab *Grado) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Nivel   Nivel //para crear el FK `gorm:"foreignkey:NivelId"`
