package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// ConverToUUID converts a string to uuid
func ConvertToUUID(uuidStr string) (*uuid.UUID, error) {
	uuidObj, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to uuid: %w", err)
	}
	return &uuidObj, nil
}

// ConvertListToUUIDs converts a list of string to a list of uuid
func ConvertListToUUIDs(uuidStrList []string) ([]uuid.UUID, error) {
	var uuidObjList []uuid.UUID

	for _, uuidStr := range uuidStrList {
		uuidObj, err := ConvertToUUID(uuidStr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to uuid: %w", err)
		}

		uuidObjList = append(uuidObjList, *uuidObj)
	}

	return uuidObjList, nil
}
