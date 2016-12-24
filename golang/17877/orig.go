package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Order struct {
	Price float64
}

func (o *Order) UnmarshalJSON(data []byte) error {
	if o == nil {
		return errors.New("o is nil")
	}
	o.Price = 99

	return nil
}

type Outer struct {
	*Order
}

func main() {
	b := []byte(`{"Price":53.2}`)
	var m Outer

	// m.Order.Price will only be set if  `m.Order != nil`.
	// NOTE this works perfectly fine without initialization if `m.Order` is not an anonymous field(see example: https://play.golang.org/p/DYuygZrLgk)
	// m.Order = &Order{}

fmt.Printf("#%v\n", m)
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", m.Order)
}
