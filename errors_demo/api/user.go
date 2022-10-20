package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-grpc/errors_demo/pkg"
	"net/http"
)

// WriteResponse used to write an error and JSON data into response.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		coder := errors.ParseCoder(err)
		coder.Code()
		c.JSON(coder.HTTPStatus(), pkg.Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

		return
	}
	c.JSON(http.StatusOK, data)
}

func GetUser(c *gin.Context) {
	err := errors.New("test_error")

	WriteResponse(c, errors.WithCode(10010, err.Error()), nil)
}
