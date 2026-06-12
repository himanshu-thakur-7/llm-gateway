package gatewayerrors

type ErrorType string

const (
	ErrorTypeRateLimit ErrorType = "rate_limit"
	ErrorTypeAuth      ErrorType = "auth"
	ErrorTypeTimeout   ErrorType = "timeout"
	ErrorTypeProvider  ErrorType = "provider"
	ErrorTypeUnknown   ErrorType = "unknown"
)

type ProviderError struct {
	Type     ErrorType
	Provider string
	Message  string
}

func (e *ProviderError) Error() string {
	return e.Message
}
