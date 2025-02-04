package adapter

import (
	"ddd-sample/internal/core/aggregate"
	"ddd-sample/internal/core/repository"
	"ddd-sample/pkg/log"
)

type coreRepository struct {
}

// 工廠
func NewCoreRepository() repository.CoreRepository {
	return &coreRepository{}
}

func (repo *coreRepository) PubEvent(ea aggregate.CoreAggregate) error {
	events := ea.PopEvents()
	for _, e := range events {
		jsonString, err := e.ParseToJSON()
		if err != nil {
			return err
		}

		log.Debug("Event: ", jsonString)
	}

	return nil
}
