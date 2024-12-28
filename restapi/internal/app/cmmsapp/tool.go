package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
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
	if tool.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "tool id is required")
	}

	if len(tool.Title) < tp.MinEntityTitleLength || len(tool.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("tool title must be between [%d] and [%d] characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) toolExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	_, err = a.db.GetTool(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}
