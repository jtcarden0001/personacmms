package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) AssociateToolWithTask(assetId string, taskId string, toolId string) (tp.Tool, error) {
	return tp.Tool{}, ae.New(ae.CodeNotImplemented, "AssociateToolWithTask not implemented")
}

func (a *App) AssociateToolWithWorkOrder(assetId string, workOrderId string, toolId string) (tp.Tool, error) {
	return tp.Tool{}, ae.New(ae.CodeNotImplemented, "AssociateToolWithWorkOrder not implemented")
}

func (a *App) CreateTool(tool tp.Tool) (tp.Tool, error) {
	if tool.Id != uuid.Nil {
		return tp.Tool{}, ae.New(ae.CodeInvalid, "tool id must be nil on create, we will create an id for you")
	}
	tool.Id = uuid.New()

	return tp.Tool{}, ae.New(ae.CodeNotImplemented, "CreateTool not implemented")
}

func (a *App) DeleteTool(toolId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteTool not implemented")
}

func (a *App) DisassociateToolWithTask(assetId string, taskId string, toolId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateToolWithTask not implemented")
}

func (a *App) DisassociateToolWithWorkOrder(assetId string, workOrderId string, toolId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateToolWithWorkOrder not implemented")
}

func (a *App) GetTool(toolId string) (tp.Tool, error) {
	return tp.Tool{}, ae.New(ae.CodeNotImplemented, "GetTool not implemented")
}

func (a *App) ListTools() ([]tp.Tool, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTools not implemented")
}

func (a *App) UpdateTool(toolId string, tool tp.Tool) (tp.Tool, error) {
	return tp.Tool{}, ae.New(ae.CodeNotImplemented, "UpdateTool not implemented")
}

func (a *App) validateTool(tool tp.Tool) error {
	return ae.New(ae.CodeNotImplemented, "validateTool not implemented")
}

func (a *App) toolExists(id uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "toolExists not implemented")
}
