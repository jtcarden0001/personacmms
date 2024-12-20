package mock

import (
	"sync"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

type MockStore struct {
	data map[string]map[string]interface{}
	mu   sync.RWMutex
}

func New() *MockStore {
	return &MockStore{
		data: make(map[string]map[string]interface{}),
	}
}

// asset
func (m *MockStore) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["assets"] == nil {
		m.data["assets"] = make(map[string]interface{})
	}
	asset.Id = uuid.New()
	m.data["assets"][asset.Id.String()] = asset
	return asset, nil
}

func (m *MockStore) DeleteAsset(id, group string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["assets"][id]; !ok {
		return ae.New(ae.CodeNotFound, "asset not found")
	}
	delete(m.data["assets"], id)
	return nil
}

func (m *MockStore) GetAsset(id, group string) (tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if asset, ok := m.data["assets"][id]; ok {
		return asset.(tp.Asset), nil
	}
	return tp.Asset{}, nil
}

func (m *MockStore) ListAssets() ([]tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var assets []tp.Asset
	for _, asset := range m.data["assets"] {
		assets = append(assets, asset.(tp.Asset))
	}
	return assets, nil
}

func (m *MockStore) ListAssetsByGroup(group string) ([]tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var assets []tp.Asset
	for _, asset := range m.data["assets"] {
		if asset.(tp.Asset).GroupTitle == group {
			assets = append(assets, asset.(tp.Asset))
		}
	}
	return assets, nil
}

func (m *MockStore) UpdateAsset(id, group string, asset tp.Asset) (tp.Asset, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["assets"][id]; !ok {
		return tp.Asset{}, ae.New(ae.CodeNotFound, "asset not found")
	}
	m.data["assets"][id] = asset
	return asset, nil
}

// task consumable
func (m *MockStore) CreateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskConsumables"] == nil {
		m.data["taskConsumables"] = make(map[string]interface{})
	}
	id := tc.TaskId.String() + tc.ConsumableId.String()
	m.data["taskConsumables"][id] = tc
	return tc, nil
}

func (m *MockStore) DeleteTaskConsumable(taskId, consumableId tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := taskId.String() + consumableId.String()
	if _, ok := m.data["taskConsumables"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task consumable not found")
	}
	delete(m.data["taskConsumables"], id)
	return nil
}

func (m *MockStore) GetTaskConsumable(taskId, consumableId tp.UUID) (tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	id := taskId.String() + consumableId.String()
	if tc, ok := m.data["taskConsumables"][id]; ok {
		return tc.(tp.TaskConsumable), nil
	}
	return tp.TaskConsumable{}, nil
}

func (m *MockStore) ListTaskConsumables() ([]tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskConsumables []tp.TaskConsumable
	for _, tc := range m.data["taskConsumables"] {
		taskConsumables = append(taskConsumables, tc.(tp.TaskConsumable))
	}
	return taskConsumables, nil
}

func (m *MockStore) ListTaskConsumablesByTaskId(taskId tp.UUID) ([]tp.TaskConsumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskConsumables []tp.TaskConsumable
	for _, tc := range m.data["taskConsumables"] {
		if tc.(tp.TaskConsumable).TaskId == taskId {
			taskConsumables = append(taskConsumables, tc.(tp.TaskConsumable))
		}
	}
	return taskConsumables, nil
}

func (m *MockStore) UpdateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := tc.TaskId.String() + tc.ConsumableId.String()
	if _, ok := m.data["taskConsumables"][id]; !ok {
		return tp.TaskConsumable{}, ae.New(ae.CodeNotFound, "task consumable not found")
	}
	m.data["taskConsumables"][id] = tc
	return tc, nil
}

// task tool
func (m *MockStore) CreateTaskTool(tt tp.TaskTool) (tp.TaskTool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskTools"] == nil {
		m.data["taskTools"] = make(map[string]interface{})
	}
	id := tt.TaskId.String() + tt.ToolId.String()
	m.data["taskTools"][id] = tt
	return tt, nil
}

