package errors

type producerError string

func (p producerError) Error() string {
	return string(p)
}

const (
	ErrCartolaProviderNil = producerError("cartola provider está nulo")
	ErrMercadoStoreNil    = producerError("mercado store está nulo")
	ErrStatusStoreNil     = producerError("status store está nulo")
)
