package services

import "fmt"

// SendProjectSolicitude will send a solicitude to selected member that are
// friends added by the manager / creator of the project
func SendProjectSolicitude(membersRequest map[string]string, manager *map[string][2]string, project string, ID string) {

	/*
		manager[0] User Name
		manager[1] User ID
	*/
	for _, info := range *manager {
		fmt.Println("Solicitude sended by " + info[0] + " of the username" + info[1] + "at the project:" + project)
	}
	fmt.Println("secret ID: " + ID)

	for userName, email := range membersRequest {
		fmt.Println("Sended request to ", userName, " . At the email: ", email)
	}
}
