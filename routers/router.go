package router

import (
	controller "gin-web/controllers"
	conts "gin-web/utils/constant"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	weather := r.Group(conts.WEATHER_URI)
	weather.GET(conts.WEATHER_NOW_URI, controller.WeatherNowHandler)
	r.Run(":8080")
}
