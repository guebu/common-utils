package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guebu/common-utils/errors"
)

func Respond(c *gin.Context, status int, body interface{}){
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}

	// Test

	// otherwise return body in JSON format
	c.JSON(status, body)
}

func RespondError(c *gin.Context, appErr *errors.ApplicationError){
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(appErr.AStatusCode, appErr)
		return
	}

	// otherwise return body in JSON format
	c.JSON(appErr.AStatusCode, appErr)
}

func GetJSONMessage(message string) string {
	jsonString := fmt.Sprintf(`{"message": "%s"}`, message)
	return jsonString
}

