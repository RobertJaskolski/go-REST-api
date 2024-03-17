package models

import (
	"time"
)

type UserRoleEnum string

const (
	SuperAdminRole UserRoleEnum = "super_admin"
	OwnerRole      UserRoleEnum = "owner"
	T3AdminRole    UserRoleEnum = "t3_admin"
	AdminRole      UserRoleEnum = "admin"
	UserRole       UserRoleEnum = "user"
	ViewerRole     UserRoleEnum = "viewer"
	SupportRole    UserRoleEnum = "support"
)

type User struct {
	ID         int            `json:"id"`
	Email      string         `json:"email"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	TimeZone   NullableString `json:"time_zone"`
	Mobile     NullableString `json:"mobile"`
	Role       string         `json:"role"`
	IsActive   bool           `json:"is_active"`
	CreatedAt  time.Time      `json:"created_at"`
	ModifiedAt time.Time      `json:"modified_at"`
}

type CreateUserDTO struct {
	Email     string         `json:"email" validate:"required,email,max=50"`
	FirstName string         `json:"first_name" validate:"required,max=50"`
	LastName  string         `json:"last_name" validate:"required,max=50"`
	TimeZone  NullableString `json:"time_zone"`
	Mobile    NullableString `json:"mobile"`
	IsActive  bool           `json:"is_active" validate:"required"`
	Password  string         `json:"password"`
	Role      string         `json:"role" validate:"required,oneof=super_admin owner t3_admin admin user viewer support"`
}
