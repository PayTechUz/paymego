package paymego

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

type PaymeTimeoutError struct{}
type PaymeError struct{}
type InvalidAmountError struct{}
type InvalidParamsError struct{}
type P2PIndenticalCardsError struct{}
type PaycomServiceNotAvailableError struct{}
type ProcessingCenterNotAvailableError struct{}
type CardExpiredError struct{}
type CardNotFoundError struct{}
type InvalidFormatTokenError struct{}
type ErrEmptyOrInvalidPaycomID struct{}
type ErrEmptyOrInvalidPaycomKey struct{}

// if the parameter is invalid
func (e *InvalidParamsError) Error() string {
	return "Invalid params"
}

// if amount is negative or more than max/min amount
func (e *InvalidAmountError) Error() string {
	return "Invalid amount"
}

// if the cards are similar, cannot be p2p process
func (e *P2PIndenticalCardsError) Error() string {
	return "similar cards cannot be p2p process"
}

// payme error unknown
func (e *PaymeError) Error() string {
	return "payme error was occurred"
}

// payme card not found
func (e *CardNotFoundError) Error() string {
	return "payme card not found"
}

// invalid format token
func (e *InvalidFormatTokenError) Error() string {
	return "invalid format token"
}

// invalid paycomID
func (e ErrEmptyOrInvalidPaycomID) Error() string {
	return "invalid paycomID"
}

// invalid paycomKey
func (e ErrEmptyOrInvalidPaycomKey) Error() string {
	return "invalid paycomKey"
}

func (e ProcessingCenterNotAvailableError) Error() string {
	return "processing center not available"
}

func (e PaycomServiceNotAvailableError) Error() string {
	return "paycom service not available"
}

// card expired error
func (e CardExpiredError) Error() string {
	return "card expired"
}

// payme timeout exceeded
func (e PaymeTimeoutError) Error() string {
	return "payme timeout exceeded"
}
