package services

import "fmt"

// SendProjectSolicitude print the @user.name with their respective email
func SendProjectSolicitude(membersRequest map[string]string) {
	for userName, email := range membersRequest {
		fmt.Println("Sended request to ", userName, " . At the email: ", email)
	}
}

// TODO: Add to SendProjectSolicitude manager name and project name and project ID at the request
