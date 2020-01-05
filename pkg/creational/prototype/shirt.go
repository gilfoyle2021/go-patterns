package prototype

import "fmt"

import "errors"

type Color byte

const (
	Red    = 1
	Blue   = 2
	Yellow = 3
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ShirtCache struct {
}

func (sc *ShirtCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case Red:
		newItem := *redPrototype
		return &newItem, nil
	case Yellow:
		newItem := *yelloPrototype
		return &newItem, nil
	case Blue:
		newItem := *bulePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Not find this type")
	}

}

func GetShirtsCloner() (ShirtCloner, error) {

	return nil, nil
}

type Shirt struct {
	Color Color
	Price float32
	SKU   string
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Color:%d price:%0.2f SKU:%s", s.Color, s.Price, s.SKU)
}

var redPrototype *Shirt = &Shirt{
	Color: Red,
	Price: 2.0,
	SKU:   "xxx",
}

var bulePrototype *Shirt = &Shirt{
	Color: Blue,
	Price: 1.0,
	SKU:   "yyy",
}

var yelloPrototype *Shirt = &Shirt{
	Color: Yellow,
	Price: 1.0,
	SKU:   "yyy",
}
