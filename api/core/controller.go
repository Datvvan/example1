package core

import (
	"net/http"

	"github.com/datvvan/sample1/util"
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) Detail(c echo.Context) error {
	resp := &ServiceResp{
		Text:  "Body of microservice 1",
		Image: "https://storage.googleapis.com/gcp-obj-udestiny-dev/udestiny/1715570454/bitkub.png",
	}
	return c.JSON(http.StatusOK, util.SuccessResponse(resp))
}