func (m *MockStore) DeleteTaskTool(taskId, toolId tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := taskId.String() + toolId.String()
	if _, ok := m.data["taskTools"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task tool not found")
	}
	delete(m.data["taskTools"], id)
	return nil
}

func (m *MockStore) GetTaskTool(taskId, toolId tp.UUID) (tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	id := taskId.String() + toolId.String()
	if tt, ok := m.data["taskTools"][id]; ok {
		return tt.(tp.TaskTool), nil
	}
	return tp.TaskTool{}, nil
}

func (m *MockStore) ListTaskTools() ([]tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTools []tp.TaskTool
	for _, tt := range m.data["taskTools"] {
		taskTools = append(taskTools, tt.(tp.TaskTool))
	}
	return taskTools, nil
}

func (m *MockStore) ListTaskToolsByTaskId(taskId tp.UUID) ([]tp.TaskTool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTools []tp.TaskTool
	for _, tt := range m.data["taskTools"] {
		if tt.(tp.TaskTool).TaskId == taskId {
			taskTools = append(taskTools, tt.(tp.TaskTool))
		}
	}
	return taskTools, nil
}

// category
func (m *MockStore) CreateCategory(category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["categories"] == nil {
		m.data["categories"] = make(map[string]interface{})
	}
	category.Id = uuid.New()
	m.data["categories"][category.Title] = category
	return category, nil
}

func (m *MockStore) DeleteCategory(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["categories"][title]; !ok {
		return ae.New(ae.CodeNotFound, "category not found")
	}
	delete(m.data["categories"], title)
	return nil
}

func (m *MockStore) GetCategory(title string) (tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if category, ok := m.data["categories"][title]; ok {
		return category.(tp.Category), nil
	}
	return tp.Category{}, nil
}

func (m *MockStore) ListCategories() ([]tp.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var categories []tp.Category
	for _, category := range m.data["categories"] {
		categories = append(categories, category.(tp.Category))
	}
	return categories, nil
}

func (m *MockStore) UpdateCategory(title string, category tp.Category) (tp.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["categories"][title]; !ok {
		return tp.Category{}, ae.New(ae.CodeNotFound, "category not found")
	}
	m.data["categories"][title] = category
	return category, nil
}

// consumable
func (m *MockStore) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["consumables"] == nil {
		m.data["consumables"] = make(map[string]interface{})
	}
	consumable.Id = uuid.New()
	m.data["consumables"][consumable.Title] = consumable
	return consumable, nil
}

func (m *MockStore) DeleteConsumable(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["consumables"][title]; !ok {
		return ae.New(ae.CodeNotFound, "consumable not found")
	}
	delete(m.data["consumables"], title)
	return nil
}

func (m *MockStore) GetConsumableByTitle(title string) (tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if consumable, ok := m.data["consumables"][title]; ok {
		return consumable.(tp.Consumable), nil
	}
	return tp.Consumable{}, nil
}

func (m *MockStore) GetConsumableById(id uuid.UUID) (tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, consumable := range m.data["consumables"] {
		if consumable.(tp.Consumable).Id == id {
			return consumable.(tp.Consumable), nil
		}
	}
	return tp.Consumable{}, nil
}

func (m *MockStore) ListConsumables() ([]tp.Consumable, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var consumables []tp.Consumable
	for _, consumable := range m.data["consumables"] {
		consumables = append(consumables, consumable.(tp.Consumable))
	}
	return consumables, nil
}

func (m *MockStore) UpdateConsumable(title string, consumable tp.Consumable) (tp.Consumable, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["consumables"][title]; !ok {
		return tp.Consumable{}, ae.New(ae.CodeNotFound, "consumable not found")
	}
	m.data["consumables"][title] = consumable
	return consumable, nil
}

// date trigger
func (m *MockStore) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["dateTriggers"] == nil {
		m.data["dateTriggers"] = make(map[string]interface{})
	}
	dt.Id = uuid.New()
	m.data["dateTriggers"][dt.Id.String()] = dt
	return dt, nil
}

