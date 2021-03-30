package k8

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodsHandler struct {
	Client *kubernetes.Clientset
}

//GetPodsByNamespace - GET - returns all the pods from the namespace
func (h *PodsHandler) GetPods(c echo.Context) error {
	c.Logger().Info("get pods request")

	//namespace := c.Param("namespace")

	pods, err := h.Client.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	// 		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return c.JSONPretty(http.StatusOK, pods, " ")
}
