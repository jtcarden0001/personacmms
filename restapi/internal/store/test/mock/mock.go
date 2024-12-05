package mock

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type MockStore struct {
	// Add fields to store mock data if necessary
}

func New() *MockStore {
	return &MockStore{}
}

// asset
func (m *MockStore) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	// Mock implementation
	return asset, nil
}

func (m *MockStore) DeleteAsset(id, group string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListAssets() ([]tp.Asset, error) {
	// Mock implementation
	return []tp.Asset{}, nil
}

func (m *MockStore) ListAssetsByGroup(group string) ([]tp.Asset, error) {
	// Mock implementation
	return []tp.Asset{}, nil
}

func (m *MockStore) GetAsset(id, group string) (tp.Asset, error) {
	// Mock implementation
	return tp.Asset{}, nil
}

func (m *MockStore) UpdateAsset(id, group string, asset tp.Asset) (tp.Asset, error) {
	// Mock implementation
	return asset, nil
}

// asset task consumable
func (m *MockStore) CreateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	// Mock implementation
	return tc, nil
}

func (m *MockStore) DeleteTaskConsumable(id, taskId tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTaskConsumables() ([]tp.TaskConsumable, error) {
	// Mock implementation
	return []tp.TaskConsumable{}, nil
}

func (m *MockStore) GetTaskConsumable(id, taskId tp.UUID) (tp.TaskConsumable, error) {
	// Mock implementation
	return tp.TaskConsumable{}, nil
}

func (m *MockStore) UpdateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	// Mock implementation
	return tc, nil
}

// asset task tool
func (m *MockStore) CreateTaskTool(tt tp.TaskTool) (tp.TaskTool, error) {
	// Mock implementation
	return tt, nil
}

