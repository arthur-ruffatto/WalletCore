package gateway

import (
	"github.com/arthur-ruffatto/wallet-ms/internal/entity"
)

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
