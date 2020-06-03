package project

// Project properties definition
type Project struct {
	ProjectName string
	Icon        uint8
	Manager     string
	Members     []string
	Description string
	Visibility  bool
	Tags        []string
}

// GenerateProject return a project instance
func (p *Project) GenerateProject(projectName string, icon uint8, manager string, members []string, description string, visibility bool, tags []string) Project {

	// Project declaration and initialization
	project := Project{
		ProjectName: projectName,
		Icon:        icon,
		Manager:     manager,
		Members:     members,
		Description: description,
		Visibility:  visibility,
		Tags:        tags,
	}

	return project
}
