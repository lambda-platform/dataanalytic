package dataanalytic

import (
	"github.com/lambda-platform/lambda/config"
	"github.com/lambda-platform/dataanalytic/handler"
	"github.com/lambda-platform/dataanalytic/utils"
	"github.com/labstack/echo/v4"
	"github.com/lambda-platform/agent/agentMW"

)


func Set(e *echo.Echo){

	if config.Config.App.Migrate == "true"{
		utils.AutoMigrateSeed()
	}

	a :=e.Group("/analytics")
	a.Use(agentMW.IsLoggedInCookie)
	a.GET("/data", handler.AnalyticsData)
	a.POST("/pivot", handler.Pivot)
	/* ROUTES */


}

