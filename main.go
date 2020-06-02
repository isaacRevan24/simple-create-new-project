package main

import (
	"fmt"

	"github.com/isaacRevan24/simple-create-new-project/types"
)

type project struct {
	ProjectName string
	Icon        uint8
	Manager     string
	Members     []string
	Description string
	Visibility  bool
	Tags        []string
}

func (p project) infoProject() {
	fmt.Println(p.Manager, "___", p.ProjectName)
}

func main() {
	project1 := project{
		ProjectName: "web1",
		Icon:        1,
		Manager:     "isaac",
		Members:     []string{"1", "2"},
		Description: "oh hi mark",
		Visibility:  true,
		Tags:        []string{"hi", "there"},
	}

	project1.infoProject()

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
