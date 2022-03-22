package source

import "github.com/luizvnasc/cartola/cartola-commons/model"

type CartolaProvider interface {
	GetStatus() (*model.Status, error)
	GetMercado() (*model.Mercado, error)
}

type MercadoStorer interface {
	GetUltimo() (*model.Mercado, error)
	Salva(*model.Mercado) error
	GetByRodada(rodada int) (*model.Mercado, error)
	Transmite(mercado *model.Mercado) error
}

type StatusStorer interface {
	GetUltimo() (*model.Status, error)
	Salva(*model.Status) error
}
