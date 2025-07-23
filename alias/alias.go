package alias

import (
	"encoding/json"
	"os"
)

// type Function struct {
// 	Name     string
// 	Function string
// }

type Project struct {
	Name      string
	Functions map[string]string
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "file not found"
}

// Save projects to a specified file location
func SaveProjects(filename string, projects []Project) error {
	b, err := json.Marshal(projects)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadProjects(filename string) ([]Project, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return []Project{}, err
	}

	var projects []Project
	if err := json.Unmarshal(b, &projects); err != nil {
		return []Project{}, err
	}

	return projects, nil
}
