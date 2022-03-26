package role_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (res *roleDelivery) GetRoles(c *gin.Context) {
	response := res.usecase.GetRoles()
	// fmt.Printf("%+v", response)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
