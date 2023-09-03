package paymego

func handleErrorResponse(responseJson PaymeResponse) (PaymeResponse, error) {
	var paymeError error

	errorCode := responseJson.Error.Code

	switch errorCode {
	case P2PIndenticalCardsErrorCode:
		paymeError = &P2PIndenticalCardsError{}
	case InvalidAmoutErrorCode:
		paymeError = &InvalidAmountError{}
	case InvalidParamsErrorCode:
		paymeError = &InvalidParamsError{}
	case CardNotFoundErrorCode:
		paymeError = &CardNotFoundError{}
	case InvalidFormatTokenErrorCode:
		paymeError = &InvalidFormatTokenError{}
	case CardNumberNotFoundCode:
		paymeError = &CardNotFoundError{}
	case CardExpiredCode:
		paymeError = &CardExpiredError{}
	case ProcessingCenterNotAvailableCode:
		paymeError = &ProcessingCenterNotAvailableError{}
	case PaycomServiceNotAvailableCode:
		paymeError = &PaycomServiceNotAvailableError{}
	default:
		if errorCode != 0 {
			paymeError = &PaymeError{}
		}
	}

	return responseJson, paymeError
}
