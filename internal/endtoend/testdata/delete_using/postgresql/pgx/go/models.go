// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type A struct {
	AID   pgtype.Text
	BIDFk pgtype.Text
}

type B struct {
	BID pgtype.Text
}
