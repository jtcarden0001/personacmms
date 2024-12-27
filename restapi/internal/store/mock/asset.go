package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

var assetTable = "assets"

func (m *MockStore) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[assetTable] == nil {
		m.data[assetTable] = make(map[uuid.UUID]interface{})
	}

	m.data[assetTable][asset.Id] = asset
	return asset, nil
}

func (m *MockStore) DeleteAsset(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[assetTable][id]; ok {
		delete(m.data[assetTable], id)
		return nil
	}

	return ae.New(ae.CodeNotFound, "deleteasset - asset not found")
}

func (m *MockStore) GetAsset(id uuid.UUID) (tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if asset, ok := m.data[assetTable][id]; ok {
		return asset.(tp.Asset), nil
	}

	return tp.Asset{}, ae.New(ae.CodeNotFound, "getasset - asset not found")
}

func (m *MockStore) ListAssets() ([]tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var assets []tp.Asset
	for _, asset := range m.data[assetTable] {
		assets = append(assets, asset.(tp.Asset))
	}

	return assets, nil
}

func (m *MockStore) UpdateAsset(asset tp.Asset) (tp.Asset, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[assetTable][asset.Id]; !ok {
		return tp.Asset{}, ae.New(ae.CodeNotFound, "updateasset - asset not found")
	}

	m.data[assetTable][asset.Id] = asset
	return asset, nil
}
