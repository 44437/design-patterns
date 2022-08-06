package observer

import "observer/model"

type Observer interface {
	Update(model.Message)
	GetId() string
}
