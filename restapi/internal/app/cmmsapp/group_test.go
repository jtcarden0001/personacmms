package cmmsapp

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingGroup := utest.SetupGroup(1, false)
	_, err = app.CreateGroup(conflictingGroup)
	if err != nil {
		t.Errorf("TestCreateGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	emptyTitleGroup := utest.SetupGroup(2, false)
	emptyTitleGroup.Title = ""

	testCases := []struct {
		name          string
		group         tp.Group
		shouldSucceed bool
	}{
		{"valid group", utest.SetupGroup(3, false), true},
		{"non nil id", utest.SetupGroup(4, true), false},
		{"empty title", emptyTitleGroup, false},
		{"conflicting title", conflictingGroup, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateGroup(tc.group)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestDeleteGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	testCases := []struct {
		name          string
		groupId       string
		shouldSucceed bool
	}{
		{"valid group deletion", createdGroup.Id.String(), true},
		{"invalid group ID", "invalid", false},
		{"nil group ID", uuid.Nil.String(), false},
		{"empty group ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteGroup(tc.groupId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestGetGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	testCases := []struct {
		name          string
		groupId       string
		shouldSucceed bool
	}{
		{"valid group", createdGroup.Id.String(), true},
		{"invalid group ID", "invalid", false},
		{"nil group ID", uuid.Nil.String(), false},
		{"empty group ID", "", false},
		{"non-existent group", uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetGroup(tc.groupId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListGroups(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListGroups")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	g := utest.SetupGroup(1, false)
	_, err = app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestListGroups: failed during setup. CreateGroup() failed: %v", err)
	}

	g = utest.SetupGroup(2, false)
	_, err = app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestListGroups: failed during setup. CreateGroup() failed: %v", err)
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
			gs, err := app.ListGroups()
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListGroups() failed: %v", err)
				} else {
					if len(gs) != tc.count {
						t.Errorf("ListGroups() failed: expected %d groups, got %d", tc.count, len(gs))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListGroups() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListGroupsByAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListGroupsByAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name                 string
		groupCount           int
		associatedGroupCount int
		shouldSucceed        bool
	}{
		{"list 2 associated groups out of 3 total", 3, 2, true},
	}

	for _, tc := range testCases {
		asset := utest.SetupAsset(1, false)
		ca, err := app.CreateAsset(asset)
		if err != nil {
			t.Errorf("TestListGroupsByAsset: failed during setup. CreateAsset() failed: %v", err)
		}

		createdGroupIds := []uuid.UUID{}
		for i := 0; i < tc.groupCount; i++ {
			g := utest.SetupGroup(i, false)
			cg, err := app.CreateGroup(g)
			if err != nil {
				t.Errorf("TestListGroupsByAsset: failed during setup. CreateGroup() failed: %v", err)
			}

			if i < tc.associatedGroupCount {
				_, err = app.AssociateAssetWithGroup(ca.Id.String(), cg.Id.String())
				if err != nil {
					t.Errorf("TestListGroupsByAsset: failed during setup. AssociateGroupToAsset() failed: %v", err)
				}
			}

			createdGroupIds = append(createdGroupIds, cg.Id)
		}

		gps, err := app.ListGroupsByAsset(ca.Id.String())
		if tc.shouldSucceed {

			if err != nil {
				t.Errorf("ListGroupsByAsset() failed: %v", err)
			} else {
				if len(gps) != tc.associatedGroupCount {
					t.Errorf("ListGroupsByAsset() failed: expected %d groups, got %d", tc.associatedGroupCount, len(gps))
				}
			}
		}

		if !tc.shouldSucceed && err == nil {
			t.Errorf("ListGroupsByAsset() should have failed with %s", tc.name)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	groupCount := 5
	var ids []string
	groups := make(map[string]tp.Group)
	nilIdGroups := make(map[string]tp.Group)
	for i := 0; i < groupCount; i++ {
		g := utest.SetupGroup(i, false)
		cg, err := app.CreateGroup(g)
		if err != nil {
			t.Errorf("TestUpdateGroup: failed during setup. CreateGroup() failed: %v", err)
		}

		ids = append(ids, cg.Id.String())
		groups[cg.Id.String()] = cg
		nilIdGroups[cg.Id.String()] = g
	}

	testCases := []struct {
		name          string
		groupId       string
		group         tp.Group
		title         string
		shouldSucceed bool
	}{
		{"valid group with matching IDs", ids[0], groups[ids[0]], "valid title1", true},
		{"valid group with Group.Id nil", ids[1], nilIdGroups[ids[1]], "valid title2", true},
		{"mismatching group ID and Group.Id", ids[2], groups[ids[3]], "valid title3", false},
		{"non-existent group", uuid.New().String(), tp.Group{}, "valid title3", false},

		{"invalid group ID", "invalid", tp.Group{}, "valid title3", false},
		{"nil group ID", uuid.Nil.String(), tp.Group{}, "valid title3", false},
		{"empty group ID", "", tp.Group{}, "valid title3", false},
		{"conflicting id", ids[4], groups[ids[3]], "valid title3", false},

		{"empty title", ids[1], groups[ids[1]], "", false},
		{"minimum length title", ids[1], groups[ids[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", ids[1], groups[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", ids[1], groups[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
		{"conflicting title", ids[2], groups[ids[2]], groups[ids[3]].Title, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.group.Title = tc.title
			_, err := app.UpdateGroup(tc.groupId, tc.group)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name          string
		group         tp.Group
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid group", utest.SetupGroup(1, false), uuid.New(), "valid title", true},
		{"nil id", utest.SetupGroup(2, false), uuid.Nil, "valid title", false},

		{"empty title", utest.SetupGroup(3, false), uuid.New(), "", false},
		{"minimum length title", utest.SetupGroup(4, false), uuid.New(), strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", utest.SetupGroup(5, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", utest.SetupGroup(6, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.group.Id = tc.id
			tc.group.Title = tc.title
			err := app.validateGroup(tc.group)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGroupExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGroupExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestGroupExists: failed during setup. CreateGroup() failed: %v", err)
	}

	testCases := []struct {
		name          string
		groupId       string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid group", createdGroup.Id.String(), true, true},
		{"non-existent group", uuid.New().String(), false, true},

		{"invalid group ID", "invalid", false, false},
		{"nil group ID", uuid.Nil.String(), false, false},
		{"empty group ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.groupExists(tc.groupId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("groupExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("groupExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("groupExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
