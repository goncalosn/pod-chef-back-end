package kubernetes

import (
	"net/http"
	"pod-chef-back-end/pkg"
)

//CreateDeployment service responsible for creating a deployment inside a new namespace
func (srv *Service) CreateDeployment(id string, role string, deployName string, replicas *int32, image string, containerPort int32) (interface{}, error) {
	//check the number os deployments by the user
	deployments, err := srv.mongoRepository.GetDeploymentsFromUser(id)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//check if user has already the max number os deployments
	if role != "admin" {
		if len(deployments) > 3 {
			//return a custom error
			return nil, &pkg.Error{Err: err, Code: http.StatusBadRequest, Message: "Max number 3 deployments permited per user"}
		}
	}

	namespaceUUID := "namespace-" + deployName //generate name for namespace

	//call driven adapter responsible for creating namespaces inside the kubernetes cluster
	_, err = srv.kubernetesRepository.CreateNamespace(namespaceUUID)
	if err != nil {
		//return error from the kubernetes repository method
		return nil, err
	}

	deploymentUUID := "deployment-" + deployName //generate name for the deployment

	//call driven adapter responsible for creating deployments inside the kubernetes cluster
	_, err = srv.kubernetesRepository.CreateDeployment(namespaceUUID, deploymentUUID, replicas, image, containerPort)
	if err != nil {
		//creation of the deployment went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	serviceUUID := "service-" + deployName //generate name for the service
	//create service to expose the deployment
	_, err = srv.kubernetesRepository.CreateClusterIPService(namespaceUUID, serviceUUID, containerPort)
	if err != nil {
		//creation of the service went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, deperr := srv.kubernetesRepository.DeleteNamespace(namespaceUUID)
		if deperr != nil {
			return nil, deperr
		}
		//return error from the kubernetes repository method
		return nil, err
	}

	ingressUUID := "ingress-" + deployName //generate name for the service
	//create ingress to expose the service
	_, err = srv.kubernetesRepository.CreateIngress(namespaceUUID, ingressUUID, deployName)
	if err != nil {
		//creation of the ingress went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, deperr := srv.kubernetesRepository.DeleteNamespace(namespaceUUID)
		if deperr != nil {
			return nil, deperr
		}
		//return error from the kubernetes repository method
		return nil, err
	}

	srv.mongoRepository.InsertDeployment(deployName, id, image)
	if err != nil {
		//delete namespace
		_, deperr := srv.kubernetesRepository.DeleteNamespace(namespaceUUID)
		if deperr != nil {
			return nil, deperr
		}
		//return error from the mongo repository method
		return nil, err
	}

	//return app uuid
	return deployName, nil
}

//GetAllDeployments service responsible for getting all deployments inside a namespace
func (srv *Service) GetAllDeployments() (interface{}, error) {
	//call driven adapter responsible for getting all deployments from database
	response, err := srv.mongoRepository.GetAllDeployments()

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetDeploymentsByUser service responsible for getting all deployments inside a namespace
func (srv *Service) GetDeploymentsByUser(id string) (interface{}, error) {
	//call driven adapter responsible for getting all deployments from database
	response, err := srv.mongoRepository.GetDeploymentsFromUser(id)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetDeploymentByUserAndName service responsible for getting a deployment
func (srv *Service) GetDeploymentByUserAndName(id string, uuid string) (interface{}, error) {
	//call driven adapter responsible for getting a deployment from the database
	response, err := srv.mongoRepository.GetDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//verify if the user requesting the deployment, it's the deployment's creator
	if id != response.User {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusForbidden, Message: "Cannot get another user's deployment"}
	}

	return response, nil
}

//DeleteDeploymentByUserAndUUID service responsible for deleting a deployment
func (srv *Service) DeleteDeploymentByUserAndUUID(id string, uuid string) (*string, error) {

	//call driven adapter responsible for getting a deployment from the database
	deployment, err := srv.mongoRepository.GetDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//verify if the user deleting the deployment, it's the deployment's creator
	if id != deployment.User {
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
	_, err = srv.mongoRepository.DeleteDeploymentByUUID(uuid)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	message := "Deployment deleted sucessfully"

	return &message, nil
}
