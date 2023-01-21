package usercontroller

import (
	"net/http"
	"time"

	"web-project/models"
	"web-project/store"
	"web-project/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Store store.Store
}

func (u *UserController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Password = hashed
		err = u.Store.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "created"})
	}
}

type loginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cred loginReq
		if err := c.ShouldBindJSON(&cred); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := u.Store.GetUser(cred.UserName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if utils.ValidatePassword(cred.Password, user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_name or password"})
			return
		}
		tokenStr, err := utils.CreateJwtToken(user.Username, time.Duration(time.Minute*30))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accessToken": tokenStr})
	}
}
