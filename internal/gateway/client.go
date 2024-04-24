package gateway

import "github.com/arthur-ruffatto/wallet-ms/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
