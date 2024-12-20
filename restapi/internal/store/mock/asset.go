package mock

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

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
