package implementation

import (
	"time"

	api "github.com/Planck1858/golang101/examples/grpc-server/gen/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	Task  TodoType = "task"
	Event TodoType = "event"
)

type (
	Todo struct {
		ID        string
		Name      string
		Type      TodoType
		CreatedAt time.Time
	}

	TodoType string
)

func (t TodoType) String() string {
	return string(t)
}

func TodoToProto(t Todo) *api.Todo {
	return &api.Todo{
		ID:        t.ID,
		Name:      t.Name,
		Type:      TodoTypeToProto(t.Type),
		CreatedAt: timestamppb.New(t.CreatedAt),
	}
}

func TodoTypeToProto(t TodoType) api.TodoType {
	return api.TodoType(api.TodoType_value[t.String()])
}

func TodoTypeFromProto(t api.TodoType) TodoType {
	switch TodoType(t.String()) {
	case Task:
		return Task
	case Event:
		return Event
	}

	return ""
}
