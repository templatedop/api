package middlewares

import (
	"context"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/templatedop/api/ecode"
	perror "github.com/templatedop/api/errors"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/modules/server/response"
	"github.com/templatedop/api/modules/server/validation"
)



func unwrapError(err error) error {
	for {
		unwrappedErr := perror.Unwrap(err)
		if unwrappedErr == nil {
			return err
		}
		err = unwrappedErr
	}
}

func ErrHandler(log *log.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		cc := c.UserContext()
		RequestID := cc.Value(RequestIDContextKey).(string)
		log.ToZerolog().Error().Str("RequestID:", RequestID).Str("Stack", perror.Stack(err)).Msg("Error handler stack")
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(response.Error("Request Timedout", nil))
		}
		if gerr, ok := err.(*perror.Error); ok {

			unwrappedErr := unwrapError(gerr)

			if pgErr, ok := unwrappedErr.(*pgconn.PgError); ok {
				status, e := handledbError(pgErr.Code)
				parsedCode, _ := strconv.Atoi(pgErr.Code)
				ers := []response.Errors{
					{Code: perror.Code(gerr).Code(), Message: perror.Code(gerr).Message()},
					{Code: parsedCode, Message: e.Error()},
				}
				return c.Status(status).JSON(response.Error(gerr.Current().Error(), ers))

			}
			if unwrappedErr != nil {
				return perrnondb(c, unwrappedErr.(*perror.Error))

			}

			return perrnondb(c, gerr)

		}

		if verr, ok := err.(*validation.Error); ok {

			return c.Status(GetCodeFromErr(verr)).
				JSON(response.Error(verr.Error(), nil, getValidationErrs(verr)...))

		} else if pgErr, ok := err.(*pgconn.PgError); ok {
			status, e := handledbError(pgErr.Code)
			parsedCode, _ := strconv.Atoi(pgErr.Code)
			ers := []response.Errors{
				{Code: parsedCode, Message: e.Error()},
			}
			emessage := "internal Error"
			return c.Status(status).JSON(response.Error(emessage, ers))
		}

		// find type of error
		errcode := GetCodeFromErr(err)
		if errcode < 500 {
			return c.Status(errcode).JSON(response.Error(err.Error(), nil))
		}

		return handleGenericError(c, err)

	}
}

func perrnondb(c *fiber.Ctx, gerr *perror.Error) error {
	status, e := getperrorcode(gerr)
	if perror.Code(gerr).Code() == perror.Code(e).Code() {
		ers := []response.Errors{
			{Code: perror.Code(gerr).Code(), Message: perror.Code(gerr).Message()},
		}
		return c.Status(status).JSON(response.Error(gerr.Current().Error(), ers))
	}
	ers := []response.Errors{
		{Code: perror.Code(gerr).Code(), Message: perror.Code(gerr).Message()},
		{Code: perror.Code(e).Code(), Message: perror.Code(e).Message()},
	}

	return c.Status(status).JSON(response.Error(gerr.Current().Error(), ers))

}

func getperrorcode(e *perror.Error) (i int, er error) {
	switch e.Code() {
	case ecode.CodeMalformedRequest:
		return fiber.StatusUnprocessableEntity, perror.NewCode(e.Code(), "Request malformed")
	}
	return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeInternalError, "Internal error")
}

func getValidationErrs(err error) []validation.FieldError {
	if vErr := new(validation.Error); errors.As(err, &vErr) {
		return vErr.FieldErrors()
	}
	return nil
}

func GetCodeFromErr(err error) int {
	if fErr := new(fiber.Error); errors.As(err, &fErr) {
		return fErr.Code
	}

	return fiber.StatusInternalServerError
}

func pgerr(err string) (int, error) {
	return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, err)
}

