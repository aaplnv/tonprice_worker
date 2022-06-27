// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/tryquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TRYQuote is the model entity for the TRYQuote schema.
type TRYQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TRYQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tryquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case tryquote.FieldID:
			values[i] = new(sql.NullInt64)
		case tryquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TRYQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TRYQuote fields.
func (tq *TRYQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tryquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tq.ID = int(value.Int64)
		case tryquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				tq.Price = value.Float64
			}
		case tryquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				tq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TRYQuote.
// Note that you need to call TRYQuote.Unwrap() before calling this method if this TRYQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (tq *TRYQuote) Update() *TRYQuoteUpdateOne {
	return (&TRYQuoteClient{config: tq.config}).UpdateOne(tq)
}

// Unwrap unwraps the TRYQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tq *TRYQuote) Unwrap() *TRYQuote {
	tx, ok := tq.config.driver.(*txDriver)
	if !ok {
		panic("ent: TRYQuote is not a transactional entity")
	}
	tq.config.driver = tx.drv
	return tq
}

// String implements the fmt.Stringer.
func (tq *TRYQuote) String() string {
	var builder strings.Builder
	builder.WriteString("TRYQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", tq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", tq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(tq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TRYQuotes is a parsable slice of TRYQuote.
type TRYQuotes []*TRYQuote

func (tq TRYQuotes) config(cfg config) {
	for _i := range tq {
		tq[_i].config = cfg
	}
}
