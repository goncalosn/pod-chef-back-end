package deployments

import (
	"mime/multipart"
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"

	appsv1 "k8s.io/api/apps/v1"

	"github.com/labstack/gommon/log"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func NewPostService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
	}
}

func (srv *Service) CreateDefaultDeployment(name string, replicas *int32, image string) (interface{}, error) {
	nameExists, err := srv.k8DeploymentsRepository.CheckRepeatedDeployName(name, "default")
	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Deployment name already used."}
	}

	node, err := srv.k8DeploymentsRepository.CreateDefaultDeployment(name, replicas, image)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (srv *Service) CreateFileDeployment(file *multipart.FileHeader) (interface{}, error) {

	src, err := file.Open()
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	defer src.Close()

	yamlToJSON := yaml.NewYAMLToJSONDecoder(src)

	// future json object parsed from yaml file
	var dep *appsv1.Deployment
	if err := yamlToJSON.Decode(&dep); err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	//check namespace
	namespaces, err := srv.k8NamespacesRepository.GetNamespaces()
	if err != nil {
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	if dep.Namespace == "" {
		dep.Namespace = "default"
	}

	//verify corret namespace name
	if ok := contains(namespaces, dep.Namespace); !ok {
		return nil, &httpError.Error{Err: err, Code: http.StatusNotFound, Message: "Namespace does not exist"}
	}

	//check if deployment name already exists
	nameExists, err := srv.k8DeploymentsRepository.CheckRepeatedDeployName(dep.Name, dep.Namespace)
	if err != nil {
		return nil, err
	}

	if nameExists {
		//deployment name exists
		return nil, &httpError.Error{Err: err, Code: http.StatusConflict, Message: "Deployment name already used."}
	}

	response, err := srv.k8DeploymentsRepository.CreateFileDeployment(dep)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}
