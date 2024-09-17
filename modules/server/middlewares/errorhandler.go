package middlewares

import (
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
	//"go.uber.org/zap"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
	Stack   string      `json:"stack,omitempty"` // Optionally include stack traces for debugging
}

func ErrHandler(log *log.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		log.ToZerolog().Info().Err(err).Msg("Error handler log")
		log.Info("Error handler log")

		if gerr, ok := err.(*perror.Error); ok {
			if unwrappedErr := perror.Unwrap(gerr); unwrappedErr != nil {
				if pgErr, ok := unwrappedErr.(*pgconn.PgError); ok {
					status, e := handledbError(pgErr.Code)
					parsedCode, _ := strconv.Atoi(pgErr.Code)
					ers := []response.Errors{
						{Code: status, Message: strconv.Itoa(perror.Code(e).Code()) + "-" + perror.Code(e).Message()},
						{Code: parsedCode, Message: e.Error()},
					}
					return c.Status(status).JSON(response.Error(gerr.Current().Error(), ers))
				}

			}
			status, pe := getperrorcode(gerr)

			ers := []response.Errors{
				{Code: status, Message: strconv.Itoa(perror.Code(pe).Code()) + "-" + perror.Code(pe).Message()},
			}

			return c.Status(status).JSON(response.Error(gerr.Current().Error(), ers))

		} else {

			log.ToZerolog().Info().Msg("Unable to parse perror")
		}

		return c.
			Status(getCodeFromErr(err)).
			JSON(response.Error(err.Error(), nil, getValidationErrs(err)...))
	}
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

func getCodeFromErr(err error) int {
	if fErr := new(fiber.Error); errors.As(err, &fErr) {
		return fErr.Code
	}

	return fiber.StatusInternalServerError
}

func handledbError(sqlState string) (status int, e error) {
	switch {
	case pgerrcode.IsCardinalityViolation(sqlState):
		return fiber.StatusBadRequest, perror.NewCode(ecode.CodeDbOperationError, "Cardinality violation")
	case pgerrcode.IsWarning(sqlState):
		return fiber.StatusOK, perror.NewCode(ecode.CodeDbOperationError, "Warning")
	case pgerrcode.IsNoData(sqlState):
		return fiber.StatusNotFound, perror.NewCode(ecode.CodeDbOperationError, "No data found")
	case pgerrcode.IsSQLStatementNotYetComplete(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "SQL statement not  complete")
	case pgerrcode.IsConnectionException(sqlState):
		return fiber.StatusServiceUnavailable, perror.NewCode(ecode.CodeDbOperationError, "Connection exception")
	case pgerrcode.IsTriggeredActionException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Triggered action exception")
	case pgerrcode.IsFeatureNotSupported(sqlState):
		return fiber.StatusNotImplemented, perror.NewCode(ecode.CodeDbOperationError, "Feature not supported")
	case pgerrcode.IsInvalidTransactionInitiation(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid transaction initiation")
	case pgerrcode.IsLocatorException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Locator exception")
	case pgerrcode.IsInvalidGrantor(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid grantor")
	case pgerrcode.IsInvalidRoleSpecification(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid role specification")
	case pgerrcode.IsDiagnosticsException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Diagnostics exception")
	case pgerrcode.IsCaseNotFound(sqlState):
		return fiber.StatusNotFound, perror.NewCode(ecode.CodeDbOperationError, "Case not found")
	case pgerrcode.IsCardinalityViolation(sqlState):
		return fiber.StatusBadRequest, perror.NewCode(ecode.CodeDbOperationError, "Cardinality violation")
	case pgerrcode.IsDataException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Data exception")
	case pgerrcode.IsIntegrityConstraintViolation(sqlState):
		return fiber.StatusConflict, perror.NewCode(ecode.CodeDbOperationError, "Integrity constraint violation")
	case pgerrcode.IsInvalidCursorState(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid cursor state")
	case pgerrcode.IsInvalidTransactionState(sqlState):
		return fiber.StatusConflict, perror.NewCode(ecode.CodeDbOperationError, "Invalid transaction state")
	case pgerrcode.IsInvalidSQLStatementName(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid SQL statement name")
	case pgerrcode.IsTriggeredDataChangeViolation(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Triggered data change violation")
	case pgerrcode.IsInvalidAuthorizationSpecification(sqlState):
		return fiber.StatusUnauthorized, perror.NewCode(ecode.CodeDbOperationError, "Invalid authorization specification")
	case pgerrcode.IsDependentPrivilegeDescriptorsStillExist(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Dependent privilege descriptors still exist")
	case pgerrcode.IsInvalidTransactionTermination(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid transaction termination")
	case pgerrcode.IsSQLRoutineException(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "SQL routine exception")
	case pgerrcode.IsInvalidCursorName(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "Invalid cursor name")
	case pgerrcode.IsExternalRoutineException(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "External routine exception")
	case pgerrcode.IsExternalRoutineInvocationException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "External routine invocation exception")
	case pgerrcode.IsSavepointException(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Savepoint exception")
	case pgerrcode.IsInvalidCatalogName(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid catalog name")
	case pgerrcode.IsInvalidSchemaName(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Invalid schema name")
	case pgerrcode.IsTransactionRollback(sqlState):
		return fiber.StatusConflict, perror.NewCode(ecode.CodeDbOperationError, "Transaction rollback")
	case pgerrcode.IsSyntaxErrororAccessRuleViolation(sqlState):
		return fiber.StatusBadRequest, perror.NewCode(ecode.CodeDbOperationError, "Syntax error or access rule violation")
	case pgerrcode.IsWithCheckOptionViolation(sqlState):
		return fiber.StatusConflict, perror.NewCode(ecode.CodeDbOperationError, "With check option violation")
	case pgerrcode.IsInsufficientResources(sqlState):
		return fiber.StatusServiceUnavailable, perror.NewCode(ecode.CodeDbOperationError, "Insufficient resources")
	case pgerrcode.IsProgramLimitExceeded(sqlState):
		return fiber.StatusServiceUnavailable, perror.NewCode(ecode.CodeDbOperationError, "Program limit exceeded")
	case pgerrcode.IsObjectNotInPrerequisiteState(sqlState):
		return fiber.StatusUnprocessableEntity, perror.NewCode(ecode.CodeDbOperationError, "Object not in prerequisite state")
	case pgerrcode.IsOperatorIntervention(sqlState):
		return fiber.StatusServiceUnavailable, perror.NewCode(ecode.CodeDbOperationError, "Operator intervention")
	case pgerrcode.IsSystemError(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "System error")
	case pgerrcode.IsSnapshotFailure(sqlState):
		return fiber.StatusConflict, perror.NewCode(ecode.CodeDbOperationError, "Snapshot failure")
	case pgerrcode.IsConfigurationFileError(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "Configuration file error")
	case pgerrcode.IsForeignDataWrapperError(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "Foreign data wrapper error")
	case pgerrcode.IsPLpgSQLError(sqlState):
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "PL/pgSQL error")
	default:
		return fiber.StatusInternalServerError, perror.NewCode(ecode.CodeDbOperationError, "Unknown database error")

	}

}
