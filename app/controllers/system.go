package controllers

import (
	"github.com/revel/revel"
)

// System controls access to information about star systems.
type System struct {
	*revel.Controller
}

// Index returns a list view of systems
func (c System) Index() revel.Result {
	greeting := "Aloha Galaxy!"
	return c.Render(greeting)
}

// Show details of a specific system
func (c System) Show(id int) revel.Result {

	return c.Render(id)
}
