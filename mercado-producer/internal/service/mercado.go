package service

import (
	"github.com/luizvnasc/cartola/mercado-producer/internal/errors"
	"github.com/luizvnasc/cartola/mercado-producer/internal/source"
)

type MercadoService struct {
	cartolaProvider source.CartolaProvider
	mercadoStore    source.MercadoStorer
	statusStore     source.StatusStorer
}

func NewMercadoService(cartolaProvider source.CartolaProvider, mercadoStore source.MercadoStorer, statusStore source.StatusStorer) (MercadoServicer, error) {
	if cartolaProvider == nil {
		return nil, errors.ErrCartolaProviderNil
	}
	if mercadoStore == nil {
		return nil, errors.ErrMercadoStoreNil
	}
	if statusStore == nil {
		return nil, errors.ErrStatusStoreNil
	}
	return &MercadoService{
		cartolaProvider: cartolaProvider,
		mercadoStore:    mercadoStore,
		statusStore:     statusStore,
	}, nil
}

func (m *MercadoService) Produz() error {
	return nil
}