func (m *MockStore) DeleteDateTrigger(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["dateTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "date trigger not found")
	}
	delete(m.data["dateTriggers"], id.String())
	return nil
}

func (m *MockStore) GetDateTrigger(id tp.UUID) (tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if dt, ok := m.data["dateTriggers"][id.String()]; ok {
		return dt.(tp.DateTrigger), nil
	}
	return tp.DateTrigger{}, nil
}

func (m *MockStore) ListDateTriggers() ([]tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var dateTriggers []tp.DateTrigger
	for _, dt := range m.data["dateTriggers"] {
		dateTriggers = append(dateTriggers, dt.(tp.DateTrigger))
	}
	return dateTriggers, nil
}

func (m *MockStore) ListDateTriggersByTaskId(id tp.UUID) ([]tp.DateTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var dateTriggers []tp.DateTrigger
	for _, dt := range m.data["dateTriggers"] {
		if dt.(tp.DateTrigger).TaskId == id {
			dateTriggers = append(dateTriggers, dt.(tp.DateTrigger))
		}
	}
	return dateTriggers, nil
}

func (m *MockStore) UpdateDateTrigger(id tp.UUID, dt tp.DateTrigger) (tp.DateTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["dateTriggers"][id.String()]; !ok {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound, "date trigger not found")
	}
	m.data["dateTriggers"][id.String()] = dt
	return dt, nil
}

// group
func (m *MockStore) CreateGroup(group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["groups"] == nil {
		m.data["groups"] = make(map[string]interface{})
	}
	group.Id = uuid.New()
	m.data["groups"][group.Title] = group
	return group, nil
}

func (m *MockStore) DeleteGroup(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["groups"][title]; !ok {
		return ae.New(ae.CodeNotFound, "group not found")
	}
	delete(m.data["groups"], title)
	return nil
}

func (m *MockStore) GetGroup(title string) (tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if group, ok := m.data["groups"][title]; ok {
		return group.(tp.Group), nil
	}
	return tp.Group{}, nil
}

func (m *MockStore) ListGroups() ([]tp.Group, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var groups []tp.Group
	for _, group := range m.data["groups"] {
		groups = append(groups, group.(tp.Group))
	}
	return groups, nil
}

func (m *MockStore) UpdateGroup(title string, group tp.Group) (tp.Group, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["groups"][title]; !ok {
		return tp.Group{}, ae.New(ae.CodeNotFound, "group not found")
	}
	m.data["groups"][title] = group
	return group, nil
}

// task
func (m *MockStore) CreateTask(task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["tasks"] == nil {
		m.data["tasks"] = make(map[string]interface{})
	}
	task.Id = uuid.New()
	m.data["tasks"][task.Id.String()] = task
	return task, nil
}

func (m *MockStore) DeleteTask(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tasks"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "task not found")
	}
	delete(m.data["tasks"], id.String())
	return nil
}

func (m *MockStore) GetTask(id tp.UUID) (tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := m.data["tasks"][id.String()]; ok {
		return task.(tp.Task), nil
	}
	return tp.Task{}, nil
}

func (m *MockStore) GetTaskByAssetId(assetId tp.UUID, taskId tp.UUID) (tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := m.data["tasks"][taskId.String()]; ok {
		if task.(tp.Task).AssetId == assetId {
			return task.(tp.Task), nil
		}
	}

	return tp.Task{}, ae.New(ae.CodeNotFound, "task not found")
}

func (m *MockStore) ListTasks() ([]tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tasks []tp.Task
	for _, task := range m.data["tasks"] {
		tasks = append(tasks, task.(tp.Task))
	}
	return tasks, nil
}

func (m *MockStore) ListTasksByAssetId(id tp.UUID) ([]tp.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tasks []tp.Task
	for _, task := range m.data["tasks"] {
		if task.(tp.Task).AssetId == id {
			tasks = append(tasks, task.(tp.Task))
		}
	}
	return tasks, nil
}

