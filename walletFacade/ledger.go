package walletFacade

import "fmt"

// Complex subsystem parts
// Make ledge entry

// ledger 分類帳
type Ledger struct{}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
