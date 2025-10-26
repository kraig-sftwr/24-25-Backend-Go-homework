package ty

type Common struct {
	Name  string
	Price float64
	Stock float64
}

func NewCommon(name string, price float64, stock float64) Common {
	return Common{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

type Commoner interface {
	GetName() string
	GetPrice() float64
	GetStock() float64
}

func (c *Common) GetName() string {
	return c.Name
}
func (c *Common) GetPrice() float64 {
	return c.Price
}
func (c *Common) GetStock() float64 {
	return c.Stock
}

func (c *Common) SetStock(stock float64) {
	c.Stock = stock
}