func handledbError(sqlState string) (status int, e error) {

	switch {
	case pgerrcode.IsCardinalityViolation(sqlState):
		return pgerr("Cardinality violation")
	case pgerrcode.IsWarning(sqlState):
		return pgerr("Warning")
	case pgerrcode.IsNoData(sqlState):
		return pgerr("No data found")
	case pgerrcode.IsSQLStatementNotYetComplete(sqlState):
		return pgerr("SQL statement not yet complete")
	case pgerrcode.IsConnectionException(sqlState):
		return pgerr("Connection exception")
	case pgerrcode.IsTriggeredActionException(sqlState):
		return pgerr("Triggered action exception")
	case pgerrcode.IsFeatureNotSupported(sqlState):
		return pgerr("Feature not supported")
	case pgerrcode.IsInvalidTransactionInitiation(sqlState):
		return pgerr("Invalid transaction initiation")
	case pgerrcode.IsLocatorException(sqlState):
		return pgerr("Locator exception")
	case pgerrcode.IsInvalidGrantor(sqlState):
		return pgerr("Invalid grantor")
	case pgerrcode.IsInvalidRoleSpecification(sqlState):
		return pgerr("Invalid role specification")
	case pgerrcode.IsDiagnosticsException(sqlState):
		return pgerr("Diagnostics exception")
	case pgerrcode.IsCaseNotFound(sqlState):
		return pgerr("Case not found")
	case pgerrcode.IsCardinalityViolation(sqlState):
		return pgerr("Cardinality violation")
	case pgerrcode.IsDataException(sqlState):
		return pgerr("Data exception")
	case pgerrcode.IsIntegrityConstraintViolation(sqlState):
		return pgerr("Integrity constraint violation")
	case pgerrcode.IsInvalidCursorState(sqlState):
		return pgerr("Invalid cursor state")
	case pgerrcode.IsInvalidTransactionState(sqlState):
		return pgerr("Invalid transaction state")
	case pgerrcode.IsInvalidSQLStatementName(sqlState):
		return pgerr("Invalid SQL statement name")
	case pgerrcode.IsTriggeredDataChangeViolation(sqlState):
		return pgerr("Triggered data change violation")
	case pgerrcode.IsInvalidAuthorizationSpecification(sqlState):
		return pgerr("Invalid authorization specification")
	case pgerrcode.IsDependentPrivilegeDescriptorsStillExist(sqlState):
		return pgerr("Dependent privilege descriptors still exist")
	case pgerrcode.IsInvalidTransactionTermination(sqlState):
		return pgerr("Invalid transaction termination")
	case pgerrcode.IsSQLRoutineException(sqlState):
		return pgerr("SQL routine exception")
	case pgerrcode.IsInvalidCursorName(sqlState):
		return pgerr("Invalid cursor name")
	case pgerrcode.IsExternalRoutineException(sqlState):
		return pgerr("External routine exception")
	case pgerrcode.IsExternalRoutineInvocationException(sqlState):
		return pgerr("External routine invocation exception")
	case pgerrcode.IsSavepointException(sqlState):
		return pgerr("Savepoint exception")
	case pgerrcode.IsInvalidCatalogName(sqlState):
		return pgerr("Invalid catalog name")
	case pgerrcode.IsInvalidSchemaName(sqlState):
		return pgerr("Invalid schema name")
	case pgerrcode.IsTransactionRollback(sqlState):
		return pgerr("Transaction rollback")
	case pgerrcode.IsSyntaxErrororAccessRuleViolation(sqlState):
		return pgerr("Syntax error or access rule violation")
	case pgerrcode.IsWithCheckOptionViolation(sqlState):
		return pgerr("With check option violation")
	case pgerrcode.IsInsufficientResources(sqlState):
		return pgerr("Insufficient resources")
	case pgerrcode.IsProgramLimitExceeded(sqlState):
		return pgerr("Program limit exceeded")
	case pgerrcode.IsObjectNotInPrerequisiteState(sqlState):
		return pgerr("Object not in prerequisite state")
	case pgerrcode.IsOperatorIntervention(sqlState):
		return pgerr("Operator intervention")
	case pgerrcode.IsSystemError(sqlState):
		return pgerr("System error")
	case pgerrcode.IsSnapshotFailure(sqlState):
		return pgerr("Snapshot failure")
	case pgerrcode.IsConfigurationFileError(sqlState):
		return pgerr("Configuration file error")
	case pgerrcode.IsForeignDataWrapperError(sqlState):
		return pgerr("Foreign data wrapper error")
	case pgerrcode.IsPLpgSQLError(sqlState):
		return pgerr("PL/pgSQL error")
	default:
		return pgerr("Unknown database error")

	}

}

func handleGenericError(c *fiber.Ctx, err error) error {
	emssage := "internal server error"
	status := GetCodeFromErr(err)
	ers := []response.Errors{
		{Code: 50, Message: "internal error"},
	}
	return c.Status(status).JSON(response.Error(emssage, ers))
}
