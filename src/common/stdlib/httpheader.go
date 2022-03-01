package httpheader

import "go.opencensus.io/plugin/ochttp/propagation/b3"

const (
	// https://datatracker.ietf.org/doc/html/rfc7540#section-8.1.2
	// Just as in HTTP/1.x, header field names are strings of ASCII
	// characters that are compared in a case-insensitive fashion. However,
	// header field names MUST be converted to lowercase prior to their
	// encoding in HTTP/2.  A request or response containing uppercase
	// header field names MUST be treated as malformed (Section 8.1.2.6).

	// Authorization
	Authorization string = `authorization`

	// B3 Propagation Header
	B3TraceID string = b3.TraceIDHeader
	B3SpanID  string = b3.SpanIDHeader
	B3Sampled string = b3.SampledHeader

	// HTTP Header Standard
	RequestID      string = `x-request-id`
	RequestMethod  string = `x-request-method`
	RequestScheme  string = `x-request-scheme`
	KeyServerRoute string = `x-key-server-route`
	ForwardedFor   string = `x-forwarded-for`
	APIKey         string = `x-api-key`

	// Custom HTTP Header
	SessionID               string = `x-session-id`
	UserID                  string = `x-user-id`
	MerchantID              string = `x-merchant-id`
	AgentID                 string = `x-agent-id`
	DriverID                string = `x-driver-id`
	WarehouseOperatorID     string = `x-warehouseop-id`
	OrderID                 string = `x-order-id`
	UserBasicInfo           string = `x-user-basic-info`
	ClientID                string = `x-client-id`
	AppClientID             string = `x-app-client-id`
	AppLang                 string = `x-app-lang`
	AppName                 string = `x-app-name`
	AppVersion              string = `x-app-version`
	AppDebug                string = `x-app-debug`
	RegisterIfNotExist      string = `x-register-if-not-exist`
	AppMigration            string = `x-app-migration`
	AuthTokenTTL            string = `x-authtoken-ttl`
	EmailSent               string = `x-email-sent`
	HubspotSignature        string = `x-hubspot-signature`
	HubspotSignatureVersion string = `x-hubspot-signature-version`
	SkipWelcomekitEmail     string = `x-skip-welcomekit-email`
	ShopifyHMACSHA256       string = `x-shopify-hmac-sha256`
	ShopifyShopDomain       string = `x-shopify-shop-domain`

	// Lang Header
	LangEN string = `en`
	LangID string = `id`

	// UserAgent Header
	UserAgent                  string = `user-agent`
	UserAgentHTTPClientDefault string = `sdkdefault/1.0`
	ContentAccept              string = `accept`
	ContentType                string = `content-type`
	ContentJSON                string = `application/json`
	ContentXML                 string = `application/xml`
	ContentFormURLEncoded      string = `application/x-www-form-urlencoded`
	ContentMultipartFormData   string = `multipart/form-data`

	// Cache Control Header
	CacheControl          string = `cache-control`
	CacheNoCache          string = `no-cache`
	CacheNoStore          string = `no-store`
	CacheMustRevalidate   string = `must-revalidate`
	CacheMustDBRevalidate string = `must-db-revalidate`

	// bpm
	BPMProcessID  string = `x-bpm-process-id`
	BPMWorkflowID string = `x-bpm-workflow-id`
	BPMInstanceID string = `x-bpm-instance-id`
	BPMJobID      string = `x-bpm-job-id`
	BPMJobType    string = `x-bpm-job-type`

	// Test-Related
	TestID           string = `x-test-id`
	TestMockPositive string = `mock-positive`
	TestMockNegative string = `mock-negative`
)
