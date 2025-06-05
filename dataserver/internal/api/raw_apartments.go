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

func GetRawApartment(router *gin.RouterGroup) {
	router.GET("/raw-apartment/:uuid", func(c *gin.Context) {
		if entity.RawApartmentsEmpty() {
			AddLoadedHeader(c, false)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		AddLoadedHeader(c, true)

		if a := entity.FindRawApartment(c.Param("uuid")); a == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("RawApartment with UUID %s not found", c.Param("uuid")),
			})
		} else {
			c.JSON(http.StatusOK, a)
		}
	})
}

func SearchRawApartments(router *gin.RouterGroup) {
	router.GET("/raw-apartments", func(c *gin.Context) {
		if entity.RawApartmentsEmpty() {
			AddLoadedHeader(c, false)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		AddLoadedHeader(c, true)

		var frm form.SearchRawApartments
		if err := c.MustBindWith(&frm, binding.Form); err != nil {
			log.Error("Failed to bind form data", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
			return
		}

		result, err := search.RawApartments(frm)
		if err != nil {
			log.Error("Failed to search raw apartments", "error", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to search raw apartments"})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
