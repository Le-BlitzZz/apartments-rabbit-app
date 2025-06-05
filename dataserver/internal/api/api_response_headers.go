package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddLoadedHeader(c *gin.Context, loaded bool) {
	c.Header("X-Loaded", strconv.FormatBool(loaded))
}
