package app

import (
	"context"

	"github.com/booomch/demo-golang/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func handlerWrapper(
	ctr controller.Ctr,
	handler func(uc interface{}, ctx context.Context, req interface{}) (interface{}, error),
	requestDecoder func(c *fiber.Ctx) (interface{}, error),
	responseEncoder func(c *fiber.Ctx, response interface{}) error,
	uc interface{},
) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var reqDecoded interface{}
		var err error
		if requestDecoder != nil {
			reqDecoded, err = requestDecoder(c)
			if err != nil {
				return ctr.WrapError(err, c)
			}
		}
		res, err := handler(uc, ctr.PrepareContext(c), reqDecoded)
		if err != nil {
			return ctr.WrapError(err, c)
		}
		if responseEncoder != nil {
			return responseEncoder(c, res)
		}
		return c.JSON(res)
	}
}
