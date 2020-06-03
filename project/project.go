package project

import (
	"github.com/isaacRevan24/simple-create-new-project/types"
)

// TODO: Make structureType just accept concrete and indeterminate type
// Can be concrete or indeterminate
type structureType interface{}

// Project properties definition
type Project struct {
	ProjectName      string
	Icon             uint8
	Manager          string
	Members          []string
	Description      string
	Visibility       bool
	ProjectStructure structureType
}

// GenerateProject return a project instance
func (p *Project) GenerateProject(typeStruct bool, projectName string, icon uint8, manager string, members []string, description string, visibility bool, sections []string, dates [2]string) Project {

	/*
		typeStruct
		true = concrete
		false = indeterminate

		dates
		dates[0] = start date
		dates[1] = end date
	*/
	// TODO: fix nil initialization
	var structure structureType
	startDate := dates[0]
	endDate := dates[1]

	if typeStruct {
		// indefinite
		std := types.Indeterminate{}
		std = std.CreateIndeterminateProjectStructure(sections, startDate)
		structure = structureType(std)
	} else {
		// concrete
		std := types.Concrete{}
		std = std.CreateConcreteProjectStructure(sections, startDate, endDate)
		structure = structureType(std)
	}

	// Project declaration and initialization
	project := Project{
		ProjectName:      projectName,
		Icon:             icon,
		Manager:          manager,
		Members:          members,
		Description:      description,
		Visibility:       visibility,
		ProjectStructure: structure,
	}

	return project
}
