package model

import (
	"api/db"
)

type Categoria struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
	Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria

func Migraciones() {
	db.Database.AutoMigrate(&Categoria{})

}
