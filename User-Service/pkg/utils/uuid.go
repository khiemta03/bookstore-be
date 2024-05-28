package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func ConvertToUUID(uuidStr string) (*uuid.UUID, error) {
	uuidObj, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to uuid: %w", err)
	}
	return &uuidObj, nil
}
