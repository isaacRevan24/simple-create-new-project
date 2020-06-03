package main

import (
	"fmt"

	"github.com/isaacRevan24/simple-create-new-project/project"
	"github.com/isaacRevan24/simple-create-new-project/types"
)

func main() {
	project := project.Project{}
	fmt.Println(project)
	fmt.Println("--------------------")
	// Create a indeterminate project
	indeterminateProject := types.Indeterminate{}
	// List of sections added
	sectionNames := []string{"Frontend", "Backend", "Base de datos", "Photoshop", "Publicidad"}
	// Start date of the project
	startDate := "24/4/2020"

	// Creating the project type and overwrite the empty instance
	indeterminateProject = indeterminateProject.CreateIndeterminateProjectStructure(sectionNames, startDate)
	fmt.Println(indeterminateProject)
	fmt.Println("--------------------")

	concreteProject := types.Concrete{}

	endDate := "24/4/2020"
	concreteProject = concreteProject.CreateConcreteProjectStructure(sectionNames, startDate, endDate)
	fmt.Println(concreteProject)
	fmt.Println("--------------------")

}

//TODO: Refactor all the project logic into a separate package
