package main

import (
	"log"
	"net/http"
	"time"

	"df-auth-service/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
}

// A sample user
var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func Login(c *gin.Context) {

	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	// Compare the user from the request, with the one we defined.
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Invalid login credentials.")
		return
	}
	token, err := CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

// Creates the jwt Access Token
func CreateToken(userid uint64) (string, error) {
	var err error
	access_secret := "uaishd8191dh98wn9d1n98w7dn1!@35s&!(ad"

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(access_secret))

	if err != nil {
		return "", err
	}

	return token, nil
}
