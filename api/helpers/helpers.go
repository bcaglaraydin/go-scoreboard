package helpers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func LogError(msg string, err error) {
	if err != nil {
		log.Fatal(msg+"\n", err)
		os.Exit(2)
	}
}

func ResponseError(ctx *fiber.Ctx, status int, msg string) {
	er := HTTPError{
		Code:    status,
		Message: msg,
	}
	ctx.JSON(er)
}

type HTTPError struct {
	Code    int    `json:"status"`
	Message string `json:"message" `
}
