package main

import (
	"fmt"
	"time"

	"github.com/isaacRevan24/simple-create-new-project/project"
	"github.com/isaacRevan24/simple-create-new-project/services"
)

func main() {
	// Time testing
	start := time.Now()

	// Arguments of GenerateProject
	typesctruct := false // true indefinite false concrete
	projectName := "Simple"
	icon := uint8(1)
	manager := map[string]string{"@isaac.revan": "isaac@email.com"}
	description := "Super cool project"
	project := project.Project{}
	sections := []string{"Backend", "Frontend", "DataBase"}
	startDate := "hoy"
	endDate := "Mañana"
	dates := [2]string{startDate, endDate}

	// Passing all the project argument and generate a complete formed project
	project = project.GenerateProject(typesctruct, projectName, icon, manager, description, sections, dates)
	fmt.Println(project)

	// Members to send project solicitude
	membersRequest := map[string]string{"@eileen02": "eileen@email.com", "@carpedi": "arlete@outlook.com", "@mab.cad": "mab_cat@gmail.com"}

	// Sending project solicitude
	services.SendProjectSolicitude(membersRequest)

	t := time.Since(start)
	fmt.Println(t)
}

// TODO: Add to manager "User name" aka the specified name of the user like twitter
// TODO: Add a better log to the project information
// TODO: Organize all the variables in order
// TODO: See hoy to deal with the icon type
// TODO: Make optimization with pointers and concurrency
// TODO: Research how to do date formating
