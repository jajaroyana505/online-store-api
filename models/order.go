package models

type Order struct {
	Id        int64   `gorm:"primaryKey" json:"id" `
	IdProduct int64   `json:"id_product" validate:"required"`
	Jumlah    int64   `json:"jumlah" validate:"required"`
	Status    string  `json:"status" validate:"required"`
	Product   Product `gorm:"foreignKey:IdProduct"` // Relasi Many-to-One
}

func GetAllOrders(orders *[]Order) error {
	return DB.Preload("Product").Find(orders).Error
}

func (o *Order) CreateOrder() error {
	// o.Status = ""
	return DB.Create(o).Error
}
func GetOrderByID(id uint, order *Order) error {
	return DB.Preload("Product").First(order, id).Error
}

func UpdateOrder(order *Order) error {
	return DB.Save(order).Error
}

func DeleteOrder(id uint) error {
	return DB.Delete(&Order{}, id).Error

}
