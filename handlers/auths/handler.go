package auths

import ports "pod-chef-back-end/internal/core/ports"

type HTTPHandler struct {
	UserServices ports.UserServices
}

func NewHTTPHandler(userService ports.UserServices) *HTTPHandler {
	return &HTTPHandler{
		UserServices: userService,
	}
}
