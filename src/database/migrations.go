package database

import (
	"log"

	model "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.OrdenGORM{})

	log.Println("Migraciones completadas")
}
