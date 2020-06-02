package types

import "fmt"

// sections list
type sections map[string]float32

// Indeterminate type project
type Indeterminate struct {
	//Base structure of a indeterminate project
	//Project how have a start but not a end date
	ProjectType string
	Structure   sections
	StartDate   string
}

// Concrete type project
type Concrete struct {
	//Base structure of a concrete project
	//Project how have a start and end date
	Indeterminate
	EndDate string
}

//ProjectInformation Show the current information of the Indeterminate project
func (i *Indeterminate) ProjectInformation() {
	fmt.Println(i.ProjectType, i.Structure, i.StartDate)
}

// AddSection add a new section to the structure
func (i *Indeterminate) AddSection(sectionNames []string) {
	// Initialice the structure map
	i.Structure = make(sections)
	sizeSectionNames := len(sectionNames) // size ot the slice

	if sizeSectionNames == 0 {
		// Default value of "General":100 if no section is specified
		i.Structure["General"] = float32(100)
	} else {
		/*
		 Assign the custom section name to the structure of the project with
		 the corresponding percentage calculate ( 100/amount of sections )
		 generatin a map like this ["a":33, "b":33, "c":33]
		*/

		// Calculating the percentage of each project over 100
		percentage := float32(100 / sizeSectionNames)

		// Adding the custom section name with his specific percentage to the project structure
		for _, value := range sectionNames {
			i.Structure[value] = percentage
		}
	}
}

func getIndeterminateStructureSections(i *Indeterminate, sectionNames []string) sections {
	i.AddSection(sectionNames)
	return i.Structure
}

func getConcreteStructureSections(c *Concrete, sectionNames []string) sections {
	c.AddSection(sectionNames)
	return c.Structure
}

// CreateIndeterminateProjectStructure return a complete structure instance of a project structure
func (i *Indeterminate) CreateIndeterminateProjectStructure(sectionNames []string, startDate string) Indeterminate {
	structure := getIndeterminateStructureSections(i, sectionNames)

	indeterminateProject := Indeterminate{
		ProjectType: "Indeterminate",
		Structure:   structure,
		StartDate:   startDate,
	}

	return indeterminateProject
}

// CreateConcreteProjectStructure return a complete structure instance of a project structure
func (c *Concrete) CreateConcreteProjectStructure(sectionNames []string, startDate string, endDate string) Concrete {
	structure := getConcreteStructureSections(c, sectionNames)

	indeterminate := Indeterminate{
		ProjectType: "Concrete",
		Structure:   structure,
		StartDate:   startDate,
	}

	concreteType := Concrete{
		Indeterminate: indeterminate,
		EndDate:       endDate,
	}

	return concreteType
}
