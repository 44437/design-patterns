package subject

import (
	"golang.org/x/exp/slices"
	"observer/model"
	"observer/observer"
)

type Provider struct {
	Observers []observer.Observer
}

func (p *Provider) Register(observer observer.Observer) {
	p.Observers = append(p.Observers, observer)
}

func (p *Provider) Unregister(observer observer.Observer) {
	var ids []string
	for _, observerItem := range p.Observers {
		ids = append(ids, observerItem.GetId())
	}

	index := slices.Index(ids, observer.GetId())

	if index >= 0 {
		p.Observers = append(p.Observers[:index], p.Observers[index+1:]...)
	}
}

func (p *Provider) NotifyAll(message model.Message) {
	for _, observerItem := range p.Observers {
		observerItem.Update(message)
	}
}

func (p *Provider) Push(message model.Message) {
	p.NotifyAll(message)
}
