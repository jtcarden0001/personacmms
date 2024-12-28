package cmmsapp

import (
	"strings"
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	creA, err := app.CreateAsset(a)
	if err != nil {
		t.Fatalf("CreateAsset() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, a, creA, diffFields)
}

func TestValidateAssetWithNilId(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	err := app.validateAsset(a)
	if err == nil {
		t.Errorf("ValidateAsset() should have returned an error with nil ID")
	}
}

func TestValidateAssetWithEmptyTitle(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	a.Title = ""
	err := app.validateAsset(a)
	if err == nil {
		t.Errorf("ValidateAsset() should have returned an error with empty Title")
	}
}

func TestValidateAssetWithTooLongTitle(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	maxTitleLen := 255
	a.Title = strings.Join(make([]string, maxTitleLen+1), "a")
	err := app.validateAsset(a)
	if err == nil {
		t.Errorf("ValidateAsset() should have returned an error with too long Title")
	}
}

func TestAssetExistsWithPresentAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	creA, err := app.CreateAsset(a)
	if err != nil {
		t.Fatalf("AssetExists() failed during setup. CreateAsset() failed: %v", err)
	}

	exists, err := app.assetExists(creA.Id)
	if err != nil {
		t.Errorf("AssetExists() failed: %v", err)
	}

	if !exists {
		t.Errorf("AssetExists() should have returned true with present asset")
	}
}

func TestAssetExistsWithAbsentAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	exists, err := app.assetExists(a.Id)
	if err != nil {
		t.Errorf("AssetExists() failed: %v", err)
	}

	if exists {
		t.Errorf("AssetExists() should have returned false with absent asset")
	}
}

func TestDeleteAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Fatalf("TestAssetDelete: failed during setup. CreateAsset() failed: %v", err)
	}

	err = app.DeleteAsset(createdAsset.Id.String())
	if err != nil {
		t.Fatalf("TestAssetDelete: DeleteAsset() failed: %v", err)
	}

	_, err = app.GetAsset(createdAsset.Id.String())
	if err == nil {
		t.Errorf("TestAssetDelete: GetAsset() returned nil error after deletion")
	}
}

func TestGetAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	creA, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetGet: failed during setup. CreateAsset() failed: %v", err)
	}

	getA, err := app.GetAsset(a.Id.String())
	if err != nil {
		t.Errorf("TestAssetGet: GetAsset() failed: %v", err)
	}

	utest.CompEntities(t, creA, getA)
}

func TestGetAssetWithAbsentAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)

	_, err := app.GetAsset(a.Id.String())
	if err == nil {
		t.Errorf("TestAssetGet: GetAsset() should have returned an error with absent asset")
	}
}

func TestListAssets(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	b := utest.SetupAsset(2, false)
	_, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetList: failed during setup. CreateAsset() failed: %v", err)
	}

	_, err = app.CreateAsset(b)
	if err != nil {
		t.Errorf("TestAssetList: failed during setup. CreateAsset() failed: %v", err)
	}

	assets, err := app.ListAssets()
	if err != nil {
		t.Errorf("TestAssetList: ListAssets() failed: %v", err)
	}

	if len(assets) != 2 {
		t.Errorf("TestAssetList: expected 1 asset, got %v", len(assets))
	}
}

func TestUpdateAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	creA, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetUpdate: failed during setup. CreateAsset() failed: %v", err)
	}

	a.Title = "New Name"
	updatedAsset, err := app.UpdateAsset(creA.Id.String(), a)
	if err != nil {
		t.Errorf("UpdateAsset() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Title"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, a, updatedAsset, diffFields)
}
