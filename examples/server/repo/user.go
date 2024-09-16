package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/templatedop/api/db"
	"github.com/templatedop/api/examples/server/core/domain"
	"github.com/templatedop/api/modules/server/route"
	. "github.com/templatedop/api/repo"

	//"github.com/templatedop/fxdb"
	//. "github.com/templatedop/fxdb"
	logger "github.com/templatedop/api/log"
)

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 */
//  type DB struct {
//     *fxdb.DB
// }
type UserRepository struct {
	Db  *db.DB
	log *logger.Logger
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(Db *db.DB, log *logger.Logger) *UserRepository {
	return &UserRepository{
		Db,
		log,
	}
}

var psql = Psql

// CreateUser creates a new user in the database
// func (ur *UserRepository) CreateUser(gctx *gin.Context, user *domain.User) (*domain.UserDB, error) {
func (ur *UserRepository) CreateUser(pcontext *route.Context, user domain.User) (domain.User, error) {

	ctx, cancel := context.WithTimeout(pcontext.Ctx, 10*time.Second)
	defer cancel()

	//ur.log.Debug("USer:", user)

	query := psql.Insert("users").SetMap(GenerateMapFromStruct(user, "insert")).Suffix("returning *")
	// sql, args, err := query.ToSql()
	// ur.log.Debug("sql:", sql)
	// ur.log.Debug("args:", args)
	// if err != nil {
	// 	ur.log.Debug("Error:", err.Error())
	// 	return err
	// }

	// e, er := ur.Db.Exec(ctx, sql, args...)
	// ur.log.Debug("e:", e)

	p, err := InsertReturning(ctx, ur.Db, query, pgx.RowToStructByName[domain.User])
	// ur.log.Debug("P:", p.Insert())
	// ur.log.Debug("error:", err)
	//ur.log.Debug("user:", p)
	ur.log.Debug("user", p)
	return p, err
	//return Insert(ctx, ur.Db, query, pgx.RowToAddrOfStructByPos[domain.UserDB], ur.log)
}

// GetUserByID gets a user by ID from the database
func (ur *UserRepository) GetUserByID(id uint64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ur.log.Info("Came inside getuser by id")
	var u1 domain.UserDB
	columns := GenerateColumnsFromStruct(u1, "select")
	query := psql.Select(columns...).
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	return SelectOne(ctx, ur.Db, query, pgx.RowToAddrOfStructByNameLax[domain.User])
}

// GetUserByEmailAndPassword gets a user by email from the database
func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//var user domain.User
	query := psql.Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)
	return SelectOneOK(ctx, ur.Db, query, pgx.RowToAddrOfStructByName[domain.User])
}

type PreparedStatement struct {
	Name           string `db:"name"`
	Statement      string
	PrepareTime    time.Time
	ParameterTypes string
	FromSql        bool
	//ResultDataTypes string
}

func (ur *UserRepository) GetPreparestatements() (r []PreparedStatement, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := psql.Select("name, statement, prepare_time, parameter_types, from_sql ").
		From("pg_catalog.pg_prepared_statements ")
	r, e = SelectRows(ctx, ur.Db, query, pgx.RowToStructByPos[PreparedStatement])
	return

}

// ListUsers lists all users from the database
func (ur *UserRepository) ListUsers(pcontext *route.Context, skip, limit uint64) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(pcontext.Ctx, 10*time.Second)
	defer cancel()
	//query := psql.Select("name,email,password").
	query := psql.Select("*").
		From("users").
		OrderBy("id").
		Limit(limit).
		Offset((skip - 1) * limit)
	a, b := SelectRows(ctx, ur.Db, query, pgx.RowToAddrOfStructByNameLax[domain.User])
	//ur.log.Debug("Error:", b)
	//ur.log.Debug("Users:", a)

	return a, b
	// sql, args, err := query.ToSql()
	// if err != nil {
	// 	ur.log.Debug("Logs :", err.Error())
	// 	return nil, err
	// }
	// ur.log.Debug("sql:", sql)
	// ur.log.Debug("args:", args)

	// //s,_:=SelectWithSecondary(ctx,ur.Db,"select",domain.User{},sql,args)

	// //return nil,nil

	// rows, err := ur.Db.Query(ctx, sql, args...)
	// if err != nil {
	// 	ur.log.Error("pgutility, err querying at Select Rows :", err.Error())
	// 	return nil, err

	// }

	// defer rows.Close()
	// //user := []*domain.User{}
	// user, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[domain.User])
	// if err != nil {
	// 	ur.log.Error("pgutility, err collecting Rows at Select Rows :", err.Error())
	// 	return nil, err
	// }
	// ur.log.Debug("Users:", user)
	// return user, nil

	// collectedRows, err := pgx.CollectRows(rows, scanFn)
	// if err != nil {
	// 	log.Error("pgutility, err collecting Rows at Select Rows :", err.Error())
	// 	return nil, err
	// }
	// return SelectRows(ctx, ur.Db, query, RowToStructByTagExp[domain.User], ur.log)
	// if ctx.Err() == context.DeadlineExceeded {
	// 	ur.log.Error("Context deadline exceeded")
	// 	return nil, errors.New("context deadline exceeded")
	// }

	// return a, b
	//return SelectRowsTag[domain.User](ur.Db.ctx, ur.Db, query, ur.log, "select")

}

// UpdateUser updates a user by ID in the database
func (ur *UserRepository) UpdateUser(pcontext context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(pcontext, 10*time.Second)
	defer cancel()
	query := psql.Update("users").
		// Set("name", sq.Expr("COALESCE(?, name)", name)).
		// Set("email", sq.Expr("COALESCE(?, email)", email)).
		// Set("password", sq.Expr("COALESCE(?, password)", password)).
		// Set("role", sq.Expr("COALESCE(?, role)", role)).
		// Set("updated_at", time.Now()).
		SetMap(GenerateMapFromStruct(user, "insert")).
		Where(sq.Eq{"id": user.ID}).
		Suffix("RETURNING *")
	return UpdateReturning(ctx, ur.Db, query, pgx.RowToAddrOfStructByPos[domain.User])

}

// DeleteUser deletes a user by ID from the database
func (ur *UserRepository) DeleteUser(pcontext context.Context, id uint64) (pgconn.CommandTag, error) {
	ctx, cancel := context.WithTimeout(pcontext, 10*time.Second)
	defer cancel()
	query := psql.Delete("users").
		Where(sq.Eq{"id": id})
	p, err := Delete(ctx, ur.Db, query)
	return p, err
}
