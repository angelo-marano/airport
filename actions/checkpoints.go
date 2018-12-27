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
// Model: Singular (Checkpoint)
// DB Table: Plural (checkpoints)
// Resource: Plural (Checkpoints)
// Path: Plural (/checkpoints)
// View Template Folder: Plural (/templates/checkpoints/)

// CheckpointsResource is the resource for the Checkpoint model
type CheckpointsResource struct {
	buffalo.Resource
}

// List gets all Checkpoints. This function is mapped to the path
// GET /checkpoints
func (v CheckpointsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	checkpoints := &models.Checkpoints{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Checkpoints from the DB
	if err := q.All(checkpoints); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, checkpoints))
}

// Show gets the data for one Checkpoint. This function is mapped to
// the path GET /checkpoints/{checkpoint_id}
func (v CheckpointsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Checkpoint
	checkpoint := &models.Checkpoint{}

	// To find the Checkpoint the parameter checkpoint_id is used.
	if err := tx.Find(checkpoint, c.Param("checkpoint_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, checkpoint))
}

// New renders the form for creating a new Checkpoint.
// This function is mapped to the path GET /checkpoints/new
func (v CheckpointsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Checkpoint{}))
}

// Create adds a Checkpoint to the DB. This function is mapped to the
// path POST /checkpoints
func (v CheckpointsResource) Create(c buffalo.Context) error {
	// Allocate an empty Checkpoint
	checkpoint := &models.Checkpoint{}

	// Bind checkpoint to the html form elements
	if err := c.Bind(checkpoint); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(checkpoint)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, checkpoint))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Checkpoint was created successfully")

	// and redirect to the checkpoints index page
	return c.Render(201, r.Auto(c, checkpoint))
}

// Edit renders a edit form for a Checkpoint. This function is
// mapped to the path GET /checkpoints/{checkpoint_id}/edit
func (v CheckpointsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Checkpoint
	checkpoint := &models.Checkpoint{}

	if err := tx.Find(checkpoint, c.Param("checkpoint_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, checkpoint))
}

// Update changes a Checkpoint in the DB. This function is mapped to
// the path PUT /checkpoints/{checkpoint_id}
func (v CheckpointsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Checkpoint
	checkpoint := &models.Checkpoint{}

	if err := tx.Find(checkpoint, c.Param("checkpoint_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Checkpoint to the html form elements
	if err := c.Bind(checkpoint); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(checkpoint)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, checkpoint))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Checkpoint was updated successfully")

	// and redirect to the checkpoints index page
	return c.Render(200, r.Auto(c, checkpoint))
}

// Destroy deletes a Checkpoint from the DB. This function is mapped
// to the path DELETE /checkpoints/{checkpoint_id}
func (v CheckpointsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Checkpoint
	checkpoint := &models.Checkpoint{}

	// To find the Checkpoint the parameter checkpoint_id is used.
	if err := tx.Find(checkpoint, c.Param("checkpoint_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(checkpoint); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Checkpoint was destroyed successfully")

	// Redirect to the checkpoints index page
	return c.Render(200, r.Auto(c, checkpoint))
}