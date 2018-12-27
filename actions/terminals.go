package actions

import (
	"github.com/angelo-marano/airport/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Terminal)
// DB Table: Plural (terminals)
// Resource: Plural (Terminals)
// Path: Plural (/terminals)
// View Template Folder: Plural (/templates/terminals/)

// TerminalsResource is the resource for the Terminal model
type TerminalsResource struct {
	buffalo.Resource
}

// List gets all Terminals. This function is mapped to the path
// GET /terminals
func (v TerminalsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	terminals := &models.Terminals{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Terminals from the DB
	if err := q.All(terminals); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, terminals))
}

// Show gets the data for one Terminal. This function is mapped to
// the path GET /terminals/{terminal_id}
func (v TerminalsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Terminal
	terminal := &models.Terminal{}

	// To find the Terminal the parameter terminal_id is used.
	if err := tx.Find(terminal, c.Param("terminal_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, terminal))
}

// New renders the form for creating a new Terminal.
// This function is mapped to the path GET /terminals/new
func (v TerminalsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Terminal{}))
}

// Create adds a Terminal to the DB. This function is mapped to the
// path POST /terminals
func (v TerminalsResource) Create(c buffalo.Context) error {
	// Allocate an empty Terminal
	terminal := &models.Terminal{}

	// Bind terminal to the html form elements
	if err := c.Bind(terminal); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(terminal)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, terminal))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Terminal was created successfully")

	// and redirect to the terminals index page
	return c.Render(201, r.Auto(c, terminal))
}

// Edit renders a edit form for a Terminal. This function is
// mapped to the path GET /terminals/{terminal_id}/edit
func (v TerminalsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Terminal
	terminal := &models.Terminal{}

	if err := tx.Find(terminal, c.Param("terminal_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, terminal))
}

// Update changes a Terminal in the DB. This function is mapped to
// the path PUT /terminals/{terminal_id}
func (v TerminalsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Terminal
	terminal := &models.Terminal{}

	if err := tx.Find(terminal, c.Param("terminal_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Terminal to the html form elements
	if err := c.Bind(terminal); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(terminal)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, terminal))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Terminal was updated successfully")

	// and redirect to the terminals index page
	return c.Render(200, r.Auto(c, terminal))
}

// Destroy deletes a Terminal from the DB. This function is mapped
// to the path DELETE /terminals/{terminal_id}
func (v TerminalsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Terminal
	terminal := &models.Terminal{}

	// To find the Terminal the parameter terminal_id is used.
	if err := tx.Find(terminal, c.Param("terminal_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(terminal); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Terminal was destroyed successfully")

	// Redirect to the terminals index page
	return c.Render(200, r.Auto(c, terminal))
}