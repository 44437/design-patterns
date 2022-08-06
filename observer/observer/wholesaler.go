package observer

import (
	"fmt"
	"golang.org/x/exp/slices"
	"observer/model"
)

type Product string

type Wholesaler struct {
	ID       string
	Name     string
	Products []Product
}

func (w *Wholesaler) Update(message model.Message) {
	fmt.Printf("Dear %s, New Message : %s\n", w.Name, message.Content)
}

func (w *Wholesaler) GetId() string {
	return w.ID
}

func (w *Wholesaler) Sell(product Product) {
	index := slices.Index(w.Products, product)
	if index >= 0 {
		w.Products = append(w.Products[:index], w.Products[index+1:]...)
	}
}
