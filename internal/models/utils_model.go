package models

import (
	"database/sql"
	"encoding/json"
)

type NullableString struct {
	Set bool
	sql.NullString
}

func (n *NullableString) UnmarshalJSON(b []byte) error {
	n.Set = true
	n.Valid = string(b) != "null"
	e := json.Unmarshal(b, &n.String)
	return e
}

func (n NullableString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.String)
}

type NullableBool struct {
	Set bool
	sql.NullBool
}

func (n *NullableBool) UnmarshalJSON(b []byte) error {
	n.Set = true
	n.Valid = string(b) != "null"
	e := json.Unmarshal(b, &n.Bool)
	return e
}

func (n NullableBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Bool)
}
