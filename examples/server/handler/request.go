package handler

type GreetWithBodyRequest struct {
	Text           string `json:"text" validate:""`
	RepeatTimes    int    `form:"repeatTimes" validate:"required"`
	NecessaryParam string `query:"necessaryParam" validate:"required"`
	//FormdataParam  string `form:"formdataParam" validate:"required"`
}

type RegisterRequest struct {
	//Name     string `json:"name" validate:"omitempty,required,min=5" u:"N1" db:"name" example:"John Doe"`
	Name        string `json:"name" validate:"omitempty,min=5" u:"N1" db:"name" example:"John Doe"`
	Email       string `json:"email" validate:"required,email" db:"email" example:"test@example.com"`
	Password    string `json:"password" validate:"required,min=8" u:"P1" db:"password" example:"12345678"`
	Check       int    `json:"check" validate:"required"`
	CreatedAt   string `json:"created_at" `
	CreatedTime string ` json:"created_time"  db:"created_time"  validate:"required"`
}
type ListUsersRequest struct {
	Skip  uint64 `query:"skip" validate:"required,min=0" example:"0"`
	Limit uint64 `query:"limit" validate:"required,min=5" example:"5"`
}
type GreetWithQueryRequest struct {
	Text        string `query:"text" validate:"required,min=3,max=100,queryTextValidator"`
	RepeatTimes int    `query:"repeatTimes" validate:"required,gte=1,lte=5"`
}

type GreetWithParamRequest struct {
	Text        string `param:"text" validate:"required,min=3,max=100,queryTextValidator"`
	RepeatTimes int    `query:"repeatTimes" validate:"required,gte=1,lte=5"`
}
