package event

import (
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/uid"
	"ddd-sample/pkg/util"
)

type CoreEvent[T any] struct {
	EventID string `json:"eventID"` // 事件ID
	Name    string `json:"name"`
	Time    string `json:"time"`
	Data    T      `json:"data"`
}

func NewCoreEvent[T any](time string, data T) CoreEvent[T] {
	var t T
	name := util.GetStructName(t)

	return CoreEvent[T]{
		EventID: uid.NewNanoID(),
		Name:    name,
		Time:    time,
		Data:    data,
	}
}

func (e CoreEvent[T]) GetEventID() string {
	return e.EventID
}

func (e CoreEvent[T]) GetName() string {
	return e.Name
}

func (e CoreEvent[T]) ParseToJSON() (string, errorcode.ErrorCode) {
	s, err := util.JSONMarshal(e.Data)
	if err != nil {
		return "", errorcode.ErrJSONMarshal
	}

	return s, nil
}
