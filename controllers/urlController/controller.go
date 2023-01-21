package urlcontroller

import (
	"net/http"
	"web-project/models"
	"web-project/store"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	Store store.Store
}

func (u *UrlController) CreateUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url models.Url
		err := c.ShouldBindJSON(&url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userName := c.MustGet("user_name").(string)
		url.UserName = userName
		err = u.Store.CreateUrl(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Url created successfully!"})
	}
}

func (u *UrlController) GetUrls() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.MustGet("user_name").(string)
		urls, err := u.Store.GetUrls(userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, urls)
	}
}

func (u *UrlController) GetUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userName := c.MustGet("user_name").(string)
		url, err := u.Store.GetUrl(userName, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, url)
	}
}
