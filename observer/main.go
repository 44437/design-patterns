package main

import (
	"observer/model"
	"observer/observer"
	"observer/subject"
)

func main() {
	var observerCustomer, observerWholesaler observer.Observer

	observerCustomer = &observer.Customer{
		ID:   "123",
		Name: "Customer",
	}

	observerWholesaler = &observer.Wholesaler{
		ID:       "321",
		Name:     "Wholesaler",
		Products: nil,
	}

	var provider subject.Subject
	provider = &subject.Provider{}

	provider.Register(observerCustomer)
	provider.Register(observerWholesaler)

	provider.NotifyAll(model.Message{
		ID:      "message_1",
		Content: "Message 1",
	})

	provider.Unregister(observerWholesaler)

	provider.NotifyAll(model.Message{
		ID:      "message_2",
		Content: "Message 2",
	})
}
