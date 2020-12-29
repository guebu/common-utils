package main

import (
	"fmt"
	"github.com/guebu/common-utils/errors"
	"github.com/guebu/common-utils/logger"
)

func main() {
	fmt.Println("Test of go-utils library...")
	logger.Info("This is an Info...", "Status:Open", "App:Test")
	err := errors.NewNotFoundError("New not found error!")
	logger.Error("This is an error", err, "Status:Error")
}
