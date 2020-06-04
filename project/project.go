package project

import (
	"fmt"

	"github.com/isaacRevan24/simple-create-new-project/users"
)

// Can be concrete or indeterminate
type structureType interface{}

// Project properties definition
type Project struct {
	ProjectID        string
	ProjectName      string
	Icon             uint8
	Member           users.User
	Description      string
	ProjectStructure structureType
}

// LogProject print all the project information
func (p *Project) LogProject() {
	logMessage := `		Project information
	ProjectID: 		 %v
	ProjectName:     %v
	Icon: 			 %v
	Member: 		 %v
	Description: 	 %v
		Structure 
	%v`
	fmt.Printf(logMessage, p.ProjectID, p.ProjectName, p.Icon, p.Member, p.Description, p.ProjectStructure)
	fmt.Println("")
}

// GenerateProject return a project instance
func (p *Project) GenerateProject(typeStruct bool, projectName string, icon uint8, managerInfo *map[string][2]string, description string, sections []string, dates [2]string) Project {

	/*
		1. typeStruct bool
		false = concrete
		true = indeterminate

		2. projectName string

		3. icon uint8

		4. manager map[string]string

		5. description string

		6. sections []string

		7. dates [2]string
		dates[0] = start date
		dates[1] = end date

		return structureType
	*/

	// Initialize structure with the return type
	var structure structureType
	// start and end date as independent variables
	startDate := dates[0]
	endDate := dates[1]

	// If true generate a indeterminate structure else a concrete structure
	if typeStruct {
		// indefinite
		structure = generateIndeterminateStructure(sections, startDate)
	} else {
		// concrete
		structure = generateConcreteStructure(sections, startDate, endDate)
	}

	// Generate a member, in this case the only member is the manager
	member := users.GenerateMember(managerInfo)

	// Generate the project id
	ID := GenerateID(projectName)

	// Project declaration and initialization
	project := Project{
		ProjectID:        ID,
		ProjectName:      projectName,
		Icon:             icon,
		Member:           member,
		Description:      description,
		ProjectStructure: structure,
	}

	return project
}
