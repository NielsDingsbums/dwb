package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/nielsdingsbums/dwb/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Item)
// DB Table: Plural (items)
// Resource: Plural (Items)
// Path: Plural (/items)
// View Template Folder: Plural (/templates/items/)

// ItemsResource is the resource for the Item model
type ItemsResource struct {
	buffalo.Resource
}

// List gets all Items. This function is mapped to the path
// GET /items
func (v ItemsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	items := &models.Items{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Items from the DB
	if err := q.All(items); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, items))
}

// Show gets the data for one Item. This function is mapped to
// the path GET /items/{item_id}
func (v ItemsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Item
	item := &models.Item{}

	// To find the Item the parameter item_id is used.
	if err := tx.Find(item, c.Param("item_id")); err != nil {
		return c.Error(404, err)
	}

	var informationSet bool
	if item.Information.String != "" {
		informationSet = true
	} else {
		informationSet = false
	}
	c.Set("informationSet", informationSet)

	return c.Render(200, r.Auto(c, item))
}

// New renders the form for creating a new Item.
// This function is mapped to the path GET /items/new
func (v ItemsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Item{}))
}

// Create adds a Item to the DB. This function is mapped to the
// path POST /items
func (v ItemsResource) Create(c buffalo.Context) error {
	// Allocate an empty Item
	item := &models.Item{}

	// Bind item to the html form elements
	if err := c.Bind(item); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(item)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, item))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Item was created successfully")

	// and redirect to the items index page
	return c.Render(201, r.Auto(c, item))
}

// Edit renders a edit form for a Item. This function is
// mapped to the path GET /items/{item_id}/edit
func (v ItemsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Item
	item := &models.Item{}

	if err := tx.Find(item, c.Param("item_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, item))
}

// Update changes a Item in the DB. This function is mapped to
// the path PUT /items/{item_id}
func (v ItemsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Item
	item := &models.Item{}

	if err := tx.Find(item, c.Param("item_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Item to the html form elements
	if err := c.Bind(item); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(item)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, item))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Item was updated successfully")

	// and redirect to the items index page
	return c.Render(200, r.Auto(c, item))
}

// Destroy deletes a Item from the DB. This function is mapped
// to the path DELETE /items/{item_id}
func (v ItemsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Item
	item := &models.Item{}

	// To find the Item the parameter item_id is used.
	if err := tx.Find(item, c.Param("item_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(item); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Item was destroyed successfully")

	// Redirect to the items index page
	return c.Render(200, r.Auto(c, item))
}