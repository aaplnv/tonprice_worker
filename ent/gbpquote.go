// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/gbpquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// GBPQuote is the model entity for the GBPQuote schema.
type GBPQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GBPQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gbpquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case gbpquote.FieldID:
			values[i] = new(sql.NullInt64)
		case gbpquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GBPQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GBPQuote fields.
func (gq *GBPQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gbpquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gq.ID = int(value.Int64)
		case gbpquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				gq.Price = value.Float64
			}
		case gbpquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				gq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GBPQuote.
// Note that you need to call GBPQuote.Unwrap() before calling this method if this GBPQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (gq *GBPQuote) Update() *GBPQuoteUpdateOne {
	return (&GBPQuoteClient{config: gq.config}).UpdateOne(gq)
}

// Unwrap unwraps the GBPQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gq *GBPQuote) Unwrap() *GBPQuote {
	tx, ok := gq.config.driver.(*txDriver)
	if !ok {
		panic("ent: GBPQuote is not a transactional entity")
	}
	gq.config.driver = tx.drv
	return gq
}

// String implements the fmt.Stringer.
func (gq *GBPQuote) String() string {
	var builder strings.Builder
	builder.WriteString("GBPQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", gq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", gq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(gq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GBPQuotes is a parsable slice of GBPQuote.
type GBPQuotes []*GBPQuote

func (gq GBPQuotes) config(cfg config) {
	for _i := range gq {
		gq[_i].config = cfg
	}
}