func (m *MockStore) UpdateTask(id tp.UUID, task tp.Task) (tp.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tasks"][id.String()]; !ok {
		return tp.Task{}, ae.New(ae.CodeNotFound, "task not found")
	}
	m.data["tasks"][id.String()] = task
	return task, nil
}

// task template
func (m *MockStore) CreateTaskTemplate(tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["taskTemplates"] == nil {
		m.data["taskTemplates"] = make(map[string]interface{})
	}
	tt.Id = uuid.New()
	m.data["taskTemplates"][tt.Id.String()] = tt
	return tt, nil
}

func (m *MockStore) DeleteTaskTemplate(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["taskTemplates"][id]; !ok {
		return ae.New(ae.CodeNotFound, "task template not found")
	}
	delete(m.data["taskTemplates"], id)
	return nil
}

func (m *MockStore) GetTaskTemplate(id string) (tp.TaskTemplate, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tt, ok := m.data["taskTemplates"][id]; ok {
		return tt.(tp.TaskTemplate), nil
	}
	return tp.TaskTemplate{}, nil
}

func (m *MockStore) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var taskTemplates []tp.TaskTemplate
	for _, tt := range m.data["taskTemplates"] {
		taskTemplates = append(taskTemplates, tt.(tp.TaskTemplate))
	}
	return taskTemplates, nil
}

func (m *MockStore) UpdateTaskTemplate(id string, tt tp.TaskTemplate) (tp.TaskTemplate, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["taskTemplates"][id]; !ok {
		return tp.TaskTemplate{}, ae.New(ae.CodeNotFound, "task template not found")
	}
	m.data["taskTemplates"][id] = tt
	return tt, nil
}

// time trigger
func (m *MockStore) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["timeTriggers"] == nil {
		m.data["timeTriggers"] = make(map[string]interface{})
	}
	tt.Id = uuid.New()
	m.data["timeTriggers"][tt.Id.String()] = tt
	return tt, nil
}

func (m *MockStore) DeleteTimeTrigger(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "time trigger not found")
	}
	delete(m.data["timeTriggers"], id.String())
	return nil
}

func (m *MockStore) GetTimeTrigger(id tp.UUID) (tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tt, ok := m.data["timeTriggers"][id.String()]; ok {
		return tt.(tp.TimeTrigger), nil
	}
	return tp.TimeTrigger{}, nil
}

func (m *MockStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data["timeTriggers"] {
		timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
	}
	return timeTriggers, nil
}

func (m *MockStore) ListTimeTriggersByTaskId(id tp.UUID) ([]tp.TimeTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeTriggers []tp.TimeTrigger
	for _, tt := range m.data["timeTriggers"] {
		if tt.(tp.TimeTrigger).TaskId == id {
			timeTriggers = append(timeTriggers, tt.(tp.TimeTrigger))
		}
	}
	return timeTriggers, nil
}

func (m *MockStore) UpdateTimeTrigger(id tp.UUID, tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeTriggers"][id.String()]; !ok {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound, "time trigger not found")
	}
	m.data["timeTriggers"][id.String()] = tt
	return tt, nil
}

// time unit
func (m *MockStore) CreateTimeUnit(tu tp.TimeUnit) (tp.TimeUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["timeUnits"] == nil {
		m.data["timeUnits"] = make(map[string]interface{})
	}
	tu.Id = uuid.New()
	m.data["timeUnits"][tu.Id.String()] = tu
	return tu, nil
}

func (m *MockStore) DeleteTimeUnit(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeUnits"][id]; !ok {
		return ae.New(ae.CodeNotFound, "time unit not found")
	}
	delete(m.data["timeUnits"], id)
	return nil
}

func (m *MockStore) GetTimeUnit(id string) (tp.TimeUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tu, ok := m.data["timeUnits"][id]; ok {
		return tu.(tp.TimeUnit), nil
	}
	return tp.TimeUnit{}, nil
}

