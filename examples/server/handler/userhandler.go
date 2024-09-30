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
type UserHandler struct {
	*handler.Base
	svc *Service
	ser *r.UserRepository
	log *log.Logger
}

func NewUserHandler(svc *Service, ser *r.UserRepository, log *log.Logger) *UserHandler {
	//prefix := V1 + "/user"
	//V1 := "/user"
	//c:=V1.AddPrefix("/user")

	c := handler.New("User Handler").SetPrefix(V1).AddPrefix("/user")

	return &UserHandler{
		Base: c,
		svc:  svc,
		ser:  ser,
		log:  log,
	}
}



func (c *UserHandler) Routes() []route.Route {
	return []route.Route{
		//route.GET("/greet", c.greetHandler).Name("Greet Route"),
		//route.GET("/query", c.greetWithQuery).Name("Greet with query"),
		//route.GET("/param/:text", c.greetWithParam).Name("Greet with param"),
		//route.POST("/body", c.greetWithBody).Name("Greet with body"),
		route.POST("/register", c.register).Name("Register"),
		route.GET("/users", c.ListUsers).Name("List Users"),
		//route.GET("/simulate-error", c.testcustomcode2).Name("Simulate Error"),
	}
}

func (c *UserHandler) testcustomcode2(ctx *route.Context, _ any) (any, error) {
	c.log.Debug("Simulate Error called")
	//ctx.Log.Info("Simulate Error called")
	err := perror.NewCode(er.CustomCodeDatabaseError, "could not initialize the service")
	//gerr:=err.(*perror.Error)

	//fmt.Println("gerr stack:",perror.Stack(err))
	//fmt.Println("gerr stack:",gerr.Stack())
	ctx.Log.Info("end of simulation error")

	return nil, err

}

func (c *UserHandler) greetHandler(ctx *route.Context, req any) (any, error) {
	ctx.Log.Info("Greet Handler called")
	return "Greetings!", nil
}

func (c *UserHandler) greetWithBody(_ *route.Context, req GreetWithBodyRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}
func (c *UserHandler) register(ctx *route.Context, req RegisterRequest) (u domain.User, err error) {
	c.log.Debug("Register called")
	c.log.ToZerolog().Debug().Str("Register called", "sda").Caller().Msg("Register called")
	user := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: zero.StringFrom(req.CreatedAt),
	}
	//u,e:=
	u, e := c.ser.CreateUser(ctx, user)
	if e != nil {
		//err:=perror.New(er.CustomCodeDatabaseError.Message())
		//err := perror.NewCode(er.CustomCodeDatabaseError, "could not create user")
		//err:=errors.New("could not create user")

		err := perror.WrapCode(er.CustomCodeDatabaseError, e, "could not create user")

		//u.StatusCode = 201
		//err := perror.NewCode(CodeDatabaseError, "could not create user")
		return domain.User{}, err
	}

	u.StatusCode = 201
	//fmt.Println("User is:", u)

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
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

func (c *UserHandler) ListUsers(ctx *route.Context, req ListUsersRequest) (ul *ListUsersResponse, err error) {

	ctx.Log.Info("List Users called")

	// user := domain.User{
	// 	Name:      req.Name,
	// 	Email:     req.Email,
	// 	Password:  req.Password,
	// 	CreatedAt: zero.StringFrom(req.CreatedAt),
	// }
	//users1,e:=
	//u,e:=
	u, e := c.ser.ListUsers(ctx, req.Skip, req.Limit)
	if e != nil {
		return nil, e
	}
	response := &ListUsersResponse{
		Meta:  newMeta(uint64(len(u)), req.Limit, req.Skip),
		Users: u,
	}

	return response, nil

	//return c.ser.ListUsers(ctx, req.Skip, req.Limit)
	//u.ToResponse()
	// if e!=nil{
	// 	return nil,e
	// }

	// total := uint64(len(users1))
	// meta := newMeta(total, req.Limit, req.Skip)
	// rsp := toMap(meta, users1, "users")

}

func (c *UserHandler) greetWithQuery(_ *route.Context, req GreetWithQueryRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}

func (c *UserHandler) greetWithParam(_ *route.Context, req GreetWithParamRequest) (string, error) {
	return c.svc.GenerateText(req.Text, req.RepeatTimes)
}
