package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateTool(tool tp.Tool) (tp.Tool, error) {
	return a.db.CreateTool(tool)
}

func (a *App) DeleteTool(title string) error {
	return a.db.DeleteTool(title)
}

func (a *App) ListTools() ([]tp.Tool, error) {
	return a.db.ListTools()
}

func (a *App) GetTool(title string) (tp.Tool, error) {
	return a.db.GetTool(title)
}

func (a *App) UpdateTool(oldTitle string, tool tp.Tool) (tp.Tool, error) {
	return a.db.UpdateTool(oldTitle, tool)
}
