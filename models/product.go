package models

type Product struct {
	Id          int64   `gorm:"primaryKey" json:"id"`
	NamaProduct string  `json:"nama_product" validate:"required"`
	Stok        int64   `json:"stok" validate:"required,number"`
	Harga       float64 `json:"harga" validate:"required,number"`
	Orders      []Order `gorm:"foreignKey:IdProduct"` // Relasi One-to-Many
}

func GetAllProducts(products *[]Product) error {
	return DB.Find(&products).Error
}

func (p *Product) CreateProduct() error {
	return DB.Create(p).Error
}

func GetProductByID(id uint, product *Product) error {
	return DB.First(product, id).Error
}

func UpdateProduct(id int64, product *Product) error {
	product.Id = id
	return DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return DB.Delete(&Product{}, id).Error
}
