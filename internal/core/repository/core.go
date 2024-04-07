package repository

import "ddd-sample/internal/core/aggregate"

type CoreRepository interface {
	PubEvent(aggregate.CoreAggregate) error
}
