package observer

import (
	"fmt"
	"observer/model"
)

type Customer struct {
	ID   string
	Name string
}

func (c *Customer) Update(message model.Message) {
	fmt.Printf("Dear %s, New Message : %s\n", c.Name, message.Content)
}

func (c *Customer) GetId() string {
	return c.ID
}
