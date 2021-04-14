package fluentrestapi

import "net/http"

type HealthCheckLiveResponse struct {
	Status bool `json:"status,omitempty"`
}

type HealthCheckReadyResponse struct {
	Status bool `json:"status,omitempty"`
}

type HealthChecAllkResponse struct {
	Live  HealthCheckLiveResponse  `json:"live,omitempty"`
	Ready HealthCheckReadyResponse `json:"ready,omitempty"`
}

func (sc ServerContext) AddHealthCheck() ServerContext {
	handler := func(req Request) HandlerResponse {
		resp :=
			HandlerResponse{
				StatusCode:   http.StatusOK,
				JsonResponse: HealthChecAllkResponse{Live: HealthCheckLiveResponse{Status: true}},
			}

		return resp
	}
	sc.Get("/health", handler)
	return sc
}
