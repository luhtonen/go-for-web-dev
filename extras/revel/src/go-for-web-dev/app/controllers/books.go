package controllers

import "github.com/revel/revel"

// Books struct
type Books struct {
	*revel.Controller
}

// Index function to get all books
func (c Books) Index() revel.Result {
	return c.RenderText("Getting all books\n")
}

// Get function to ge book by id
func (c Books) Get(id string) revel.Result {
	return c.RenderText("Getting book with ID " + id + "\n")
}

// Delete function to delete book by id
func (c Books) Delete(id string) revel.Result {
	return c.RenderText("Deleting book with ID " + id + "\n")
}
