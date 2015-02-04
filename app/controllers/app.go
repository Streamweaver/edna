package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha Galaxy!"
	return c.Render(greeting)
}

func (c App) System(sysName string) revel.Result {
	c.Validation.Required(sysName).Message("A system name is required!")
	c.Validation.MinSize(sysName, 3).Message("You must enter at least 3 letters!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(sysName)
}
