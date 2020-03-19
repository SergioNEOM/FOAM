package database

import (
	"github.com/SergioNEOM/FOAM/models"
)

// GetUserByID get user info from DB
//
func (g *GDB) GetUserByID(id uint) (*models.User, error) {
	u := &models.User{}
	// ...
	return u, nil // !!!
}
