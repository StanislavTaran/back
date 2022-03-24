package mysqlClient

import (
	"database/sql"
)

// NullBool Nullable Bool that overrides sql.NullBool
type NullBool struct {
	sql.NullBool
}

func (nb NullBool) Value() bool {
	return nb.Bool
}

// NullFloat64 Nullable Float64 that overrides sql.NullFloat64
type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) Value() float64 {
	return nf.Float64

}

// NullInt64 Nullable Int64 that overrides sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) Value() int64 {
	return ni.Int64
}

// NullString Nullable String that overrides sql.NullString
type NullString struct {
	sql.NullString
}

func (ns NullString) Value() string {
	return ns.String
}
