package ty

type ElectronicDevice struct {
	Common
	Model float64
	Brand string
}

func NewElectronicDevice(com *Common, m float64, b string) ElectronicDevice {
	return ElectronicDevice{
		Common: Common{
			Name:  com.Name,
			Price: com.Price,
			Stock: com.Stock,
		},
		Model: m,
		Brand: b,
	}
}

func (e *ElectronicDevice) GetModel() float64 {
	return e.Model
}
func (e *ElectronicDevice) GetBrand() string {
	return e.Brand
}
