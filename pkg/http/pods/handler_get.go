package pods

import (
	"net/http"
	. "pod-chef-back-end/pkg/domain/pods"
	httpError "pod-chef-back-end/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PodHandler struct {
	PodInteractor PodInteractor
}

//GetPodsByNodeAndNamespace - GET - returns all the pods from the namespace
func (h *PodHandler) GetPodsByNodeAndNamespace(c echo.Context) error {
	log.Info("GetPodsByNodeAndNamespace request")

	namespace := c.FormValue("namespace")
	node := c.FormValue("node")

	if namespace == "" || node == "" {
		return c.JSON(http.StatusBadRequest, "invalid form")
	}

	response, err := h.PodInteractor.GetPodsByNodeAndNamespaceInteractor(node, namespace)

	if err != nil {
		kubernetesError := err.(httpError.KubernetesError)
		return c.JSON(kubernetesError.GetStatus(), kubernetesError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
