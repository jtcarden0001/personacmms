package cmmsapp

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssociateConsumableWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateConsumableWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithTask: failed during setup. CreateConsumable() failed: %v", err)
	}

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	t1 := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), t1)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		taskId        string
		consumableId  string
		cq            tp.ConsumableQuantity
		shouldSucceed bool
	}{
		{"valid consumable association", createdTask.Id.String(), createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"}, true},
		{"invalid task ID", "invalid", createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"}, false},
		{"invalid consumable ID", createdTask.Id.String(), "invalid", tp.ConsumableQuantity{Quantity: "10"}, false},
		{"non-existent consumable", createdTask.Id.String(), uuid.New().String(), tp.ConsumableQuantity{Quantity: "10"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cq, err := app.AssociateConsumableWithTask(createdAsset.Id.String(), tc.taskId, tc.consumableId, tc.cq)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("AssociateConsumableWithTask() failed: %v", err)
					return
				}

				if cq.Quantity != tc.cq.Quantity {
					t.Errorf("AssociateConsumableWithTask() failed: expected %s, got %s", tc.cq.Quantity, cq.Quantity)
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateConsumableWithTask() should have failed with %s", tc.name)
				return
			}
		})
	}
}

func TestAssociateConsumableWithWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateConsumableWithWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithWorkOrder: failed during setup. CreateConsumable() failed: %v", err)
	}

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo1 := utest.SetupWorkOrder(1, createdAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(createdAsset.Id.String(), wo1)
	if err != nil {
		t.Errorf("TestAssociateConsumableWithWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		workOrderId   string
		consumableId  string
		cq            tp.ConsumableQuantity
		shouldSucceed bool
	}{
		{"valid consumable association", createdWorkOrder.Id.String(), createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"}, true},
		{"invalid work order ID", "invalid", createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"}, false},
		{"invalid consumable ID", createdWorkOrder.Id.String(), "invalid", tp.ConsumableQuantity{Quantity: "10"}, false},
		{"non-existent consumable", createdWorkOrder.Id.String(), uuid.New().String(), tp.ConsumableQuantity{Quantity: "10"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cq, err := app.AssociateConsumableWithWorkOrder(createdAsset.Id.String(), tc.workOrderId, tc.consumableId, tc.cq)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("AssociateConsumableWithWorkOrder() failed: %v", err)
					return
				}

				if cq.Quantity != tc.cq.Quantity {

					t.Errorf("AssociateConsumableWithWorkOrder() failed: expected %s, got %s", tc.cq.Quantity, cq.Quantity)
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateConsumableWithWorkOrder() should have failed with %s", tc.name)
				return
			}
		})
	}
}

func TestCreateConsumable(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateConsumable")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingConsumable := utest.SetupConsumable(1, false)
	_, err = app.CreateConsumable(conflictingConsumable)
	if err != nil {
		t.Errorf("TestCreateConsumable: failed during setup. CreateConsumable() failed: %v", err)
	}

	emptyTitleConsumable := utest.SetupConsumable(2, false)
	emptyTitleConsumable.Title = ""

	testCases := []struct {
		name          string
		consumable    tp.Consumable
		shouldSucceed bool
	}{
		{"valid consumable", utest.SetupConsumable(3, false), true},
		{"non nil id", utest.SetupConsumable(4, true), false},
		{"empty title", emptyTitleConsumable, false},
		{"conflicting title", conflictingConsumable, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateConsumable(tc.consumable)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateConsumable() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateConsumable() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteConsumable(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteConsumable")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestDeleteConsumable: failed during setup. CreateConsumable() failed: %v", err)
	}

	testCases := []struct {
		name          string
		consumableId  string
		shouldSucceed bool
	}{
		{"valid consumable deletion", createdConsumable.Id.String(), true},
		{"invalid consumable ID", "invalid", false},
		{"nil consumable ID", uuid.Nil.String(), false},
		{"empty consumable ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteConsumable(tc.consumableId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteConsumable() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteConsumable() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateConsumableWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateConsumableWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithTask: failed during setup. CreateConsumable() failed: %v", err)
	}

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	t1 := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), t1)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	cq, err := app.AssociateConsumableWithTask(createdAsset.Id.String(), createdTask.Id.String(), createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"})
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithTask: failed during setup. AssociateConsumableWithTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		taskId        string
		consumableId  string
		cq            tp.ConsumableQuantity
		shouldSucceed bool
	}{
		{"valid consumable disassociation", createdTask.Id.String(), createdConsumable.Id.String(), cq, true},
		{"invalid task ID", "invalid", createdConsumable.Id.String(), cq, false},
		{"invalid consumable ID", createdTask.Id.String(), "invalid", cq, false},
		{"non-existent consumable", createdTask.Id.String(), uuid.New().String(), cq, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateConsumableWithTask(createdAsset.Id.String(), tc.taskId, tc.consumableId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateConsumableWithTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateConsumableWithTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateConsumableWithWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateConsumableWithWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithWorkOrder: failed during setup. CreateConsumable() failed: %v", err)
	}

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo1 := utest.SetupWorkOrder(1, createdAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(createdAsset.Id.String(), wo1)
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	cq, err := app.AssociateConsumableWithWorkOrder(createdAsset.Id.String(), createdWorkOrder.Id.String(), createdConsumable.Id.String(), tp.ConsumableQuantity{Quantity: "10"})
	if err != nil {
		t.Errorf("TestDisassociateConsumableWithWorkOrder: failed during setup. AssociateConsumableWithWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		workOrderId   string
		consumableId  string
		cq            tp.ConsumableQuantity
		shouldSucceed bool
	}{
		{"valid consumable disassociation", createdWorkOrder.Id.String(), createdConsumable.Id.String(), cq, true},
		{"invalid work order ID", "invalid", createdConsumable.Id.String(), cq, false},
		{"invalid consumable ID", createdWorkOrder.Id.String(), "invalid", cq, false},
		{"non-existent consumable", createdWorkOrder.Id.String(), uuid.New().String(), cq, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateConsumableWithWorkOrder(createdAsset.Id.String(), tc.workOrderId, tc.consumableId)
			if tc.shouldSucceed && err != nil {

				t.Errorf("DisassociateConsumableWithWorkOrder() failed: %v", err)
			}
		})
	}
}

func TestGetConsumable(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetConsumable")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestGetConsumable: failed during setup. CreateConsumable() failed: %v", err)
	}

	testCases := []struct {
		name          string
		consumableId  string
		shouldSucceed bool
	}{
		{"valid consumable", createdConsumable.Id.String(), true},
		{"invalid consumable ID", "invalid", false},
		{"nil consumable ID", uuid.Nil.String(), false},
		{"empty consumable ID", "", false},
		{"non-existent consumable", uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetConsumable(tc.consumableId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetConsumable() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetConsumable() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListConsumables(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListConsumables")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	_, err = app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestListConsumables: failed during setup. CreateConsumable() failed: %v", err)
	}

	c = utest.SetupConsumable(2, false)
	_, err = app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestListConsumables: failed during setup. CreateConsumable() failed: %v", err)
	}

	testCases := []struct {
		name          string
		count         int
		shouldSucceed bool
	}{
		{"valid list", 2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs, err := app.ListConsumables()
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListConsumables() failed: %v", err)
				} else {
					if len(cs) != tc.count {
						t.Errorf("ListConsumables() failed: expected %d consumables, got %d", tc.count, len(cs))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListConsumables() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateConsumable(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateConsumable")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	consumableCount := 5
	var ids []string
	consumables := make(map[string]tp.Consumable)
	nilIdConsumables := make(map[string]tp.Consumable)
	for i := 0; i < consumableCount; i++ {
		c := utest.SetupConsumable(i, false)
		cc, err := app.CreateConsumable(c)
		if err != nil {
			t.Errorf("TestUpdateConsumable: failed during setup. CreateConsumable() failed: %v", err)
		}

		ids = append(ids, cc.Id.String())
		consumables[cc.Id.String()] = cc
		nilIdConsumables[cc.Id.String()] = c
	}

	testCases := []struct {
		name          string
		consumableId  string
		consumable    tp.Consumable
		title         string
		shouldSucceed bool
	}{
		{"valid consumable with matching IDs", ids[0], consumables[ids[0]], "valid title1", true},
		{"valid consumable with Consumable.Id nil", ids[1], nilIdConsumables[ids[1]], "valid title2", true},
		{"mismatching consumable ID and Consumable.Id", ids[2], consumables[ids[3]], "valid title3", false},
		{"non-existent consumable", uuid.New().String(), tp.Consumable{}, "valid title3", false},

		{"invalid consumable ID", "invalid", tp.Consumable{}, "valid title3", false},
		{"nil consumable ID", uuid.Nil.String(), tp.Consumable{}, "valid title3", false},
		{"empty consumable ID", "", tp.Consumable{}, "valid title3", false},
		{"conflicting id", ids[4], consumables[ids[3]], "valid title3", false},

		{"empty title", ids[1], consumables[ids[1]], "", false},
		{"minimum length title", ids[1], consumables[ids[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", ids[1], consumables[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", ids[1], consumables[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
		{"conflicting title", ids[2], consumables[ids[2]], consumables[ids[3]].Title, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.consumable.Title = tc.title
			_, err := app.UpdateConsumable(tc.consumableId, tc.consumable)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateConsumable() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateConsumable() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateConsumable(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateConsumable")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name          string
		consumable    tp.Consumable
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid consumable", utest.SetupConsumable(1, false), uuid.New(), "valid title", true},
		{"nil id", utest.SetupConsumable(2, false), uuid.Nil, "valid title", false},

		{"empty title", utest.SetupConsumable(3, false), uuid.New(), "", false},
		{"minimum length title", utest.SetupConsumable(4, false), uuid.New(), strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", utest.SetupConsumable(5, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", utest.SetupConsumable(6, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.consumable.Id = tc.id
			tc.consumable.Title = tc.title
			err := app.validateConsumable(tc.consumable)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateConsumable() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateConsumable() should have failed with %s", tc.name)
			}
		})
	}
}

func TestConsumableExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestConsumableExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestConsumableExists: failed during setup. CreateConsumable() failed: %v", err)
	}

	testCases := []struct {
		name          string
		consumableId  string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid consumable", createdConsumable.Id.String(), true, true},
		{"non-existent consumable", uuid.New().String(), false, true},

		{"invalid consumable ID", "invalid", false, false},
		{"nil consumable ID", uuid.Nil.String(), false, false},
		{"empty consumable ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.consumableExists(tc.consumableId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("consumableExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("consumableExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("consumableExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
