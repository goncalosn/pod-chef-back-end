package kubernetes

import (
	"net/http"
	"pod-chef-back-end/pkg"

	uuid "github.com/satori/go.uuid"
)

//CreateDeployment service responsible for creating a deployment inside a new namespace
func (srv *Service) CreateDeployment(email string, replicas *int32, image string) (interface{}, error) {

	appUUID := uuid.NewV4().String() //generate uuid for this deployment

	namespaceUUID := "namespace-" + appUUID //generate name for namespace

	//call driven adapter responsible for creating namespaces inside the kubernetes cluster
	_, err := srv.kubernetesRepository.CreateNamespace(namespaceUUID)
	if err != nil {
		//return error from the kubernetes repository method
		return nil, err
	}

	deploymentUUID := "deployment-" + appUUID //generate name for the deployment

	//call driven adapter responsible for creating deployments inside the kubernetes cluster
	_, err = srv.kubernetesRepository.CreateDeployment(namespaceUUID, deploymentUUID, replicas, image)
	if err != nil {
		//creation of the deployment went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	serviceUUID := "service-" + appUUID //generate name for the service
	//create service to expose the deployment
	_, err = srv.kubernetesRepository.CreateClusterIPService(namespaceUUID, serviceUUID)
	if err != nil {
		//creation of the service went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	ingressUUID := "ingress-" + appUUID //generate name for the service
	//create ingress to expose the service
	_, err = srv.kubernetesRepository.CreateIngress(namespaceUUID, ingressUUID, "app-"+appUUID)
	if err != nil {
		//creation of the ingress went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	srv.mongoRepository.InsertDeployment(appUUID, email, image)
	if err != nil {
		//return error from the mongo repository method
		return nil, err
	}

	//return app uuid
	return appUUID, err
}

//GetDeploymentsByUser service responsible for getting all deployments inside a namespace
func (srv *Service) GetDeploymentsByUser(email string) (interface{}, error) {
	//call driven adapter responsible for getting all deployments from database
	response, err := srv.mongoRepository.GetAllDeploymentsByUser(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetDeploymentByUserAndName service responsible for getting a deployment
func (srv *Service) GetDeploymentByUserAndName(email string, uuid string) (interface{}, error) {
	//call driven adapter responsible for getting a deployment from the database
	response, err := srv.mongoRepository.GetDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//verify if the user requesting the deployment, it's the deployment's creator
	if email != response.User {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusForbidden, Message: "Cannot get another user's deployment"}
	}

	return response, nil
}

//DeleteDeploymentByUserAndUUID service responsible for deleting a deployment
func (srv *Service) DeleteDeploymentByUserAndUUID(email string, uuid string) (interface{}, error) {

	//call driven adapter responsible for getting a deployment from the database
	deployment, err := srv.mongoRepository.GetDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//verify if the user deleting the deployment, it's the deployment's creator
	if email != deployment.User {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusForbidden, Message: "Cannot delete another user's deployment"}
	}

	//call driven adapter responsible for getting a deployment from the database
	_, err = srv.kubernetesRepository.DeleteNamespace("namespace-" + uuid)
	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//call driven adapter responsible for getting a deployment from the database
	response, err := srv.mongoRepository.DeleteDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
