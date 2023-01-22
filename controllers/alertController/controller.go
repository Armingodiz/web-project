package alertcontroller

import (
	"net/http"
	"web-project/store"

	"github.com/gin-gonic/gin"
)

type AlertController struct {
	Store store.Store
}

func (al *AlertController) GetAlerts() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlId := c.Param("id")
		requests, err := al.Store.GetAlerts(urlId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, requests)
	}
}
