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

func NewPostService(kubernetesRepository ports.Deployment) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) CreateDefaultDeployment(name string, replicas *int32, image string) (interface{}, error) {

	//TODO: VERIFICAR SE O NOME DO DEPLOY JÁ EXISTE
	node, err := srv.kubernetesRepository.CreateDefaultDeployment(name, replicas, image)

	if err != nil {
		return nil, err
	}

	return node, nil
}

func (srv *Service) CreateFileDeployment(file *multipart.FileHeader) (interface{}, error) {
	//TODO: VERIFICAR SE O NOME DO DEPLOY JÁ EXISTE

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

	response, err := srv.kubernetesRepository.CreateFileDeployment(dep)

	if err != nil {
		return nil, err
	}

	return response, nil
}
