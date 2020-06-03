package main

import (
	"fmt"

	"github.com/isaacRevan24/simple-create-new-project/project"
)

func main() {
	// start := time.Now() Calculate start time

	// Arguments of GenerateProject
	typesctruct := true
	projectName := "Simple"
	icon := uint8(1)
	manager := "Isaac Atencio"
	members := []string{"eileen", "arlette", "mab"}
	description := "Super cool project"
	visibility := true
	project := project.Project{}
	sections := []string{"Backend", "Frontend", "DataBase"}
	startDate := "hoy"
	endDate := "Mañana"
	dates := [2]string{startDate, endDate}

	// Passing all the project argument and generate a complete formed project
	project = project.GenerateProject(typesctruct, projectName, icon, manager, members, description, visibility, sections, dates)
	fmt.Println(project)

	// t := time.Since(start) Get the time elapse since start time
	// fmt.Println(t) Print the elapse
}

// TODO: Add significative comments
// TODO: Make optimization with pointers and concurrency
// TODO: Make user package
// TODO: Research how to do date formating
