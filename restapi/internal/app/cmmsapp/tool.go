package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateToolWithTask(assetId string, taskId string, toolId string, ts apitp.ToolSizeRequest) (apitp.ToolSizeResponse, error) {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.ToolSizeResponse{}, err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return apitp.ToolSizeResponse{}, errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return apitp.ToolSizeResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	// TODO check that ts doesnt conflict with path params

	s := ""
	if ts.Size != nil {
		s = *ts.Size
	}

	stToolSizeResponse, err := a.db.AssociateToolWithTask(task.Id, tUid, s)
	if err != nil {
		return apitp.ToolSizeResponse{}, errors.Wrapf(err, "AssociateToolWithTask failed")
	}

	return convertStoreToolSizeToApiToolSizeResponse(stToolSizeResponse)
}

func (a *App) AssociateToolWithWorkOrder(assetId string, workOrderId string, toolId string, ts apitp.ToolSizeRequest) (apitp.ToolSizeResponse, error) {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return apitp.ToolSizeResponse{}, err
	}

	tUid, tFound, err := a.toolExists(toolId)
	if err != nil {
		return apitp.ToolSizeResponse{}, errors.Wrapf(err, "error checking tool exists")
	}

	if !tFound {
		return apitp.ToolSizeResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("tool with id [%s] not found", toolId))
	}

	// TODO check that ts doesnt conflict with path params

	s := ""
	if ts.Size != nil {
		s = *ts.Size
	}

	stToolSizeResponse, err := a.db.AssociateToolWithWorkOrder(workOrder.Id, tUid, s)
	if err != nil {
		return apitp.ToolSizeResponse{}, errors.Wrapf(err, "AssociateToolWithWorkOrder failed")
	}

	return convertStoreToolSizeToApiToolSizeResponse(stToolSizeResponse)
}

func (a *App) CreateTool(tool apitp.ToolRequest) (apitp.ToolResponse, error) {
	if tool.Id != uuid.Nil {
		return apitp.ToolResponse{}, ae.New(ae.CodeInvalid, "tool id must be nil on create, we will create an id for you")
	}
	tool.Id = uuid.New()

	err := a.validateTool(tool)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "CreateTool validation failed")
	}

	stToolRequest, err := convertApiToolRequestToStoreTool(tool)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "CreateTool - error converting to store type")
	}

	stToolResponse, err := a.db.CreateTool(stToolRequest)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "CreateTool - error creating tool")
	}

	return convertStoreToolToApiToolResponse(stToolResponse)
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

func (a *App) GetTool(toolId string) (apitp.ToolResponse, error) {
	tUid, err := uuid.Parse(toolId)
	if err != nil {
		return apitp.ToolResponse{}, ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	stToolResponse, err := a.db.GetTool(tUid)
	if err != nil {
		return apitp.ToolResponse{}, err
	}

	return convertStoreToolToApiToolResponse(stToolResponse)
}

func (a *App) ListTools() ([]apitp.ToolResponse, error) {
	stToolResponses, err := a.db.ListTools()
	if err != nil {
		return nil, errors.Wrapf(err, "ListTools failed")
	}

	return convertStoreToolListToApiToolResponseList(stToolResponses)
}

func (a *App) UpdateTool(toolId string, tool apitp.ToolRequest) (apitp.ToolResponse, error) {
	tuid, err := uuid.Parse(toolId)
	if err != nil {
		return apitp.ToolResponse{}, ae.New(ae.CodeInvalid, "tool id must be a valid uuid")
	}

	if tool.Id != uuid.Nil && tool.Id != tuid {
		return apitp.ToolResponse{}, ae.New(ae.CodeInvalid,
			fmt.Sprintf("tool id mismatch between [%s] and [%s]", toolId, tool.Id))
	}

	tool.Id = tuid
	err = a.validateTool(tool)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "UpdateTool validation failed")
	}

	stToolRequest, err := convertApiToolRequestToStoreTool(tool)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "UpdateTool - error converting to store type")
	}

	stToolResponse, err := a.db.UpdateTool(stToolRequest)
	if err != nil {
		return apitp.ToolResponse{}, errors.Wrapf(err, "UpdateTool - error updating tool")
	}

	return convertStoreToolToApiToolResponse(stToolResponse)
}

func (a *App) validateTool(tool apitp.ToolRequest) error {
	if tool.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "tool id is required")
	}

	if len(tool.Title) < apitp.MinEntityTitleLength || len(tool.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("tool title must be between [%d] and [%d] characters",
				apitp.MinEntityTitleLength,
				apitp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) toolExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil || uid == uuid.Nil {
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

func convertApiToolRequestToStoreTool(tool apitp.ToolRequest) (storetp.Tool, error) {
	return storetp.Tool{}, ae.New(ae.CodeNotImplemented, "ConvertApiToolRequestToStoreTool not implemented")
}

func convertApiToolSizeRequestToStoreToolSize(toolSize apitp.ToolSizeRequest) (storetp.ToolSize, error) {
	return storetp.ToolSize{}, ae.New(ae.CodeNotImplemented, "ConvertApiToolSizeRequestToStoreToolSize not implemented")
}

func convertStoreToolListToApiToolResponseList(st []storetp.Tool) ([]apitp.ToolResponse, error) {
	return []apitp.ToolResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreToolListToApiToolResponseList not implemented")
}

func convertStoreToolSizeListToApiToolSizeResponseList(st []storetp.ToolSize) ([]apitp.ToolSizeResponse, error) {
	return []apitp.ToolSizeResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreToolSizeListToApiToolSizeResponseList not implemented")
}

func convertStoreToolSizeToApiToolSizeResponse(st storetp.ToolSize) (apitp.ToolSizeResponse, error) {
	return apitp.ToolSizeResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreToolSizeToApiToolSizeResponse not implemented")
}

func convertStoreToolToApiToolResponse(st storetp.Tool) (apitp.ToolResponse, error) {
	return apitp.ToolResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreToolToApiToolResponse not implemented")
}
