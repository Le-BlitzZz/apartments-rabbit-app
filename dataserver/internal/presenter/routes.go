package presenter

import (
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/api"
	"github.com/gin-gonic/gin"
)

func registerRoutes(API *gin.RouterGroup) {
	api.GetRawApartment(API)
	api.SearchRawApartments(API)

	api.GetTransformedApartment(API)
	api.SearchTransformedApartments(API)
}
