package cmmsapp

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

func (cmms *App) CreateTool(title string) (int, error) {
	return cmms.store.CreateTool(title)
}

func (cmms *App) DeleteTool(id int) error {
	return cmms.store.DeleteTool(id)
}

func (cmms *App) GetAllTools() ([]tp.Tool, error) {
	return cmms.store.GetAllTools()
}

func (cmms *App) GetTool(id int) (tp.Tool, error) {
	return cmms.store.GetTool(id)
}

func (cmms *App) UpdateTool(id int, title string) error {
	return cmms.store.UpdateTool(id, title)
}

func (cmms *AppTest) ResetSequenceTool(id int) error {
	return cmms.store.ResetSequenceTool(id)
}
