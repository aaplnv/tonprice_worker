// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/aedquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AEDQuote is the model entity for the AEDQuote schema.
type AEDQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AEDQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case aedquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case aedquote.FieldID:
			values[i] = new(sql.NullInt64)
		case aedquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AEDQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AEDQuote fields.
func (aq *AEDQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case aedquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aq.ID = int(value.Int64)
		case aedquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				aq.Price = value.Float64
			}
		case aedquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				aq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AEDQuote.
// Note that you need to call AEDQuote.Unwrap() before calling this method if this AEDQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (aq *AEDQuote) Update() *AEDQuoteUpdateOne {
	return (&AEDQuoteClient{config: aq.config}).UpdateOne(aq)
}

// Unwrap unwraps the AEDQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aq *AEDQuote) Unwrap() *AEDQuote {
	tx, ok := aq.config.driver.(*txDriver)
	if !ok {
		panic("ent: AEDQuote is not a transactional entity")
	}
	aq.config.driver = tx.drv
	return aq
}

// String implements the fmt.Stringer.
func (aq *AEDQuote) String() string {
	var builder strings.Builder
	builder.WriteString("AEDQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", aq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", aq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(aq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AEDQuotes is a parsable slice of AEDQuote.
type AEDQuotes []*AEDQuote

func (aq AEDQuotes) config(cfg config) {
	for _i := range aq {
		aq[_i].config = cfg
	}
}