func (m *MockStore) ListTimeUnits() ([]tp.TimeUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var timeUnits []tp.TimeUnit
	for _, tu := range m.data["timeUnits"] {
		timeUnits = append(timeUnits, tu.(tp.TimeUnit))
	}
	return timeUnits, nil
}

func (m *MockStore) UpdateTimeUnit(id string, tu tp.TimeUnit) (tp.TimeUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["timeUnits"][id]; !ok {
		return tp.TimeUnit{}, ae.New(ae.CodeNotFound, "time unit not found")
	}
	m.data["timeUnits"][id] = tu
	return tu, nil
}

// tool
func (m *MockStore) CreateTool(tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["tools"] == nil {
		m.data["tools"] = make(map[string]interface{})
	}
	tool.Id = uuid.New()
	m.data["tools"][tool.Title] = tool
	return tool, nil
}

func (m *MockStore) DeleteTool(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tools"][title]; !ok {
		return ae.New(ae.CodeNotFound, "tool not found")
	}
	delete(m.data["tools"], title)
	return nil
}

func (m *MockStore) GetTool(title string) (tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tool, ok := m.data["tools"][title]; ok {
		return tool.(tp.Tool), nil
	}
	return tp.Tool{}, nil
}

func (m *MockStore) GetToolById(id uuid.UUID) (tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, tool := range m.data["tools"] {
		if tool.(tp.Tool).Id == id {
			return tool.(tp.Tool), nil
		}
	}

	return tp.Tool{}, nil
}

func (m *MockStore) ListTools() ([]tp.Tool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var tools []tp.Tool
	for _, tool := range m.data["tools"] {
		tools = append(tools, tool.(tp.Tool))
	}
	return tools, nil
}

func (m *MockStore) UpdateTool(title string, tool tp.Tool) (tp.Tool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["tools"][title]; !ok {
		return tp.Tool{}, ae.New(ae.CodeNotFound, "tool not found")
	}
	m.data["tools"][title] = tool
	return tool, nil
}

// usage trigger
func (m *MockStore) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["usageTriggers"] == nil {
		m.data["usageTriggers"] = make(map[string]interface{})
	}
	ut.Id = uuid.New()
	m.data["usageTriggers"][ut.Id.String()] = ut
	return ut, nil
}

func (m *MockStore) DeleteUsageTrigger(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageTriggers"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "usage trigger not found")
	}
	delete(m.data["usageTriggers"], id.String())
	return nil
}

func (m *MockStore) GetUsageTrigger(id tp.UUID) (tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ut, ok := m.data["usageTriggers"][id.String()]; ok {
		return ut.(tp.UsageTrigger), nil
	}
	return tp.UsageTrigger{}, nil
}

func (m *MockStore) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageTriggers []tp.UsageTrigger
	for _, ut := range m.data["usageTriggers"] {
		usageTriggers = append(usageTriggers, ut.(tp.UsageTrigger))
	}
	return usageTriggers, nil
}

func (m *MockStore) ListUsageTriggersByTaskId(id tp.UUID) ([]tp.UsageTrigger, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageTriggers []tp.UsageTrigger
	for _, ut := range m.data["usageTriggers"] {
		if ut.(tp.UsageTrigger).TaskId == id {
			usageTriggers = append(usageTriggers, ut.(tp.UsageTrigger))
		}
	}
	return usageTriggers, nil
}

func (m *MockStore) UpdateUsageTrigger(id tp.UUID, ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageTriggers"][id.String()]; !ok {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, "usage trigger not found")
	}
	m.data["usageTriggers"][id.String()] = ut
	return ut, nil
}

// usage unit
func (m *MockStore) CreateUsageUnit(uu tp.UsageUnit) (tp.UsageUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["usageUnits"] == nil {
		m.data["usageUnits"] = make(map[string]interface{})
	}
	uu.Id = uuid.New()
	m.data["usageUnits"][uu.Title] = uu
	return uu, nil
}

