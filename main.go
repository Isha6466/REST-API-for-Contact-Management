package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
  "gorm.io/gorm"
)

type Contact struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func main() {
	app := gofr.New()

	// Connect to DB (omitted for brevity)

	// Set Gofr database connection
	app.Data.DB = db

	// ... remaining API endpoints using app and db ...

	app.GET("/contacts", func(ctx *gofr.Context) {
		var contacts []Contact
		err := app.Data.DB.Find(&contacts)
		if err != nil {
			ctx.JSON(500, gofr.Error{Message: err.Error()})
			return
		}
		ctx.JSON(200, contacts)
	})

	app.GET("/contacts/:id", func(ctx *gofr.Context) {
		id := ctx.Param("id")
		var contact Contact
		err := app.Data.DB.First(&contact, "id = ?", id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(404, gofr.Error{Message: "Contact not found"})
			} else {
				ctx.JSON(500, gofr.Error{Message: err.Error()})
			}
			return
		}
		ctx.JSON(200, contact)
	})

	app.POST("/contacts", func(ctx *gofr.Context) {
		var contact Contact
		err := ctx.BindJSON(&contact)
		if err != nil {
			ctx.JSON(400, gofr.Error{Message: err.Error()})
			return
		}
		// Generate ID if needed
		// ...
		err = app.Data.DB.Create(&contact)
		if err != nil {
			ctx.JSON(500, gofr.Error{Message: err.Error()})
			return
		}
		ctx.JSON(201, contact)
	})

	app.PUT("/contacts/:id", func(ctx *gofr.Context) {
		id := ctx.Param("id")
		var contact Contact
		err := ctx.BindJSON(&contact)
		if err != nil {
			ctx.JSON(400, gofr.Error{Message: err.Error()})
			return
		}
		contact.ID = id // Ensure ID matches
		err = app.Data.DB.Updates(&contact)
		if err != nil {
			ctx.JSON(500, gofr.ErrorResponse{Message: err.Error()})
			return
		}
		ctx.JSON(200, contact)
	})

	app.DELETE("/contacts/:id", func(ctx *gofr.Context) {
		id := ctx.Param("id")
		err := app.Data.DB.Delete(&Contact{ID: id})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(404, gofr.Error{Message: "Contact not found"})
			} else {
				ctx.JSON(500, gofr.ErrorResponse{Message: err.Error()})
			}
		}
	})

	// Start Gofr server
	app.Start()
}

// Unit tests

func TestGetContacts(t *testing.T) {
	// Mock dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockGofrCtx := mock_gofr.NewMockContext(mockCtrl)
	mockDB := mock_sql.NewMockDB(mockCtrl)

	// Mock expected behavior
	contacts := []Contact{
		{ID: "1", Name: "John Doe"},
		{ID: "2", Name: "Jane Doe"},
	}
	mockDB.EXPECT().Find(&contacts).Return(nil) // Successful call

	// Set expectations for Gofr context
	mockGofrCtx.EXPECT().JSON(http.StatusOK, contacts).Times(1)

	// Execute handler with mocks
	app.GET("/contacts")(mockGofrCtx)

	// Verify expectations
	mockCtrl.Verify()
}
