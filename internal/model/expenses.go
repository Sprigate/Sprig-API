package model

import "time"

// ExpenseType merepresentasikan kolom enum('pengeluaran','pemasukan') di MySQL.
type ExpenseType string
type Wallet string

const (
	ExpenseTypePengeluaran ExpenseType = "pengeluaran"
	ExpenseTypePemasukan   ExpenseType = "pemasukan"
	ExpenseTypeUtang       ExpenseType = "utang"

	// masuk ke grup apa di budget? (Dompet)
	WalletKebutuhan Wallet = "kebutuhan"
	WalletKeinginan Wallet = "keinginan"
	WalletTabungan  Wallet = "tabungan"
)

type Expense struct {
	ID          uint64      `db:"id"`
	UserID      uint64      `db:"user_id"`
	CategoryID  uint64      `db:"category_id"`
	Amount      uint64      `db:"amount"`
	Description *string     `db:"description"` // pointer karena kolom DEFAULT NULL (boleh kosong)
	ExpenseDate time.Time   `db:"expense_date"`
	Type        ExpenseType `db:"type"`
	Wallet      Wallet      `db:"wallet"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}
