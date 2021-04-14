package deployments

import ports "pod-chef-back-end/internal/core/ports"

type HTTPHandler struct {
	DeploymentServices ports.DeploymentServices
}

func NewHTTPHandler(deploymentService ports.DeploymentServices) *HTTPHandler {
	return &HTTPHandler{
		DeploymentServices: deploymentService,
	}
}
