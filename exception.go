package paymego

import "errors"

const (
	InvalidAmoutErrorCode       = -31611
	InvalidParamsErrorCode      = -32602
	P2PIndenticalCardsErrorCode = -31630
	CardNotFoundErrorCode       = -31400
	InvalidFormatTokenErrorCode = -32500
	CardNumberNotFoundCode      = -31300
	CardExpiredCode             = -31301

	// fail
	PaycomServiceNotAvailableCode    = -31001
	ProcessingCenterNotAvailableCode = -31002
)

var (
	ErrPaymeTimeoutError                 = errors.New("payme timeout exceeded")
	ErrCardExpiredError                  = errors.New("card expired")
	ErrPaycomServiceNotAvailableError    = errors.New("paycom service not available")
	ErrProcessingCenterNotAvailableError = errors.New("processing center not available")
	ErrEmptyOrInvalidPaycomKey           = errors.New("invalid paycomKey")
	ErrErrEmptyOrInvalidPaycomID         = errors.New("invalid paycomID")
	ErrInvalidFormatTokenError           = errors.New("invalid format token")
	ErrCardNotFoundError                 = errors.New("payme card not found")
	ErrPaymeError                        = errors.New("payme error was occurred")
	ErrInvalidAmountError                = errors.New("invalid amount")
	ErrInvalidParamsError                = errors.New("invalid params")
	ErrP2PIdenticalCards                 = errors.New("similar cards cannot be used for P2P processing")
)