func (m *MockStore) DeleteTaskTool(id, taskId tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTaskTools() ([]tp.TaskTool, error) {
	// Mock implementation
	return []tp.TaskTool{}, nil
}

func (m *MockStore) GetTaskTool(id, taskId tp.UUID) (tp.TaskTool, error) {
	// Mock implementation
	return tp.TaskTool{}, nil
}

// category
func (m *MockStore) CreateCategory(category tp.Category) (tp.Category, error) {
	// Mock implementation
	return category, nil
}

func (m *MockStore) DeleteCategory(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListCategories() ([]tp.Category, error) {
	// Mock implementation
	return []tp.Category{}, nil
}

func (m *MockStore) GetCategory(id string) (tp.Category, error) {
	// Mock implementation
	return tp.Category{}, nil
}

func (m *MockStore) UpdateCategory(id string, category tp.Category) (tp.Category, error) {
	// Mock implementation
	return category, nil
}

// consumable
func (m *MockStore) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	// Mock implementation
	return consumable, nil
}

func (m *MockStore) DeleteConsumable(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListConsumables() ([]tp.Consumable, error) {
	// Mock implementation
	return []tp.Consumable{}, nil
}

func (m *MockStore) GetConsumable(id string) (tp.Consumable, error) {
	// Mock implementation
	return tp.Consumable{}, nil
}

func (m *MockStore) UpdateConsumable(id string, consumable tp.Consumable) (tp.Consumable, error) {
	// Mock implementation
	return consumable, nil
}

// date trigger
func (m *MockStore) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	// Mock implementation
	return dt, nil
}

func (m *MockStore) DeleteDateTrigger(id tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListDateTriggers() ([]tp.DateTrigger, error) {
	// Mock implementation
	return []tp.DateTrigger{}, nil
}

func (m *MockStore) ListDateTriggersByTaskId(id tp.UUID) ([]tp.DateTrigger, error) {
	// Mock implementation
	return []tp.DateTrigger{}, nil
}

func (m *MockStore) GetDateTrigger(id tp.UUID) (tp.DateTrigger, error) {
	// Mock implementation
	return tp.DateTrigger{}, nil
}

func (m *MockStore) UpdateDateTrigger(id tp.UUID, dt tp.DateTrigger) (tp.DateTrigger, error) {
	// Mock implementation
	return dt, nil
}

// group
func (m *MockStore) CreateGroup(group tp.Group) (tp.Group, error) {
	// Mock implementation
	return group, nil
}

func (m *MockStore) DeleteGroup(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListGroups() ([]tp.Group, error) {
	// Mock implementation
	return []tp.Group{}, nil
}

func (m *MockStore) GetGroup(id string) (tp.Group, error) {
	// Mock implementation
	return tp.Group{}, nil
}

func (m *MockStore) UpdateGroup(id string, group tp.Group) (tp.Group, error) {
	// Mock implementation
	return group, nil
}

// task
func (m *MockStore) CreateTask(task tp.Task) (tp.Task, error) {
	// Mock implementation
	return task, nil
}

func (m *MockStore) DeleteTask(id tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTasks() ([]tp.Task, error) {
	// Mock implementation
	return []tp.Task{}, nil
}

func (m *MockStore) GetTask(id tp.UUID) (tp.Task, error) {
	// Mock implementation
	return tp.Task{}, nil
}

func (m *MockStore) UpdateTask(id tp.UUID, task tp.Task) (tp.Task, error) {
	// Mock implementation
	return task, nil
}

// task template
func (m *MockStore) CreateTaskTemplate(tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	// Mock implementation
	return tt, nil
}

func (m *MockStore) DeleteTaskTemplate(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	// Mock implementation
	return []tp.TaskTemplate{}, nil
}

func (m *MockStore) GetTaskTemplate(id string) (tp.TaskTemplate, error) {
	// Mock implementation
	return tp.TaskTemplate{}, nil
}

func (m *MockStore) UpdateTaskTemplate(id string, tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	// Mock implementation
	return tt, nil
}

// time trigger
func (m *MockStore) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	// Mock implementation
	return tt, nil
}

func (m *MockStore) DeleteTimeTrigger(id tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	// Mock implementation
	return []tp.TimeTrigger{}, nil
}

func (m *MockStore) ListTimeTriggersByTaskId(id tp.UUID) ([]tp.TimeTrigger, error) {
	// Mock implementation
	return []tp.TimeTrigger{}, nil
}

func (m *MockStore) GetTimeTrigger(id tp.UUID) (tp.TimeTrigger, error) {
	// Mock implementation
	return tp.TimeTrigger{}, nil
}

func (m *MockStore) UpdateTimeTrigger(id tp.UUID, tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	// Mock implementation
	return tt, nil
}

// time unit
func (m *MockStore) CreateTimeUnit(tu tp.TimeUnit) (tp.TimeUnit, error) {
	// Mock implementation
	return tu, nil
}

func (m *MockStore) DeleteTimeUnit(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTimeUnits() ([]tp.TimeUnit, error) {
	// Mock implementation
	return []tp.TimeUnit{}, nil
}

func (m *MockStore) GetTimeUnit(id string) (tp.TimeUnit, error) {
	// Mock implementation
	return tp.TimeUnit{}, nil
}

func (m *MockStore) UpdateTimeUnit(id string, tu tp.TimeUnit) (tp.TimeUnit, error) {
	// Mock implementation
	return tu, nil
}

// tool
func (m *MockStore) CreateTool(tool tp.Tool) (tp.Tool, error) {
	// Mock implementation
	return tool, nil
}

func (m *MockStore) DeleteTool(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListTools() ([]tp.Tool, error) {
	// Mock implementation
	return []tp.Tool{}, nil
}

func (m *MockStore) GetTool(id string) (tp.Tool, error) {
	// Mock implementation
	return tp.Tool{}, nil
}

func (m *MockStore) UpdateTool(id string, tool tp.Tool) (tp.Tool, error) {
	// Mock implementation
	return tool, nil
}

// usage trigger
func (m *MockStore) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	// Mock implementation
	return ut, nil
}

func (m *MockStore) DeleteUsageTrigger(id tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	// Mock implementation
	return []tp.UsageTrigger{}, nil
}

func (m *MockStore) GetUsageTrigger(id tp.UUID) (tp.UsageTrigger, error) {
	// Mock implementation
	return tp.UsageTrigger{}, nil
}

func (m *MockStore) UpdateUsageTrigger(id tp.UUID, ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	// Mock implementation
	return ut, nil
}

// usage unit
func (m *MockStore) CreateUsageUnit(uu tp.UsageUnit) (tp.UsageUnit, error) {
	// Mock implementation
	return uu, nil
}

func (m *MockStore) DeleteUsageUnit(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListUsageUnits() ([]tp.UsageUnit, error) {
	// Mock implementation
	return []tp.UsageUnit{}, nil
}

func (m *MockStore) GetUsageUnit(id string) (tp.UsageUnit, error) {
	// Mock implementation
	return tp.UsageUnit{}, nil
}

func (m *MockStore) UpdateUsageUnit(id string, uu tp.UsageUnit) (tp.UsageUnit, error) {
	// Mock implementation
	return uu, nil
}

// work order
func (m *MockStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	// Mock implementation
	return wo, nil
}

func (m *MockStore) DeleteWorkOrder(id tp.UUID) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	// Mock implementation
	return []tp.WorkOrder{}, nil
}

func (m *MockStore) GetWorkOrder(id tp.UUID) (tp.WorkOrder, error) {
	// Mock implementation
	return tp.WorkOrder{}, nil
}

func (m *MockStore) UpdateWorkOrder(id tp.UUID, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// Mock implementation
	return wo, nil
}

// work order status
func (m *MockStore) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	// Mock implementation
	return wos, nil
}

func (m *MockStore) DeleteWorkOrderStatus(id string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) ListWorkOrderStatuses() ([]tp.WorkOrderStatus, error) {
	// Mock implementation
	return []tp.WorkOrderStatus{}, nil
}

func (m *MockStore) GetWorkOrderStatus(id string) (tp.WorkOrderStatus, error) {
	// Mock implementation
	return tp.WorkOrderStatus{}, nil
}

func (m *MockStore) UpdateWorkOrderStatus(id string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	// Mock implementation
	return wos, nil
}

// utilities
func (m *MockStore) Exec(query string) error {
	// Mock implementation
	return nil
}

func (m *MockStore) Close() error {
	// Mock implementation
	return nil
}
