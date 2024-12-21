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
	m.data["assets"][asset.Title] = asset
	return asset, nil
}

func (m *MockStore) DeleteAsset(group, title string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["assets"][title]; !ok {
		return ae.New(ae.CodeNotFound, "asset not found")
	}
	delete(m.data["assets"], title)
	return nil
}

func (m *MockStore) GetAsset(group, title string) (tp.Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if asset, ok := m.data["assets"][title]; ok {
		return asset.(tp.Asset), nil
	}
	return tp.Asset{}, ae.ErrNotFound
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

func (m *MockStore) UpdateAsset(group, title string, asset tp.Asset) (tp.Asset, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data["assets"][title]; !ok {
		return tp.Asset{}, ae.New(ae.CodeNotFound, "asset not found")
	}
	// delete the old asset
	delete(m.data["assets"], title)
	m.data["assets"][asset.Title] = asset
	return asset, nil
}
