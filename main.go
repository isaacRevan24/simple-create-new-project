package main

import (
	"fmt"

	"github.com/isaacRevan24/simple-create-new-project/project"
)

func main() {
	members := []string{"eileen", "arlette", "mab"}
	project := project.Project{}
	sections := []string{"Backend", "Frontend", "DataBase"}
	dates := [2]string{"hoy", "mañana"}

	project = project.GenerateProject(false, "Simple", 1, "isaac", members, "supe cool", true, sections, dates)
	fmt.Println(project)
}

//TODO: Refactor all the project logic into a separate package
//TODO: Add significative comments
//TODO: Improve main funcion order
