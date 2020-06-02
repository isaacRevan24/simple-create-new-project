package types

/*
Base structure of a concrete project
Project how have a determined start date and a determines end date
*/

// Concrete type project
type Concrete struct {
	Indeterminate
	EndDate string
}

func getConcreteStructureSections(c *Concrete, sectionNames []string) sections {
	c.AddSection(sectionNames)
	return c.Structure
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

//TODO: Refactor this package in a single file for every type of project
