package packa

type temporaryError string

func (err *temporaryError) Error() string {
	return "packa: temporary error"
}
