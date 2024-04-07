package entity

import "ddd-sample/pkg/uid"

type CoreEntity struct {
	uid string
}

func NewCoreEntity() CoreEntity {
	return CoreEntity{
		uid: uid.NewNanoID(),
	}
}

func BuildCoreEntity(uid string) CoreEntity {
	return CoreEntity{
		uid: uid,
	}
}

func (e CoreEntity) UID() string {
	return e.uid
}
