package ingresses

// import (
// 	"net/http"

// 	httpError "pod-chef-back-end/pkg/errors"

// 	"github.com/labstack/gommon/log"
// 	apiv1 "k8s.io/api/core/v1"
// )

// func (serviceHandler *KubernetesClient) PatchIngress() (interface{}, error) {
// 	list, err := serviceHandler.Clientset.NetworkingV1().Ingresses(apiv1.NamespaceDefault).Patch()
// 	if err != nil {
// 		log.Error(err)
// 		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
// 	}

// 	for _, ingress := range list.Items {
// 		response = append(response, Deploy{
// 			Name:      dep.Name,
// 			Namespace: dep.Namespace,
// 			Images:    dep.Spec.Template.Spec.Containers,
// 			Status:    dep.Status,
// 		})
// 	}

// 	return list, nil
// }
