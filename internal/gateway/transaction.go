package gateway

import "github.com/arthur-ruffatto/wallet-ms/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
