package paymego

func handleErrorResponse(responseJson PaymeResponse) (PaymeResponse, error) {
	var paymeError error

	errorCode := responseJson.Error.Code

	switch errorCode {
	case P2PIndenticalCardsErrorCode:
		paymeError = ErrP2PIdenticalCards
	case InvalidAmoutErrorCode:
		paymeError = ErrInvalidAmountError
	case InvalidParamsErrorCode:
		paymeError = ErrInvalidParamsError
	case CardNotFoundErrorCode:
		paymeError = ErrCardNotFoundError
	case InvalidFormatTokenErrorCode:
		paymeError = ErrInvalidFormatTokenError
	case CardNumberNotFoundCode:
		paymeError = ErrCardNotFoundError
	case CardExpiredCode:
		paymeError = ErrCardExpiredError
	case ProcessingCenterNotAvailableCode:
		paymeError = ErrProcessingCenterNotAvailableError
	case PaycomServiceNotAvailableCode:
		paymeError = ErrPaycomServiceNotAvailableError
	default:
		if errorCode != 0 {
			paymeError = ErrPaymeError
		}
	}

	return responseJson, paymeError
}

func FromSoumToTiyin(amount int) int {
	return amount * 100
}

func FromTiyinToSoum(amount int) int {
	return amount / 100
}
