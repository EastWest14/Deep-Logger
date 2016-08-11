package deeplogger

type CountWriter struct {
	V int
}

func (iw *CountWriter) Write(input []byte) (n int, err error) {
	iw.V++
	return 0, nil
}
