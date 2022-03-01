package errors

import (
	"net/http"

	errors "github.com/mytoko2796/sdk-go/stdlib/error"
)

func CompileError(err error, lang string, debugMode bool) (int, error) {
	var (
		appError *AppError
		httpCode int
	)
	var debugErr *string
	if debugMode {
		errStr := err.Error()
		if len(errStr) > 0 {
			debugErr = &errStr
		}
	}
	code := errors.ErrCode(err)
	switch code {

	case CodeValueInvalid, CodeHTTPBadRequest, CodeSQLRecordIsExpired, CodeHTTPClientErrorOnReadBody:
		httpCode = http.StatusBadRequest
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "badrequest"),
			sys:          err,
			DebugError:   debugErr,
		}

	case CodeHTTPNotFound, CodeSqlRecordNotFound:
		httpCode = http.StatusNotFound
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "notfound"),
			sys:          err,
			DebugError:   debugErr,
		}

	case CodeSQLUniqueConstraint:
		httpCode = http.StatusBadRequest
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "uniqueconst"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeSQLRecordAlreadyExist:
		httpCode = http.StatusBadRequest
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "alreadyexist"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeSQLRecordDoesNotMatch:
		httpCode = http.StatusBadRequest
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "doesnotmatch"),
			sys:          err,
			DebugError:   debugErr,
		}

	case CodeAuthRefreshTokenExpired:
		httpCode = http.StatusUnauthorized
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "refreshtokenexpired"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeAuthAccessTokenExpired:
		httpCode = http.StatusUnauthorized
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "accesstokenexpired"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeHTTPUnauthorized, CodeOauthBadUsernamePassword, CodeOauthUsernameHasNotBeenVerified:
		httpCode = http.StatusUnauthorized
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "unauthorized"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeHTTPServiceUnavailable:
		httpCode = http.StatusServiceUnavailable
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "serviceunavailable"),
			sys:          err,
			DebugError:   debugErr,
		}
	default:
		httpCode = http.StatusInternalServerError
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "internal"),
			sys:          err,
			DebugError:   debugErr,
		}
	}

	return httpCode, appError
}
