package dominio

// ProductoGORM es el modelo de producto para GORM de Producto
type ProductoGORM struct {
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	SKU         string `gorm:"type:varchar(255);not null"`
	Precio      string `gorm:"type:varchar(255);not null"`
	Descripcion string `gorm:"type:varchar(255);not null"`
}

// TableName especifica el nombre de la tabla para ProductoGORM
func (ProductoGORM) TableName() string {
	return "productos"
}

func (productoGORM *ProductoGORM) ToGQL() (*Producto, error) {

	return &Producto{
		Nombre:      productoGORM.Nombre,
		Precio:      productoGORM.Precio,
		Descripcion: productoGORM.Descripcion,
	}, nil
}
