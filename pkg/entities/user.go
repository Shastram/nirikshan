package entities

import (
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"nirikshan-backend/pkg/utils"
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty" json:"_id"`
	UserName  string    `bson:"username,omitempty" json:"username"`
	Password  string    `bson:"password,omitempty" json:"password,omitempty"`
	Email     string    `bson:"email,omitempty" json:"email,omitempty"`
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt string    `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// UserPassword defines structure for password related requests
type UserPassword struct {
	Username    string `json:"username,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

type UserResponse struct {
	JwtToken string `json:"jwt_token"`
	User     *User  `json:"user"`
}

// SanitizedUser returns the user object without sensitive information
func (user *User) SanitizedUser() *User {
	user.Password = ""
	return user
}

// GetSignedJWT generates the JWT Token for the user object
func (user *User) GetSignedJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(utils.JWTExpiryDuration)).Unix()
	claims["iat"] = time.Now().Unix()
	tokenString, err := token.SignedString([]byte(utils.JwtSecret))
	if err != nil {
		logrus.Info(err)
		return "", err
	}
	return tokenString, nil
}
