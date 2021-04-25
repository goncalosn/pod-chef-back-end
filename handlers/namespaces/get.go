package namespaces

import (
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	NamespaceServices ports.NamespaceServices
}

func NewHTTPHandler(namespaceService ports.NamespaceServices) *HTTPHandler {
	return &HTTPHandler{
		NamespaceServices: namespaceService,
	}
}

func (h *HTTPHandler) GetNamespaces(c echo.Context) error {
	response, err := h.NamespaceServices.GetNamespaces()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
