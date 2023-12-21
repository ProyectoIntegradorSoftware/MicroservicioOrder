package dominio

type ActualizarOrdenInput struct {
	Nombre      *string         `json:"nombre,omitempty"`
	SKU         *string         `json:"sku,omitempty"`
	Precio      *string         `json:"precio,omitempty"`
	Descripcion *string         `json:"descripcion,omitempty"`
	Productos   []ProductoInput `json:"productos,omitempty"`
}

type CrearOrdenInput struct {
	Nombre      string          `json:"nombre"`
	SKU         string          `json:"sku"`
	Precio      string          `json:"precio"`
	Descripcion string          `json:"descripcion"`
	Productos   []ProductoInput `json:"productos,omitempty"`
}

type ProductoInput struct {
	Nombre      string `json:"nombre"`
	Precio      string `json:"precio"`
	Descripcion string `json:"descripcion"`
}

type RespuestaEliminacion struct {
	Mensaje string `json:"mensaje"`
}

type Orden struct {
	ID          string     `json:"id"`
	Nombre      string     `json:"nombre"`
	SKU         string     `json:"sku"`
	Precio      string     `json:"precio"`
	Descripcion string     `json:"descripcion"`
	Productos   []Producto `json:"productos,omitempty"`
}

type Producto struct {
	Nombre      string `json:"nombre"`
	Precio      string `json:"precio"`
	Descripcion string `json:"descripcion"`
}

type CrearProductoInput struct {
	Nombre      string `json:"nombre"`
	SKU         string `json:"sku"`
	Precio      string `json:"precio"`
	Descripcion string `json:"descripcion"`
}
