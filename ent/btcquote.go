// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/btcquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// BTCQuote is the model entity for the BTCQuote schema.
type BTCQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BTCQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case btcquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case btcquote.FieldID:
			values[i] = new(sql.NullInt64)
		case btcquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BTCQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BTCQuote fields.
func (bq *BTCQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case btcquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bq.ID = int(value.Int64)
		case btcquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				bq.Price = value.Float64
			}
		case btcquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				bq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this BTCQuote.
// Note that you need to call BTCQuote.Unwrap() before calling this method if this BTCQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (bq *BTCQuote) Update() *BTCQuoteUpdateOne {
	return (&BTCQuoteClient{config: bq.config}).UpdateOne(bq)
}

// Unwrap unwraps the BTCQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bq *BTCQuote) Unwrap() *BTCQuote {
	tx, ok := bq.config.driver.(*txDriver)
	if !ok {
		panic("ent: BTCQuote is not a transactional entity")
	}
	bq.config.driver = tx.drv
	return bq
}

// String implements the fmt.Stringer.
func (bq *BTCQuote) String() string {
	var builder strings.Builder
	builder.WriteString("BTCQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", bq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", bq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(bq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BTCQuotes is a parsable slice of BTCQuote.
type BTCQuotes []*BTCQuote

func (bq BTCQuotes) config(cfg config) {
	for _i := range bq {
		bq[_i].config = cfg
	}
}
