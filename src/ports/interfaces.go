package ports

import (
	model "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/dominio"
)

// puerto de salida
type OrderRepository interface {
	CrearOrden(input model.CrearOrdenInput) (*model.Orden, error)
	Orden(id string) (*model.Orden, error)
	ActualizarOrden(id string, input *model.ActualizarOrdenInput) (*model.Orden, error)
	EliminarOrden(id string) (*model.RespuestaEliminacion, error)
	Ordens() ([]*model.Orden, error)
}
