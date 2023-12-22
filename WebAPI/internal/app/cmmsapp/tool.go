package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type Tool interface {
	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTool() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

func (cmms *App) CreateTool(title string, size string) (int, error) {
	return cmms.db.CreateTool(title, size)
}

func (cmms *App) DeleteTool(id int) error {
	return cmms.db.DeleteTool(id)
}

func (cmms *App) GetAllTool() ([]tp.Tool, error) {
	return cmms.db.GetAllTool()
}

func (cmms *App) GetTool(id int) (tp.Tool, error) {
	return cmms.db.GetTool(id)
}

func (cmms *App) UpdateTool(id int, title string, size string) error {
	return cmms.db.UpdateTool(id, title, size)
}
