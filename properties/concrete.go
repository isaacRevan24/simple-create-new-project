package properties

/*
Base structure of a concrete project
Project how have a determined start date and a determines end date
*/

// Concrete type project
type Concrete struct {
	Indeterminate
	EndDate string
}

// Return the concrete structure with the sections added and the %
func getConcreteStructureSections(c *Concrete, sectionNames []string) sections {
	c.AddSection(sectionNames)
	return c.Structure
}

// CreateConcreteProjectStructure return a complete structure instance of a project structure
func (c *Concrete) CreateConcreteProjectStructure(sectionNames []string, startDate string, endDate string) Concrete {
	// Generate the project structure with all sections and %
	structure := getConcreteStructureSections(c, sectionNames)

	// Embedded indterminate in concrete type
	indeterminate := Indeterminate{
		ProjectType: "Concrete",
		Structure:   structure,
		StartDate:   startDate,
	}

	// Concrete type declaration and initialization
	concreteType := Concrete{
		Indeterminate: indeterminate,
		EndDate:       endDate,
	}

	return concreteType
}
