package mock

import (
	"API-GO/domain"

	"github.com/google/uuid"
)

func MockUsers() []domain.USER {
	return []domain.USER{
		{
			ID:        uuid.New(),
			FirstName: "mock_name",
			LastName:  "mock_last_name",
			Email:     "mock@email.com",
			// CreatedAt: time.Now(),
		},
	}
}
