package transactions

import (
	"time"

	"github.com/google/uuid"
	"github.com/kahshiuhtang/Insights/internal/blockchain"
)
const (
	Exchange = iota
	Review
	Status
)
type TransactionType = int
// Group Transactions?
type Transaction struct {
	transType TransactionType
	id int64
	timestamp time.Time
	fromUUID uuid.UUID
	toUUID uuid.UUID
	itemUUID uuid.UUID
}
func (t Transaction) AddToChain(b *blockchain.Blockchain) bool{
	b.AddTransactionToChain(t);
	return true;
}
func (t Transaction) FindTransactionInChain(b *blockchain.Blockchain) *blockchain.Block{
	return nil;
}
func (t Transaction) toString() map[string]interface{}{
	blockData := map[string]interface{}{
                "transType":   t.transType,
				"id": 			t.id,
                "from":     	t.fromUUID,
				"to": 			t.toUUID,
                "amount": 		t.itemUUID,
    }
	return blockData;
}