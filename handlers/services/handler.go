package services

import ports "pod-chef-back-end/internal/core/ports"

type HTTPHandler struct {
	ServiceServices ports.ServiceServices
}

func NewHTTPHandler(serviceService ports.ServiceServices) *HTTPHandler {
	return &HTTPHandler{
		ServiceServices: serviceService,
	}
}
