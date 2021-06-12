package controller

import (
	"net/http"
	"reflect"

	"github.com/booomch/demo-golang/pkg/codes"
	"github.com/booomch/demo-golang/pkg/httperr"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (ctr *Ctr) WrapError(err error, c *fiber.Ctx) error {
	if reflect.TypeOf(err).String() == "*httperr.HTTPErr" {
		httpErr := err.(*httperr.HTTPErr)
		return httpErr.Send(c)
	}

	logrus.Error("Unhandled error ", "error: ", err, c.Request())
	return ErrInternal.SetDetail(err).Send(c)
}

var (
	ErrInternal  = httperr.New(codes.Omit, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	ErrParseBody = httperr.New(codes.Omit, http.StatusBadRequest, "Failed to parse body")
)
