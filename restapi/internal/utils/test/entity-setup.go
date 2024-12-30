package test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func SetupAsset(identifier int, generateId bool) tp.Asset {
	year := 2000 + (identifier % 22) // Random year based on identifier
	return tp.Asset{
		Id:           getTestEntityId(generateId),
		Title:        fmt.Sprintf("Asset %d", identifier),
		Year:         ToPtr(year),
		Manufacturer: ToPtr(fmt.Sprintf("Asset %d manufacturer", identifier)),
		Make:         ToPtr(fmt.Sprintf("Asset %d make", identifier)),
		ModelNumber:  ToPtr(fmt.Sprintf("Asset %d model number", identifier)),
		SerialNumber: ToPtr(fmt.Sprintf("Asset %d serial number", identifier)),
		Description:  ToPtr(fmt.Sprintf("Asset %d description", identifier)),
	}
}

func SetupAssetAndTask(identifier int, generateId bool) (tp.Asset, tp.Task) {
	asset := SetupAsset(identifier, generateId)
	task := SetupTask(identifier, asset.Id, generateId)
	return asset, task
}

func SetupCategory(identifier int, generateId bool) tp.Category {
	return tp.Category{
		Id:          getTestEntityId(generateId),
		Title:       fmt.Sprintf("Category %d", identifier),
		Description: ToPtr(fmt.Sprintf("Category %d description", identifier)),
	}
}

func SetupConsumable(identifier int, generateId bool) tp.Consumable {
	return tp.Consumable{
		Id:    getTestEntityId(generateId),
		Title: fmt.Sprintf("Consumable %d", identifier),
	}
}

func SetupConsumableQuantity(identifier int, consumableId uuid.UUID, taskId uuid.UUID, generateId bool) tp.ConsumableQuantity {
	return tp.ConsumableQuantity{
		Id:       getTestEntityId(generateId),
		Title:    fmt.Sprintf("Consumable Quantity %d", identifier),
		Quantity: fmt.Sprintf("%d", identifier),
	}
}

func SetupDateTrigger(identifier int, taskId uuid.UUID, generateId bool) tp.DateTrigger {
	return tp.DateTrigger{
		Id:            getTestEntityId(generateId),
		ScheduledDate: time.Now().AddDate(0, identifier, 0),
		TaskId:        taskId,
	}
}

func SetupGroup(identifier int, generateId bool) tp.Group {
	return tp.Group{
		Id:    getTestEntityId(generateId),
		Title: fmt.Sprintf("Group %d", identifier),
	}
}

func SetupTask(identifier int, assetId uuid.UUID, generateId bool) tp.Task {
	return tp.Task{
		Id:           getTestEntityId(generateId),
		Title:        fmt.Sprintf("Task %d", identifier),
		Instructions: ToPtr(fmt.Sprintf("Task %d instructions", identifier)),
		AssetId:      assetId,
	}
}

func SetupTimeTrigger(identifier int, taskId uuid.UUID, generateId bool) tp.TimeTrigger {
	return tp.TimeTrigger{
		Id:       getTestEntityId(generateId),
		TaskId:   taskId,
		Quantity: identifier,
		TimeUnit: tp.TimeTriggerUnitDays,
	}
}

func SetupTool(identifier int, generateId bool) tp.Tool {
	return tp.Tool{
		Id:    getTestEntityId(generateId),
		Title: "Tool " + strconv.Itoa(identifier),
	}
}

func SetupToolSize(identifier int, toolId uuid.UUID, generateId bool) tp.ToolSize {
	return tp.ToolSize{
		Id:    getTestEntityId(generateId),
		Title: "Tool Size " + strconv.Itoa(identifier),
		Size:  ToPtr("Size " + strconv.Itoa(identifier)),
	}
}

func SetupUsageTrigger(identifier int, taskId uuid.UUID, generateId bool) tp.UsageTrigger {
	return tp.UsageTrigger{
		Id:        getTestEntityId(generateId),
		TaskId:    taskId,
		Quantity:  identifier,
		UsageUnit: tp.UsageTriggerUnitDays,
	}
}

func SetupWorkOrder(identifier int, assetId uuid.UUID, generateId bool) tp.WorkOrder {
	return tp.WorkOrder{
		Id:              getTestEntityId(generateId),
		Title:           "Work Order " + strconv.Itoa(identifier),
		CreatedDate:     time.Now().AddDate(0, 0, -identifier),
		CompletedDate:   ToPtr(time.Now().AddDate(0, 0, identifier)),
		Instructions:    ToPtr("Instructions " + strconv.Itoa(identifier)),
		Notes:           ToPtr("Notes " + strconv.Itoa(identifier)),
		CumulativeMiles: ToPtr(identifier * 100),
		CumulativeHours: ToPtr(identifier * 10),
		AssetId:         assetId,
		Status:          tp.WorkOrderStatusComplete,
	}
}

func getTestEntityId(generateId bool) uuid.UUID {
	if generateId {
		return uuid.New()
	}
	return uuid.Nil
}
