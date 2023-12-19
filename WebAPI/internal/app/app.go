package app

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/app/cmmsapp"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

// TODO: break out this interface into smaller embedded interfaces
type App interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error

	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

type AppTest interface {
	App
	ResetSequenceEquipment(int) error
	ResetSequenceTool(int) error
}

func New(db store.Store) *imp.App {
	return imp.New(db)
}

func NewTest(db store.StoreTest) *imp.AppTest {
	return imp.NewTest(db)
}
