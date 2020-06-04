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

	// Manager info
	manager := make(map[string][2]string)
	userID := "@isaac.revan"
	userInfo := [2]string{"Isaac Atencio", "isaac2498@email.com"}
	manager = map[string][2]string{userID: userInfo}

	// Arguments of GenerateProject
	typesctruct := false // true indefinite false concrete
	projectName := "Simple"
	icon := uint8(1)
	description := "Super cool project"
	project := project.Project{}
	sections := []string{"Backend", "Frontend", "DataBase"}
	startDate := "hoy"
	endDate := "Mañana"
	dates := [2]string{startDate, endDate}

	// Passing all the project argument and generate a complete formed project
	project = project.GenerateProject(typesctruct, projectName, icon, &manager, description, sections, dates)

	// fmt.Println(project)
	project.LogProject()

	// Members to send project solicitude
	membersRequest := map[string]string{"@eileen02": "eileen@email.com", "@carpedi": "arlete@outlook.com", "@mab.cad": "mab_cat@gmail.com"}

	// Sending project solicitude
	services.SendProjectSolicitude(membersRequest)

	// Execute time test
	t := time.Since(start)
	fmt.Println(t)
}

// TODO: See how to deal with the icon type
// TODO: Make optimization with pointers and concurrency
// TODO: Research how to do date formating
