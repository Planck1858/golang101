package implementation

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	api "github.com/Planck1858/golang101/examples/grpc-server/gen/go"
)

func (i *Implementation) CreateTodo(_ context.Context, req *api.CreateTodoRequest) (_ *api.Todo, err error) {
	i.log.Info("CreateTodo started...")
	defer func() {
		if err != nil {
			i.log.Error("CreateTodo error: %v", err)
		} else {
			i.log.Info("CreateTodo success")
		}
	}()

	err = req.ValidateAll()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate all: %v", err)
	}

	todo := Todo{
		ID:        uuid.New().String(),
		Name:      req.GetName(),
		Type:      TodoTypeFromProto(req.GetType()),
		CreatedAt: time.Now(),
	}

	i.store[todo.ID] = todo

	resp := TodoToProto(todo)

	return resp, nil
}

func (i *Implementation) CreateTodoV2(_ context.Context, req *api.CreateTodoRequestV2) (_ *api.Todo, err error) {
	i.log.Info("CreateTodoV2 started...")
	defer func() {
		if err != nil {
			i.log.Error("CreateTodoV2 error: %v", err)
		} else {
			i.log.Info("CreateTodoV2 success")
		}
	}()

	err = req.ValidateAll()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate all: %v", err)
	}

	todo := Todo{
		ID:        uuid.New().String(),
		Name:      req.GetTodo().GetName(),
		Type:      TodoTypeFromProto(req.GetTodo().GetType()),
		CreatedAt: time.Now(),
	}

	i.store[todo.ID] = todo

	resp := TodoToProto(todo)

	return resp, nil
}

func (i *Implementation) GetAllTodo(_ context.Context, req *emptypb.Empty) (_ *api.TodoList, err error) {
	i.log.Info("GetAllTodo started...")
	defer func() {
		if err != nil {
			i.log.Error("GetAllTodo error: %v", err)
		} else {
			i.log.Info("GetAllTodo  success")
		}
	}()

	resp := &api.TodoList{
		Todos: make([]*api.Todo, 0, len(i.store)),
	}

	for _, t := range i.store {
		resp.Todos = append(resp.Todos, TodoToProto(t))
	}

	return resp, nil
}

func (i *Implementation) GetTodoByID(_ context.Context, req *api.GetTodoByIDRequest) (_ *api.Todo, err error) {
	i.log.Info("GetTodoByID started...")
	defer func() {
		if err != nil {
			i.log.Error("GetTodoByID error: %v", err)
		} else {
			i.log.Info("GetTodoByID success")
		}
	}()

	err = req.ValidateAll()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate all: %v", err)
	}

	t, ok := i.store[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo is not found")
	}

	resp := TodoToProto(t)

	return resp, nil
}
func (i *Implementation) GetTodoByIDWithQueryParams(_ context.Context, req *api.GetTodoByIDRequestWithQueryParams) (_ *api.TodoList, err error) {
	i.log.Info("GetTodoByIDWithQueryParams started...")
	defer func() {
		if err != nil {
			i.log.Error("GetTodoByIDWithQueryParams error: %v", err)
		} else {
			i.log.Info("GetTodoByIDWithQueryParams success")
		}
	}()

	if req.GetId() != "" {
		t, ok := i.store[req.GetId()]
		if ok {
			return &api.TodoList{Todos: []*api.Todo{TodoToProto(t)}}, nil
		}
	}

	if req.GetType() != "" {
		if !(req.GetType() == Task.String() || req.GetType() == Event.String()) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid type")
		}

		todos := make([]*api.Todo, 0, len(i.store))
		for _, t := range i.store {
			if t.Type.String() == req.GetType() {
				todos = append(todos, TodoToProto(t))
			}
		}

		return &api.TodoList{Todos: todos}, nil
	}

	return nil, status.Errorf(codes.NotFound, "todos are not found")
}
