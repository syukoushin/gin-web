package controller

import (
	vo "gin-web/structs/vos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WeatherNowHandler(c *gin.Context) {

	w := &vo.WeatherNowVO{Temp: "37℃"}
	c.JSON(http.StatusOK, w)
}
