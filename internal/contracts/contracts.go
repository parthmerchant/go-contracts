package contracts

import (
	"github.com/jackc/pgtype"
)

type User struct {
	UserId   pgtype.Int8    `json:"user_id"`
	FullName pgtype.Text    `json:"full_name"`
	Balance  pgtype.Numeric `json:"balance"`
}

type Contract struct {
	ContractId pgtype.Int8    `json:"contract_id"`
	IssuerId   pgtype.Int8    `json:"issuer_id"`
	SignerId   pgtype.Int8    `json:"signer_id"`
	Note       pgtype.Text    `json:"note"`
	Amount     pgtype.Numeric `json:"amount"`
	Signed     pgtype.Bool    `json:"signed"`
}
