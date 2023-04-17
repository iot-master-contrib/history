package internal

import "github.com/zgwit/iot-master/v3/pkg/lib"

type Device struct {
	//Id     string
	Update int64
	Store  map[string]string
	Values map[string]float64
	Last   map[string]float64
}

func NewDevice() *Device {
	return &Device{
		Values: map[string]float64{},
		Last:   map[string]float64{},
	}
}

func EnsureDevice(id string) *Device {
	dev := devices.Load(id)
	if dev == nil {
		dev = NewDevice()
		devices.Store(id, dev)
	}
	return dev
}

var devices lib.Map[Device]
