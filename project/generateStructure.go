package project

import "github.com/isaacRevan24/simple-create-new-project/properties"

// generateIndeterminateStructure generate the Indeterminate project structure
func generateIndeterminateStructure(sections []string, startDate string) structureType {
	std := properties.Indeterminate{}
	std = std.CreateIndeterminateProjectStructure(sections, startDate)
	return structureType(std)
}

// generateConcreteStructure generate the concrete project structure
func generateConcreteStructure(sections []string, startDate string, endDate string) structureType {
	std := properties.Concrete{}
	std = std.CreateConcreteProjectStructure(sections, startDate, endDate)
	return structureType(std)
}
