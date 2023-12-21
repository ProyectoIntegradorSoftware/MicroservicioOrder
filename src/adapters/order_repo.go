package adapters

import (
	"errors"
	"fmt"
	"log"

	"github.com/ProyectoIntegradorSoftware/MicroservicioOrder/database"
	model "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/dominio"
	"github.com/ProyectoIntegradorSoftware/MicroservicioOrder/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type orderRepository struct {
	db *database.DB
}

func NewOrderRepository(db *database.DB) ports.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

// ObtenerOrden obtiene un orden por su ID.
func (ur *orderRepository) Orden(id string) (*model.Orden, error) {
	if id == "" {
		return nil, errors.New("El ID de orden es requerido")
	}

	var ordenGORM model.OrdenGORM
	result := ur.db.GetConn().First(&ordenGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el orden con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return ordenGORM.ToGQL()
}

// Ordens obtiene todos los ordens de la base de datos.
func (ur *orderRepository) Ordens() ([]*model.Orden, error) {
	var ordensGORM []model.OrdenGORM
	result := ur.db.GetConn().Find(&ordensGORM)

	if result.Error != nil {
		log.Printf("Error al obtener los ordens: %v", result.Error)
		return nil, result.Error
	}

	var ordens []*model.Orden
	for _, ordenGORM := range ordensGORM {
		orden, _ := ordenGORM.ToGQL()
		ordens = append(ordens, orden)
	}

	return ordens, nil
}
func (ur *orderRepository) CrearOrden(input model.CrearOrdenInput) (*model.Orden, error) {
	// Convertir la lista de productos en modelo ProductoGORM
	var productosGORM []model.ProductoGORM
	for _, productoInput := range input.Productos {
		productoGORM := &model.ProductoGORM{
			Nombre:      productoInput.Nombre,
			Precio:      productoInput.Precio,
			Descripcion: productoInput.Descripcion,
			// Otros campos del producto si es necesario
		}
		productosGORM = append(productosGORM, *productoGORM)
	}

	// Crear la orden con la lista de productos
	ordenGORM := &model.OrdenGORM{
		ListaProductos: productosGORM,
		// Otros campos de la orden si es necesario
	}

	result := ur.db.GetConn().Create(&ordenGORM)
	if result.Error != nil {
		log.Printf("Error al crear orden: %v", result.Error)
		return nil, result.Error
	}

	response, err := ordenGORM.ToGQL()
	return response, err
}
func (ur *orderRepository) ActualizarOrden(id string, input *model.ActualizarOrdenInput) (*model.Orden, error) {
	var ordenGORM model.OrdenGORM
	if id == "" {
		return nil, errors.New("El ID de orden es requerido")
	}

	result := ur.db.GetConn().Preload("ListaProductos").First(&ordenGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Orden con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Actualizar la lista de productos
	if input.Productos != nil {
		// Convertir la lista de productos en modelo ProductoGORM
		var productosGORM []model.ProductoGORM
		for _, productoInput := range input.Productos {
			productoGORM := &model.ProductoGORM{
				Nombre:      productoInput.Nombre,
				Precio:      productoInput.Precio,
				Descripcion: productoInput.Descripcion,
				// Otros campos del producto si es necesario
			}
			productosGORM = append(productosGORM, *productoGORM)
		}
		ordenGORM.ListaProductos = productosGORM
	}

	result = ur.db.GetConn().Save(&ordenGORM)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Printf("Orden actualizado: %v", ordenGORM)

	// Convertir el modelo GORM actualizado a modelo GraphQL
	response, err := ordenGORM.ToGQL()
	return response, err
}

// EliminarOrden elimina un orden de la base de datos por su ID.
func (ur *orderRepository) EliminarOrden(id string) (*model.RespuestaEliminacion, error) {
	// Intenta buscar el orden por su ID
	var ordenGORM model.OrdenGORM
	result := ur.db.GetConn().First(&ordenGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El orden no se encontró en la base de datos
			response := &model.RespuestaEliminacion{
				Mensaje: "El orden no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el orden con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al buscar el orden",
		}
		return response, result.Error
	}

	// Elimina el orden de la base de datos
	result = ur.db.GetConn().Delete(&ordenGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el orden con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al eliminar el orden",
		}
		return response, result.Error
	}

	// Éxito al eliminar el orden
	response := &model.RespuestaEliminacion{
		Mensaje: "Orden eliminado con éxito",
	}
	return response, result.Error

}
