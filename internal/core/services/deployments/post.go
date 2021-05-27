package deployments

import (
	ports "pod-chef-back-end/internal/core/ports"

	uuid "github.com/satori/go.uuid"
)

func NewPostService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace, k8IngressesRepository ports.Ingress) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
		k8IngressesRepository:   k8IngressesRepository,
	}
}

func (srv *Service) CreateDeployment(token string, replicas *int32, image string) (interface{}, error) {

	appUuid := uuid.NewV4().String() //generate uuid for this deployment

	namespaceUuid := "namespace-" + appUuid //generate name for namespace
	//create new namespace for user app
	_, err := srv.k8NamespacesRepository.CreateNamespace(namespaceUuid)
	if err != nil {
		return nil, err
	}

	deploymentUuid := "deployment-" + appUuid //generate name for the deployment

	//create deployment in namespace
	_, err = srv.k8DeploymentsRepository.CreateDeployment(namespaceUuid, deploymentUuid, replicas, image)
	if err != nil {
		//delete namespace
		_, _ = srv.k8NamespacesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	serviceUuid := "service-" + appUuid //generate name for the service
	//create service to expose the deployment
	_, err = srv.k8ServicesRepository.CreateClusterIPService(namespaceUuid, serviceUuid)
	if err != nil {
		//delete namespace
		_, _ = srv.k8NamespacesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	ingressUuid := "ingress-" + appUuid //generate name for the service
	//create ingress to expose the service
	_, err = srv.k8IngressesRepository.CreateIngress(namespaceUuid, ingressUuid, "app-"+appUuid)
	if err != nil {
		//delete namespace
		_, _ = srv.k8NamespacesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	//TODO: save on the database the name of the namespace and app

	//TODO: return link to app
	return deploymentUuid, err
}
