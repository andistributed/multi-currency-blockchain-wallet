package models

import (
	"upex-wallet/wallet-base/db"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

// TxIn represents tx input.
type TxIn struct {
	gorm.Model
	TxSequenceID string          `gorm:"size:32;index"`
	Address      string          `gorm:"size:100;default:''"`
	Code         int             `gorm:"type:int;default:0"`
	Amount       decimal.Decimal `gorm:"type:decimal(32,20);default:0"`
}

func (*TxIn) TableName() string { return "tx_in" }

// FirstOrCreate find first matched record or create a new one.
func (in *TxIn) FirstOrCreate() error {
	return db.Default().FirstOrCreate(in, "tx_sequence_id = ? and address = ? and code = ?",
		in.TxSequenceID, in.Address, in.Code).Error
}

// GetTxInsBySequenceID gets TxIns by tx sequence id.
func GetTxInsBySequenceID(txSequenceID string) ([]*TxIn, error) {
	var txIns []*TxIn
	err := db.Default().Where("tx_sequence_id = ?", txSequenceID).Find(&txIns).Error
	return txIns, err
}