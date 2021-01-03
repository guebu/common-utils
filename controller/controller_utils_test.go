package controller

import (
	"fmt"
	"github.com/guebu/common-utils/logger"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetJSONMessage(t *testing.T) {
	jsonString := GetJSONMessage("Ping")
	compareString := fmt.Sprintf(`{"message": "%s"}`, "Ping")

	logger.Info("JSON String:" + jsonString)
	logger.Info("Compare String:" + compareString)


	assert.EqualValues(t,jsonString, compareString)
}
