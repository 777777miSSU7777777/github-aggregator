package token

type TokenNotFoundError struct {
}

func (e *TokenNotFoundError) Error() string {
	return "Token not found in both directories $HOME and current"
}
