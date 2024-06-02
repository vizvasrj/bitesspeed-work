package myerror

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	pg "github.com/lib/pq"
)

type ErrorHandler func(*gin.Context) error

func HandleError(h ErrorHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h(c)
		if err != nil {
			// PrintDetailedError(err)
			switch err2 := err.(type) {
			case MyError:
				if err2.Custom {
					c.JSON(err2.Status(), gin.H{
						"message": err2.Error() + " " + err2.Message,
						"success": false,
					})
					return
				}
				c.AbortWithStatus(err2.Status())

			default:
				c.AbortWithStatus(500)
				return
			}

		}

	}
}

func PrintDetailedError(err error) {
	var builder strings.Builder
	for err != nil {
		switch e := err.(type) {
		case MyError:
			builder.WriteString(fmt.Sprintf("MyError: %#v ", e))
			err = e.Inner
		case *pg.Error:
			builder.WriteString(fmt.Sprintf("pg.Error: %#v ", e))
			err = nil // pg.Error does not wrap another error, so we stop here
		default:
			// If the error is a type we don't know how to handle, print it and stop
			builder.WriteString(fmt.Sprintf("Unknown error type: %#v ", e))
			err = nil
		}
	}
	log.Print(builder.String())
}
