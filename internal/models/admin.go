package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
	"golang.org/x/crypto/bcrypt"
)

// Admin belongs to Credential
type Admin struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Credential   Credential
	CredentialID int
	FullName     string
}

func (admin *Admin) FromRequest(adminReq *dtos.AdminRequest) (*Admin, []error) {
	var cre Credential
	cre.RegName = adminReq.RegName

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminReq.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}
	cre.HashedPassword = string(hashedPassword)

	admin.Credential = cre
	admin.FullName = adminReq.FullName

	return admin, nil
}

func (admin *Admin) UpdateFromRequest(adminReq *dtos.AdminRequest) *Admin {
	admin.FullName = adminReq.FullName

	return admin
}

func (admin *Admin) ToResponse() *dtos.AdminResponse {
	return &dtos.AdminResponse{
		ID:       admin.ID,
		FullName: admin.FullName,
	}
}
