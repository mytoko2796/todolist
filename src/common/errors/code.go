package errors

import (
	errors "github.com/mytoko2796/sdk-go/stdlib/error"
)

const (
	CodeValue       = 1100
	CodeSQL         = 1200
	CodeES          = 1300
	CodeCache       = 1400
	CodeHTTPClient  = 1500
	Code3rdDep      = 1600
	CodeAuth        = 1700
	CodeHTTPHandler = 1800

	CodeApp = 2000
)

const (
	// Error On Values
	CodeValueInvalid = errors.Code(iota + CodeValue)
)
const (
	// Error On SQL
	CodeSQLBuilder = errors.Code(iota + CodeSQL)
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLCreate
	CodeSQLUpdate
	CodeSQLDelete
	CodeSQLUnlink
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLPrepareStmt
	CodeSQLRecordMustExist
	CodeSQLRecordAlreadyExist
	CodeSQLCannotRetrieveLastInsertID
	CodeSQLCannotRetrieveAffectedRows
	CodeSQLUniqueConstraint
	CodeSQLRecordDoesNotMatch
	CodeSqlRecordNotFound
	CodeSQLRecordIsExpired
	CodeSQLRollback
)
const (
	// Error on ES
	CodeESBuilder = errors.Code(iota + CodeES)
	CodeESRequestSearchAPI
	CodeESRequestIndexAPI
	CodeESUnmarshal
	CodeESMarshal
)
const (
	// Error On Cache
	CodeCacheMarshal = errors.Code(iota + CodeCache)
	CodeCacheUnmarshal
	CodeCacheGetSimpleKey
	CodeCacheSetSimpleKey
	CodeCacheDeleteSimpleKey
	CodeCacheGetHashKey
	CodeCacheSetHashKey
	CodeCacheDeleteHashKey
	CodeCacheSetExpiration
	CodeCacheDecode
	CodeCacheLockNotAcquired
	CodeCacheLockFailed
	CodeCacheInvalidCastType
)
const (
	// Error on HTTP Client
	CodeHTTPClientMarshal = errors.Code(iota + CodeHTTPClient)
	CodeHTTPClientUnmarshal
	CodeHTTPClientErrorOnRequest
	CodeHTTPClientErrorOnReadBody
)
const (
	// Error on 3rd Dep.
	CodeSMSFailure = errors.Code(iota + Code3rdDep)
	CodeMailerFailure
)
const (
	// Code Auth
	CodeAuthRefreshTokenExpired = errors.Code(iota + CodeAuth)
	CodeAuthAccessTokenExpired
	CodeAuthAuthCodeExpired
)

const (
	// Code HTTP Handler
	CodeHTTPBadRequest = errors.Code(iota + CodeHTTPHandler)
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPMarshal
	CodeHTTPServiceUnavailable
)

const (
	CodeOauthRedirectionURIDoesNotMatch = errors.Code(iota + CodeApp)
	CodeOauthBadUsernamePassword
	CodeOauthUsernameHasNotBeenVerified
	CodeOauthInvalidScope
	CodeAppHasBeenVerified
	CodePasswordConfirmationDoesNotMatch
	CodePasswordLengthIsTooFew
)
