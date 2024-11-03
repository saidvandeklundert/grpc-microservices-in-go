package db
import (
	"fmt"
	"github.com/saidvandeklundert/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	)
type Order struct{
	gorm.Model // adds entity metadata such as ID to struct
	CustomerId int64
	Status string
	OrderItems []OrderItem // reference to OrderItem
}	

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice float32
	Quantity int32
	OrderID uint // back reference to Order model
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter (dataSourceUrl string) (*Adapter,error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr!= nil{
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Order{}, OrderItem{}) // be sure the tables are created correctly
	if err != nil{
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db:db}, nil
}

// Get method returns the domain.Order core model
func (a Adapter) Get(id string) (domain.Order, error){
	var orderEntity Order

	// finds by ID and puts it into orderEntity
	res := a.db.First(&orderEntity,id)
	var orderItems []domain.OrderItem

	// converts Order items
	for _,orderItem := range orderEntity.OrderItems{
		orderItems=append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}
	// converts Order
	order := domain.Order{
		ID: int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerId,
		Status: orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt: orderEntity.CreatedAt.UnixNano(),
	}
	return order,res.Error
}

// accepts the domain.Order core model
func (a Adapter) Save(order *domain.Order) error{
	var orderItems []OrderItem

	// converts Order items
	for _, orderItem := range order.OrderItems{
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}

	// converts order
	orderModel := Order{
		CustomerId: order.CustomerID,
		Status: order.Status,
		OrderItems: orderItems,
	}
	// saves data into database
	res := a.db.Create(&orderModel)
	if res.Error == nil{
		order.ID = int64(orderModel.ID)
	}

	return res.Error
}