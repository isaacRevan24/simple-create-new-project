package main

import (
	"fmt"
	"time"

	"github.com/isaacRevan24/simple-create-new-project/project"
)

func main() {
	start := time.Now()

	// Arguments of GenerateProject
	typesctruct := false // true indefinite false concrete
	projectName := "Simple"
	icon := uint8(1)
	manager := "Isaac Atencio"
	memberRequest := map[string]string{"@eileen02": "eileen@email.com", "@carpedi": "arlete@outlook.com", "@mab.cad": "mab_cat@gmail.com"}
	description := "Super cool project"
	visibility := true
	project := project.Project{}
	sections := []string{"Backend", "Frontend", "DataBase"}
	startDate := "hoy"
	endDate := "Mañana"
	dates := [2]string{startDate, endDate}

	// Passing all the project argument and generate a complete formed project
	project = project.GenerateProject(typesctruct, projectName, icon, manager, memberRequest, description, visibility, sections, dates)
	fmt.Println(project)

	t := time.Since(start)
	fmt.Println(t)
}

// TODO: Make optimization with pointers and concurrency
// TODO: Make user package
// TODO: Research how to do date formating
