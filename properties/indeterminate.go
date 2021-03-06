package properties

/*
Base structure of a indeterminate project
Project how have a defined start date but not a defined end date
*/

// sections list
type sections []string

// Indeterminate type project
type Indeterminate struct {
	ProjectType string
	Structure   sections
	StartDate   string
}

// AddSection add a new section to the structure
func (i *Indeterminate) AddSection(sectionNames []string) {
	// Initialice the structure map
	sizeSectionNames := len(sectionNames) // size ot the slice

	if sizeSectionNames == 0 {
		// Default value of "General":100 if no section is specified
		i.Structure = append(i.Structure, "General")
	} else {
		/*
		 Assign the custom section name to the structure of the project with
		 the corresponding percentage calculate ( 100/amount of sections )
		 generatin a map like this ["a":33, "b":33, "c":33]
		*/

		// Adding the custom section name with his specific percentage to the project structure
		for _, value := range sectionNames {
			i.Structure = append(i.Structure, value)
		}
	}
}

// Return the indeterminate structure with the sections added and the %
func getIndeterminateStructureSections(i *Indeterminate, sectionNames []string) sections {
	i.AddSection(sectionNames)
	return i.Structure
}

// CreateIndeterminateProjectStructure return a complete structure instance of a project structure
func (i *Indeterminate) CreateIndeterminateProjectStructure(sectionNames []string, startDate string) Indeterminate {
	// Generate the project structure with all sections and %
	structure := getIndeterminateStructureSections(i, sectionNames)

	// Indeterminate type declaration and initialization
	indeterminateProject := Indeterminate{
		ProjectType: "Indeterminate",
		Structure:   structure,
		StartDate:   startDate,
	}

	return indeterminateProject
}
