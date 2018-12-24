package packa

type temporaryError string

func (*temporaryError) Error() string {
	return "packa: temporary error"
}

func (*temporaryError) temporary() bool {
	return true
}
