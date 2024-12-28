package test

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func SetupAsset(identifier int, generateId bool) tp.Asset {
	year := 2000 + (identifier % 22) // Random year based on identifier
	id := uuid.Nil
	if generateId {
		id = uuid.New()
	}

	return tp.Asset{
		Id:           id,
		Title:        fmt.Sprintf("Asset %d", identifier),
		Year:         ToPtr(year),
		Manufacturer: ToPtr(fmt.Sprintf("Asset %d manufacturer", identifier)),
		Make:         ToPtr(fmt.Sprintf("Asset %d make", identifier)),
		ModelNumber:  ToPtr(fmt.Sprintf("Asset %d model number", identifier)),
		SerialNumber: ToPtr(fmt.Sprintf("Asset %d serial number", identifier)),
		Description:  ToPtr(fmt.Sprintf("Asset %d description", identifier)),
	}
}
