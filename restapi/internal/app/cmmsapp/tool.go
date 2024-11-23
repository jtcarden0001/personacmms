package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Tool interface {
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	ListTools() ([]tp.Tool, error)
	GetTool(string) (tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)
}

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
