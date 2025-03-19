package user_account

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
)

var accountRepo = repository.NewUserAccountRepository()
var ledgerRepo = repository.NewLedgerRepository()
