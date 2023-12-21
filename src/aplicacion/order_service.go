package service

import (
	"context"

	model "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/dominio"
	repository "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/ports"
	pb "github.com/ProyectoIntegradorSoftware/MicroservicioOrder/proto"
)

// este servicio implementa la interfaz OrderServiceServer
// que se genera a partir del archivo proto
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Crear la lista de productos
	var productosInput []model.CrearProductoInput
	for _, prod := range req.GetProductos() {
		productoInput := model.CrearProductoInput{
			Nombre:      prod.GetNombre(),
			Precio:      prod.GetPrecio(),
			Descripcion: prod.GetDescripcion(),
		}
		productosInput = append(productosInput, productoInput)
	}

	// Crear la orden con la lista de productos
	crearOrdenInput := model.CrearOrdenInput{
		Nombre:      req.GetNombre(),
		SKU:         req.GetSKU(),
		Precio:      req.GetPrecio(),
		Descripcion: req.GetDescripcion(),
	}

	u, err := s.repo.CrearOrden(crearOrdenInput)
	if err != nil {
		return nil, err
	}

	// Crear la respuesta protobuf
	response := &pb.CreateOrderResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
	}

	return response, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	u, err := s.repo.Orden(req.GetId())
	if err != nil {
		return nil, err
	}

	// Crear la lista de productos en la respuesta protobuf
	var productos []*pb.Producto
	for _, prod := range u.Productos {
		product := &pb.Producto{
			Nombre:      prod.Nombre,
			Precio:      prod.Precio,
			Descripcion: prod.Descripcion,
		}
		productos = append(productos, product)
	}

	// Crear la respuesta protobuf
	response := &pb.GetOrderResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
		Productos:   productos,
	}

	return response, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.repo.Ordens()
	if err != nil {
		return nil, err
	}

	var response []*pb.Order
	for _, u := range orders {
		// Crear la lista de productos en la respuesta protobuf
		var productos []*pb.Producto
		for _, prod := range u.Productos {
			product := &pb.Producto{
				Nombre:      prod.Nombre,
				Precio:      prod.Precio,
				Descripcion: prod.Descripcion,
			}
			productos = append(productos, product)
		}

		// Crear la respuesta protobuf para cada orden
		order := &pb.Order{
			Id:          u.ID,
			Nombre:      u.Nombre,
			SKU:         u.SKU,
			Precio:      u.Precio,
			Descripcion: u.Descripcion,
			Productos:   productos,
		}
		response = append(response, order)
	}

	return &pb.ListOrdersResponse{Orders: response}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	nombre := req.GetNombre()
	SKU := req.GetSKU()
	precio := req.GetPrecio()
	descripcion := req.GetDescripcion()

	// Crear la lista de productos en el mensaje de actualización
	var productosInput []model.CrearProductoInput
	for _, prod := range req.GetProductos() {
		productoInput := model.CrearProductoInput{
			Nombre:      prod.GetNombre(),
			Precio:      prod.GetPrecio(),
			Descripcion: prod.GetDescripcion(),
		}
		productosInput = append(productosInput, productoInput)
	}

	// Crear el mensaje de actualización
	actualizarOrdenInput := &model.ActualizarOrdenInput{
		Nombre:      &nombre,
		SKU:         &SKU,
		Precio:      &precio,
		Descripcion: &descripcion,
	}

	// Actualizar la orden en el repositorio
	u, err := s.repo.ActualizarOrden(req.GetId(), actualizarOrdenInput)
	if err != nil {
		return nil, err
	}

	// Crear la respuesta protobuf
	response := &pb.UpdateOrderResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
	}

	return response, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	respuesta, err := s.repo.EliminarOrden(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeleteOrderResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}
