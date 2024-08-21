package chain_of_responsibility_pattern

import (
	"gorm.io/gorm"
	"testing"
)

type HandlerData struct {
	Handler1 string
	Handler2 string
}

type Handler interface {
	Execute(*HandlerData, *gorm.DB) error
	SetNext(Handler) Handler
	Do(*HandlerData, *gorm.DB) error
}

type BaseHandler struct {
	next Handler
}

type StartHandler struct {
	BaseHandler
}

func (h *BaseHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *BaseHandler) Do(data *HandlerData, db *gorm.DB) error {
	if h.next != nil {
		return h.next.Do(data, db)
	}

	return h.next.Execute(data, db)
}

func (h *BaseHandler) Execute(data *HandlerData, db *gorm.DB) error {
	// do something

	if h.next != nil {
		return h.next.Do(data, db)
	}

	return nil
}

type Handler1 struct {
	BaseHandler
}

func (h *Handler1) Execute(data *HandlerData, db *gorm.DB) error {
	// do something
	data.Handler1 = "handler1"
	return nil
}

type Handler2 struct {
	BaseHandler
}

func (h *Handler2) Execute(data *HandlerData, db *gorm.DB) error {
	// do something
	data.Handler2 = "handler2"
	return nil
}

func Test01(t *testing.T) {
	h1 := &Handler1{}
	h2 := &Handler2{}

	h1.SetNext(h2)

	data := &HandlerData{}
	db := &gorm.DB{}

	h1.Do(data, db)

	if data.Handler1 != "handler1" {
		t.Error("handler1 failed")
	}

	if data.Handler2 != "handler2" {
		t.Error("handler2 failed")
	}

}