func (m *MockStore) DeleteUsageUnit(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageUnits"][title]; !ok {
		return ae.New(ae.CodeNotFound, "usage unit not found")
	}
	delete(m.data["usageUnits"], title)
	return nil
}

func (m *MockStore) GetUsageUnit(title string) (tp.UsageUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if uu, ok := m.data["usageUnits"][title]; ok {
		return uu.(tp.UsageUnit), nil
	}
	return tp.UsageUnit{}, nil
}

func (m *MockStore) ListUsageUnits() ([]tp.UsageUnit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var usageUnits []tp.UsageUnit
	for _, uu := range m.data["usageUnits"] {
		usageUnits = append(usageUnits, uu.(tp.UsageUnit))
	}
	return usageUnits, nil
}

func (m *MockStore) UpdateUsageUnit(title string, uu tp.UsageUnit) (tp.UsageUnit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["usageUnits"][title]; !ok {
		return tp.UsageUnit{}, ae.New(ae.CodeNotFound, "usage unit not found")
	}
	m.data["usageUnits"][title] = uu
	return uu, nil
}

// work order
func (m *MockStore) CreateWorkOrder(wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["workOrders"] == nil {
		m.data["workOrders"] = make(map[string]interface{})
	}
	wo.Id = uuid.New()
	m.data["workOrders"][wo.Id.String()] = wo
	return wo, nil
}

func (m *MockStore) DeleteWorkOrder(id tp.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrders"][id.String()]; !ok {
		return ae.New(ae.CodeNotFound, "work order not found")
	}
	delete(m.data["workOrders"], id.String())
	return nil
}

func (m *MockStore) GetWorkOrder(id tp.UUID) (tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wo, ok := m.data["workOrders"][id.String()]; ok {
		return wo.(tp.WorkOrder), nil
	}
	return tp.WorkOrder{}, nil
}

func (m *MockStore) GetWorkOrderForTask(taskId tp.UUID, woId tp.UUID) (tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wo, ok := m.data["workOrders"][woId.String()]; ok {
		if wo.(tp.WorkOrder).TaskId == taskId {
			return wo.(tp.WorkOrder), nil
		}
	}
	return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "work order not found")
}

func (m *MockStore) ListWorkOrders() ([]tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var workOrders []tp.WorkOrder
	for _, wo := range m.data["workOrders"] {
		workOrders = append(workOrders, wo.(tp.WorkOrder))
	}
	return workOrders, nil
}

func (m *MockStore) ListWorkOrdersByTaskId(taskId tp.UUID) ([]tp.WorkOrder, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var workOrders []tp.WorkOrder
	for _, wo := range m.data["workOrders"] {
		if wo.(tp.WorkOrder).TaskId == taskId {
			workOrders = append(workOrders, wo.(tp.WorkOrder))
		}
	}
	return workOrders, nil
}

func (m *MockStore) UpdateWorkOrder(id tp.UUID, wo tp.WorkOrder) (tp.WorkOrder, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrders"][id.String()]; !ok {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, "work order not found")
	}
	m.data["workOrders"][id.String()] = wo
	return wo, nil
}

// work order status
func (m *MockStore) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data["workOrderStatuses"] == nil {
		m.data["workOrderStatuses"] = make(map[string]interface{})
	}
	wos.Id = uuid.New()
	m.data["workOrderStatuses"][wos.Title] = wos
	return wos, nil
}

func (m *MockStore) DeleteWorkOrderStatus(title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["workOrderStatuses"][title]; !ok {
		return ae.New(ae.CodeNotFound, "work order status not found")
	}
	delete(m.data["workOrderStatuses"], title)
	return nil
}

func (m *MockStore) GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if wos, ok := m.data["workOrderStatuses"][title]; ok {
		return wos.(tp.WorkOrderStatus), nil
	}
	return tp.WorkOrderStatus{}, nil
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
