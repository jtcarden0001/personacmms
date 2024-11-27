package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateAssetTaskConsumable(atc tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error) {
	return a.db.CreateAssetTaskConsumable(atc)
}

func (a *App) CreateAssetTaskConsumableWithValidation(groupTitle, assetTitle, assetTaskId, consumableId, quantityNote string) (tp.AssetTaskConsumable, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	return a.db.CreateAssetTaskConsumable(tp.AssetTaskConsumable{AssetTaskId: atId, ConsumableId: cId, QuantityNote: quantityNote})
}

func (a *App) DeleteAssetTaskConsumable(groupTitle, assetTitle, assetTaskId, consumableId string) error {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return err
	}

	return a.db.DeleteAssetTaskConsumable(atId, cId)
}

func (a *App) ListAssetTaskConsumables(groupTitle, assetTitle, assetTaskId string) ([]tp.AssetTaskConsumable, error) {
	// TODO: implement validation
	atcs, err := a.db.ListAssetTaskConsumables()
	if err != nil {
		return nil, err
	}

	// TODO: filter asset task consumables by asset task id

	return atcs, nil
}

func (a *App) GetAssetTaskConsumable(groupTitle, assetTitle, assetTaskId, consumableId string) (tp.AssetTaskConsumable, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	return a.db.GetAssetTaskConsumable(atId, cId)
}

func (a *App) UpdateAssetTaskConsumable(atc tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error) {
	return a.db.UpdateAssetTaskConsumable(atc)
}

func (a *App) UpdateAssetTaskConsumableWithValidation(groupTitle, assetTitle, assetTaskId, consumableId, quantityNote string) (tp.AssetTaskConsumable, error) {
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.AssetTaskConsumable{}, err
	}

	return a.db.UpdateAssetTaskConsumable(tp.AssetTaskConsumable{
		AssetTaskId:  atId,
		ConsumableId: cId,
		QuantityNote: quantityNote,
	})
}
