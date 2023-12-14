package cmmsapp

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

func (cmms *App) CreateTool(title string) (int, error) {
	return cmms.db.CreateTool(title)
}

func (cmms *App) DeleteTool(id int) error {
	return cmms.db.DeleteTool(id)
}

func (cmms *App) GetAllTools() ([]tp.Tool, error) {
	return cmms.db.GetAllTools()
}

func (cmms *App) GetTool(id int) (tp.Tool, error) {
	return cmms.db.GetTool(id)
}

func (cmms *App) UpdateTool(id int, title string) error {
	return cmms.db.UpdateTool(id, title)
}

func (cmms *AppTest) ResetSequenceTool(id int) error {
	return cmms.db.ResetSequenceTool(id)
}
