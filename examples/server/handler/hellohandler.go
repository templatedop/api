package handler

import (
	"github.com/guregu/null/zero"
	perror "github.com/templatedop/api/errors"
	"github.com/templatedop/api/examples/server/core/domain"
	er "github.com/templatedop/api/examples/server/core/port"
	r "github.com/templatedop/api/examples/server/repo"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/modules/server/handler"
	"github.com/templatedop/api/modules/server/route"
)

type HelloHandler struct {
	*handler.Base
	svc *Service
	ser *r.UserRepository
	log *log.Logger
}

func NewHelloHandler(svc *Service, ser *r.UserRepository, log *log.Logger) *HelloHandler {
	c := handler.New("Hello Handler").SetPrefix("/hello")
	return &HelloHandler{
		Base: c,
		svc:  svc,
		ser:  ser,
		log:  log,
	}
}

func (c *HelloHandler) Routes() []route.Route {
	return []route.Route{
		route.Get("/greet", c.greetHandler).Name("Greet Route"),
		route.Get("/query", c.greetWithQuery).Name("Greet with query"),
		route.Get("/param/:text", c.greetWithParam).Name("Greet with param"),
		route.Post("/body", c.greetWithBody).Name("Greet with body"),
		route.Post("/register", c.register).Name("Register"),
		route.Get("/users", c.ListUsers).Name("List Users"),
		route.Get("/simulate-error", c.testcustomcode2).Name("Simulate Error"),
	}
}

func (c *HelloHandler) testcustomcode2(ctx *route.Context, _ any) (any, error) {
	ctx.Log.Info("Simulate Error called")
	err := perror.NewCode(er.CustomCodeDatabaseError, "could not initialize the service")
	//gerr:=err.(*perror.Error)

	//fmt.Println("gerr stack:",perror.Stack(err))
	//fmt.Println("gerr stack:",gerr.Stack())
	ctx.Log.Info("end of simulation error")

	return nil, err

}

func (c *HelloHandler) greetHandler(ctx *route.Context, req any) (any, error) {
	ctx.Log.Info("Greet Handler called")
	return "Greetings!", nil
}

func (c *HelloHandler) greetWithBody(_ *route.Context, req GreetWithBodyRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}
func (c *HelloHandler) register(ctx *route.Context, req RegisterRequest) (u domain.User, err error) {

	user := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: zero.StringFrom(req.CreatedAt),
	}
	//u,e:=
	u, e := c.ser.CreateUser(ctx, user)
	if e != nil {
		err := perror.WrapCode(er.CustomCodeDatabaseError, e, "could not create user")
		//err := perror.NewCode(CodeDatabaseError, "could not create user")
		return domain.User{}, err
	}

	return u, nil
	// if e!=nil{
	// 	return u,e
	// }
	// return u.ToResponse(),nil

}

// type meta struct {
// 	Total uint64 `json:"total" example:"100"`
// 	Limit uint64 `json:"limit" example:"10"`
// 	Skip  uint64 `json:"skip" example:"0"`
// }

type meta struct {
	Total int `json:"total" example:"100"`
	Limit int `json:"limit" example:"10"`
	Skip  int `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip int) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

func (c *HelloHandler) ListUsers(ctx *route.Context, req ListUsersRequest) (users []*domain.User, err error) {

	ctx.Log.Info("List Users called")

	// user := domain.User{
	// 	Name:      req.Name,
	// 	Email:     req.Email,
	// 	Password:  req.Password,
	// 	CreatedAt: zero.StringFrom(req.CreatedAt),
	// }
	//users1,e:=
	//u,e:=
	return c.ser.ListUsers(ctx, req.Skip, req.Limit)
	//u.ToResponse()
	// if e!=nil{
	// 	return nil,e
	// }

	// total := uint64(len(users1))
	// meta := newMeta(total, req.Limit, req.Skip)
	// rsp := toMap(meta, users1, "users")

}

func (c *HelloHandler) greetWithQuery(_ *route.Context, req GreetWithQueryRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}

func (c *HelloHandler) greetWithParam(_ *route.Context, req GreetWithParamRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}
