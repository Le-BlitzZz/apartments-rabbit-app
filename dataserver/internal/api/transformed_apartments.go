package api

import (
	"fmt"
	"net/http"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity/search"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/form"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

func GetTransformedApartment(router *gin.RouterGroup) {
	router.GET("/transformed-apartment/:uuid", func(c *gin.Context) {
		if entity.TransformedApartmentsEmpty() {
			AddLoadedHeader(c, false)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		AddLoadedHeader(c, true)

		if a := entity.FindTransformedApartment(c.Param("uuid")); a == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("TransformedApartment with UUID %s not found", c.Param("uuid")),
			})
		} else {
			c.JSON(http.StatusOK, a)
		}
	})
}

func SearchTransformedApartments(router *gin.RouterGroup) {
	router.GET("/transformed-apartments", func(c *gin.Context) {
		if entity.TransformedApartmentsEmpty() {
			AddLoadedHeader(c, false)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		AddLoadedHeader(c, true)

		var frm form.SearchTransformedApartments
		if err := c.MustBindWith(&frm, binding.Form); err != nil {
			log.Error("Failed to bind form data", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
			return
		}

		result, err := search.TransformedApartments(frm)
		if err != nil {
			log.Error("Failed to search transformed apartments", "error", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to search transformed apartments"})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
