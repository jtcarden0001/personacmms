package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateToolWithTask(assetId string, taskId string, toolId string, ts tp.ToolSize) (tp.ToolSize, error) {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.ToolSize{}, err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return tp.ToolSize{}, errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return tp.ToolSize{}, ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	// TODO check that ts doesnt conflict with path params

	s := ""
	if ts.Size != nil {
		s = *ts.Size
	}

	return a.db.AssociateToolWithTask(task.Id, tUid, s)
}

func (a *App) AssociateToolWithWorkOrder(assetId string, workOrderId string, toolId string, ts tp.ToolSize) (tp.ToolSize, error) {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return tp.ToolSize{}, err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return tp.ToolSize{}, errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return tp.ToolSize{}, ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	// TODO check that ts doesnt conflict with path params

	s := ""
	if ts.Size != nil {
		s = *ts.Size
	}

	return a.db.AssociateToolWithWorkOrder(workOrder.Id, tUid, s)
}

func (a *App) CreateTool(tool tp.Tool) (tp.Tool, error) {
	if tool.Id != uuid.Nil {
		return tp.Tool{}, ae.New(ae.CodeInvalid, "tool id must be nil on create, we will create an id for you")
	}
	tool.Id = uuid.New()

	err := a.validateTool(tool)
	if err != nil {
		return tp.Tool{}, errors.Wrapf(err, "CreateTool validation failed")
	}

	return a.db.CreateTool(tool)
}

func (a *App) DeleteTool(toolId string) error {
	tUid, err := uuid.Parse(toolId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	// TODO: block deletion if the tool is in use

	return a.db.DeleteTool(tUid)
}

func (a *App) DisassociateToolWithTask(assetId string, taskId string, toolId string) error {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	return a.db.DisassociateToolWithTask(task.Id, tUid)
}

func (a *App) DisassociateToolWithWorkOrder(assetId string, workOrderId string, toolId string) error {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	return a.db.DisassociateToolWithWorkOrder(workOrder.Id, tUid)
}

func (a *App) GetTool(toolId string) (tp.Tool, error) {
	tUid, err := uuid.Parse(toolId)
	if err != nil {
		return tp.Tool{}, ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	return a.db.GetTool(tUid)
}

func (a *App) ListTools() ([]tp.Tool, error) {
	return a.db.ListTools()
}

func (a *App) UpdateTool(toolId string, tool tp.Tool) (tp.Tool, error) {
	tuid, err := uuid.Parse(toolId)
	if err != nil {
		return tp.Tool{}, ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	if tool.Id != uuid.Nil && tool.Id != tuid {
		return tp.Tool{}, ae.New(ae.CodeInvalid,
			fmt.Sprintf("tool id mismatch between [%s] and [%s]", toolId, tool.Id))
	}

	tool.Id = tuid
	err = a.validateTool(tool)
	if err != nil {
		return tp.Tool{}, errors.Wrapf(err, "UpdateTool validation failed")
	}

	return a.db.UpdateTool(tool)
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
