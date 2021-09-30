package profileutils

import "github.com/golang-jwt/jwt"

//PermissionGroup for grouping user permissions
type PermissionGroup struct {
}

//Permission the system allowed user permission
type Permission struct {
	Group       PermissionGroup `json:"group,omitempty"`
	Scope       string          `json:"scope,omitempty"`
	Description string          `json:"description,omitempty"`
}

//User is the model that defines auth users saved in database
type User struct {
	UID          string `json:"id,omitempty" bson:"_id"`
	SchoolID     string `json:"schoolID,omitempty" bson:"schoolID"`
	FirstName    string `json:"firstName,omitempty" validate:"required, min=2, max=100" bson:"firstName"`
	LastName     string `json:"lastName,omitempty" validate:"required, min=2, max=100" bson:"lastName"`
	Email        string `json:"email,omitempty" validate:"email, required"`
	Password     string `json:"password,omitempty"`
	Phone        string `json:"phone,omitempty" validate:"required"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken"`
	CreatedAt    int64  `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt,omitempty" bson:"updatedAt"`
	Active       bool   `json:"active,omitempty"`

	//Scopes are the permission scopes the user has
	Scopes []string `json:"scopes,omitempty"`
}

// SignedDetails are jwt details create by jwt
type SignedDetails struct {
	UID       string
	Email     string
	FirstName string
	LastName  string
	jwt.StandardClaims
}
