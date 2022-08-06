package subject

import (
	"observer/model"
	"observer/observer"
)

type Subject interface {
	Register(observer.Observer)
	Unregister(observer.Observer)
	NotifyAll(model.Message)
}
