package urlcontroller

import (
	"net/http"
	"web-project/models"
	"web-project/services/requester"
	"web-project/store"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	Store     store.Store
	Requester *requester.Requester
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
		urls, err := u.Store.GetUrls(userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(urls) >= 20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You can have only 20 urls"})
			return
		}
		url.UserName = userName
		id, err := u.Store.CreateUrl(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		url.Id = id
		u.Requester.AddUrl(url)
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
		url, err := u.Store.GetUrl(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, url)
	}
}
