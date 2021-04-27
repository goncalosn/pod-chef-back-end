package volumes

import (
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	VolumeServices ports.VolumeServices
}

func NewHTTPHandler(volumeService ports.VolumeServices) *HTTPHandler {
	return &HTTPHandler{
		VolumeServices: volumeService,
	}
}

func (h *HTTPHandler) GetVolumes(c echo.Context) error {
	response, err := h.VolumeServices.GetVolumes()

	if err != nil {
		kubernetesError := err.(*httpError.Error)
		return c.JSON(kubernetesError.Code, kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
