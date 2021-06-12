package controller

import (
	sconfiguration "github.com/booomch/demo-golang/internal/entities/_shared/configuration"
)

func ComposeEndpoints() []sconfiguration.EndpointConfiguration {

	res := []sconfiguration.EndpointConfiguration{
		{
			Path:            "/users/ping",
			Method:          "get",
			Handler:         PingServer,
			ResponseEncoder: sconfiguration.SuccessResponse,
		},
		{
			Path:    "/users/getusers",
			Method:  "get",
			Handler: GetUsers,
		},
	}

	return res
}
