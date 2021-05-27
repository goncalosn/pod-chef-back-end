package deployments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	ports "pod-chef-back-end/internal/core/ports"
	httpError "pod-chef-back-end/pkg/errors"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/labstack/gommon/log"
)

func NewPostService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace, k8IngressesRepository ports.Ingress) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
		k8IngressesRepository:   k8IngressesRepository,
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

	deploy, err := srv.k8DeploymentsRepository.CreateDefaultDeployment(name, replicas, image)
	if err != nil {
		return nil, err
	}

	return deploy, nil
}

func (srv *Service) CreateFileDeployment(file *multipart.FileHeader) (interface{}, error) {

	src, err := file.Open()
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	defer src.Close()

	var buffer bytes.Buffer

	_, err = buffer.ReadFrom(src)
	if err != nil {
		log.Error(err)
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	content := buffer.String()
	deployParts := strings.Split(content, "---")

	var responses []interface{}

	for _, part := range deployParts {
		buffer.Reset()
		buffer.WriteString(part)

		parsedJSON, err := yaml.ToJSON(buffer.Bytes())
		if err != nil {
			log.Error(err)
			return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
		}

		response, err := selectKind(srv, parsedJSON)
		if err != nil {
			log.Error(err)
			return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func contains(array []string, str string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

func selectKind(srv *Service, parsedJSON []byte) (interface{}, error) {
	//check namespace
	namespaces, err := srv.k8NamespacesRepository.GetNamespaces()
	if err != nil {
		return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	formatedJSON := strings.ToLower(string(parsedJSON))

	if strings.Contains(formatedJSON, "\"kind\":\"deployment\"") {
		var dep *appsv1.Deployment
		if err := json.Unmarshal(parsedJSON, &dep); err != nil {
			log.Error(err)
			return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
		}

		if dep.Namespace == "" {
			dep.Namespace = "default"
		}

		//verify corret namespace name
		if ok := contains(namespaces, dep.Namespace); !ok {
			return nil, &httpError.Error{Err: err, Code: http.StatusNotFound, Message: "Namespace does not exist"}
		}

		response, err := srv.k8DeploymentsRepository.CreateFileDeployment(dep)
		if err != nil {
			return nil, err
		}
		return response, nil
	} else if strings.Contains(formatedJSON, "\"kind\":\"service\"") {
		var serv *v1.Service
		if err := json.Unmarshal(parsedJSON, &serv); err != nil {
			log.Error(err)
			return nil, &httpError.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
		}

		if serv.Namespace == "" {
			serv.Namespace = "default"
		}

		//verify corret namespace name
		if ok := contains(namespaces, serv.Namespace); !ok {
			return nil, &httpError.Error{Err: err, Code: http.StatusNotFound, Message: "Namespace does not exist"}
		}

		response, err := srv.k8ServicesRepository.CreateService(serv)
		if err != nil {
			return nil, err
		}
		return response, nil
	} else {
		fmt.Println("unknown deploy")
		return nil, nil
	}
}
