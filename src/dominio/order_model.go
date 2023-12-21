package dominio

import (
	"strconv"
)

// OrdenGORM es el modelo de orden para GORM de Orden
type OrdenGORM struct {
	ID             uint           `gorm:"primaryKey:autoIncrement" json:"id"`
	UserId         string         `gorm:"type:varchar(255);not null"`
	ListaProductos []ProductoGORM `gorm:"many2many:orden_productos;" json:"lista_productos"`
}

// TableName especifica el nombre de la tabla para OrdenGORM
func (OrdenGORM) TableName() string {
	return "ordens"
}

func (ordenGORM *OrdenGORM) ToGQL() (*Orden, error) {
	// Convertir la lista de productos a un formato que desees
	var productos []*Producto
	for _, productoGORM := range ordenGORM.ListaProductos {
		productos = append(productos, &Producto{
			Nombre:      productoGORM.Nombre,
			Precio:      productoGORM.Precio,
			Descripcion: productoGORM.Descripcion,
		})
	}

	return &Orden{
		ID:          strconv.Itoa(int(ordenGORM.ID)),
		Nombre:      "",
		SKU:         "",
		Precio:      "",
		Descripcion: "",
		Productos:   []Producto{},
	}, nil
}
