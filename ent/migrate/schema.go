// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AEDColumns holds the columns for the "AED" table.
	AEDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// AEDTable holds the schema information for the "AED" table.
	AEDTable = &schema.Table{
		Name:       "AED",
		Columns:    AEDColumns,
		PrimaryKey: []*schema.Column{AEDColumns[0]},
	}
	// ARSColumns holds the columns for the "ARS" table.
	ARSColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// ARSTable holds the schema information for the "ARS" table.
	ARSTable = &schema.Table{
		Name:       "ARS",
		Columns:    ARSColumns,
		PrimaryKey: []*schema.Column{ARSColumns[0]},
	}
	// AUDColumns holds the columns for the "AUD" table.
	AUDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// AUDTable holds the schema information for the "AUD" table.
	AUDTable = &schema.Table{
		Name:       "AUD",
		Columns:    AUDColumns,
		PrimaryKey: []*schema.Column{AUDColumns[0]},
	}
	// BHDColumns holds the columns for the "BHD" table.
	BHDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BHDTable holds the schema information for the "BHD" table.
	BHDTable = &schema.Table{
		Name:       "BHD",
		Columns:    BHDColumns,
		PrimaryKey: []*schema.Column{BHDColumns[0]},
	}
	// BRLColumns holds the columns for the "BRL" table.
	BRLColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BRLTable holds the schema information for the "BRL" table.
	BRLTable = &schema.Table{
		Name:       "BRL",
		Columns:    BRLColumns,
		PrimaryKey: []*schema.Column{BRLColumns[0]},
	}
	// BTCColumns holds the columns for the "BTC" table.
	BTCColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BTCTable holds the schema information for the "BTC" table.
	BTCTable = &schema.Table{
		Name:       "BTC",
		Columns:    BTCColumns,
		PrimaryKey: []*schema.Column{BTCColumns[0]},
	}
	// CADColumns holds the columns for the "CAD" table.
	CADColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// CADTable holds the schema information for the "CAD" table.
	CADTable = &schema.Table{
		Name:       "CAD",
		Columns:    CADColumns,
		PrimaryKey: []*schema.Column{CADColumns[0]},
	}
	// CHFColumns holds the columns for the "CHF" table.
	CHFColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// CHFTable holds the schema information for the "CHF" table.
	CHFTable = &schema.Table{
		Name:       "CHF",
		Columns:    CHFColumns,
		PrimaryKey: []*schema.Column{CHFColumns[0]},
	}
	// CLPColumns holds the columns for the "CLP" table.
	CLPColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// CLPTable holds the schema information for the "CLP" table.
	CLPTable = &schema.Table{
		Name:       "CLP",
		Columns:    CLPColumns,
		PrimaryKey: []*schema.Column{CLPColumns[0]},
	}
	// CNYColumns holds the columns for the "CNY" table.
	CNYColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// CNYTable holds the schema information for the "CNY" table.
	CNYTable = &schema.Table{
		Name:       "CNY",
		Columns:    CNYColumns,
		PrimaryKey: []*schema.Column{CNYColumns[0]},
	}
	// CZKColumns holds the columns for the "CZK" table.
	CZKColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// CZKTable holds the schema information for the "CZK" table.
	CZKTable = &schema.Table{
		Name:       "CZK",
		Columns:    CZKColumns,
		PrimaryKey: []*schema.Column{CZKColumns[0]},
	}
	// EUROColumns holds the columns for the "EURO" table.
	EUROColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// EUROTable holds the schema information for the "EURO" table.
	EUROTable = &schema.Table{
		Name:       "EURO",
		Columns:    EUROColumns,
		PrimaryKey: []*schema.Column{EUROColumns[0]},
	}
	// GBPColumns holds the columns for the "GBP" table.
	GBPColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// GBPTable holds the schema information for the "GBP" table.
	GBPTable = &schema.Table{
		Name:       "GBP",
		Columns:    GBPColumns,
		PrimaryKey: []*schema.Column{GBPColumns[0]},
	}
	// HKDColumns holds the columns for the "HKD" table.
	HKDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// HKDTable holds the schema information for the "HKD" table.
	HKDTable = &schema.Table{
		Name:       "HKD",
		Columns:    HKDColumns,
		PrimaryKey: []*schema.Column{HKDColumns[0]},
	}
	// HUFColumns holds the columns for the "HUF" table.
	HUFColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// HUFTable holds the schema information for the "HUF" table.
	HUFTable = &schema.Table{
		Name:       "HUF",
		Columns:    HUFColumns,
		PrimaryKey: []*schema.Column{HUFColumns[0]},
	}
	// IDRColumns holds the columns for the "IDR" table.
	IDRColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// IDRTable holds the schema information for the "IDR" table.
	IDRTable = &schema.Table{
		Name:       "IDR",
		Columns:    IDRColumns,
		PrimaryKey: []*schema.Column{IDRColumns[0]},
	}
	// ILSColumns holds the columns for the "ILS" table.
	ILSColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// ILSTable holds the schema information for the "ILS" table.
	ILSTable = &schema.Table{
		Name:       "ILS",
		Columns:    ILSColumns,
		PrimaryKey: []*schema.Column{ILSColumns[0]},
	}
	// INRColumns holds the columns for the "INR" table.
	INRColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// INRTable holds the schema information for the "INR" table.
	INRTable = &schema.Table{
		Name:       "INR",
		Columns:    INRColumns,
		PrimaryKey: []*schema.Column{INRColumns[0]},
	}
	// JPYColumns holds the columns for the "JPY" table.
	JPYColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// JPYTable holds the schema information for the "JPY" table.
	JPYTable = &schema.Table{
		Name:       "JPY",
		Columns:    JPYColumns,
		PrimaryKey: []*schema.Column{JPYColumns[0]},
	}
	// MXNColumns holds the columns for the "MXN" table.
	MXNColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// MXNTable holds the schema information for the "MXN" table.
	MXNTable = &schema.Table{
		Name:       "MXN",
		Columns:    MXNColumns,
		PrimaryKey: []*schema.Column{MXNColumns[0]},
	}
	// NOKColumns holds the columns for the "NOK" table.
	NOKColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// NOKTable holds the schema information for the "NOK" table.
	NOKTable = &schema.Table{
		Name:       "NOK",
		Columns:    NOKColumns,
		PrimaryKey: []*schema.Column{NOKColumns[0]},
	}
	// NZDColumns holds the columns for the "NZD" table.
	NZDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// NZDTable holds the schema information for the "NZD" table.
	NZDTable = &schema.Table{
		Name:       "NZD",
		Columns:    NZDColumns,
		PrimaryKey: []*schema.Column{NZDColumns[0]},
	}
	// PKRColumns holds the columns for the "PKR" table.
	PKRColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// PKRTable holds the schema information for the "PKR" table.
	PKRTable = &schema.Table{
		Name:       "PKR",
		Columns:    PKRColumns,
		PrimaryKey: []*schema.Column{PKRColumns[0]},
	}
	// PLNColumns holds the columns for the "PLN" table.
	PLNColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// PLNTable holds the schema information for the "PLN" table.
	PLNTable = &schema.Table{
		Name:       "PLN",
		Columns:    PLNColumns,
		PrimaryKey: []*schema.Column{PLNColumns[0]},
	}
	// RUBColumns holds the columns for the "RUB" table.
	RUBColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// RUBTable holds the schema information for the "RUB" table.
	RUBTable = &schema.Table{
		Name:       "RUB",
		Columns:    RUBColumns,
		PrimaryKey: []*schema.Column{RUBColumns[0]},
	}
	// SARColumns holds the columns for the "SAR" table.
	SARColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// SARTable holds the schema information for the "SAR" table.
	SARTable = &schema.Table{
		Name:       "SAR",
		Columns:    SARColumns,
		PrimaryKey: []*schema.Column{SARColumns[0]},
	}
	// SEKColumns holds the columns for the "SEK" table.
	SEKColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// SEKTable holds the schema information for the "SEK" table.
	SEKTable = &schema.Table{
		Name:       "SEK",
		Columns:    SEKColumns,
		PrimaryKey: []*schema.Column{SEKColumns[0]},
	}
	// TRYColumns holds the columns for the "TRY" table.
	TRYColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// TRYTable holds the schema information for the "TRY" table.
	TRYTable = &schema.Table{
		Name:       "TRY",
		Columns:    TRYColumns,
		PrimaryKey: []*schema.Column{TRYColumns[0]},
	}
	// TWDColumns holds the columns for the "TWD" table.
	TWDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// TWDTable holds the schema information for the "TWD" table.
	TWDTable = &schema.Table{
		Name:       "TWD",
		Columns:    TWDColumns,
		PrimaryKey: []*schema.Column{TWDColumns[0]},
	}
	// UAHColumns holds the columns for the "UAH" table.
	UAHColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// UAHTable holds the schema information for the "UAH" table.
	UAHTable = &schema.Table{
		Name:       "UAH",
		Columns:    UAHColumns,
		PrimaryKey: []*schema.Column{UAHColumns[0]},
	}
	// USDColumns holds the columns for the "USD" table.
	USDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// USDTable holds the schema information for the "USD" table.
	USDTable = &schema.Table{
		Name:       "USD",
		Columns:    USDColumns,
		PrimaryKey: []*schema.Column{USDColumns[0]},
	}
	// UsersColumns holds the columns for the "Users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "telegram_id", Type: field.TypeInt64, Unique: true},
		{Name: "active_stable", Type: field.TypeString, Nullable: true},
		{Name: "all_stables", Type: field.TypeString, Nullable: true},
		{Name: "reg_time", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "Users" table.
	UsersTable = &schema.Table{
		Name:       "Users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ZARColumns holds the columns for the "ZAR" table.
	ZARColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// ZARTable holds the schema information for the "ZAR" table.
	ZARTable = &schema.Table{
		Name:       "ZAR",
		Columns:    ZARColumns,
		PrimaryKey: []*schema.Column{ZARColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AEDTable,
		ARSTable,
		AUDTable,
		BHDTable,
		BRLTable,
		BTCTable,
		CADTable,
		CHFTable,
		CLPTable,
		CNYTable,
		CZKTable,
		EUROTable,
		GBPTable,
		HKDTable,
		HUFTable,
		IDRTable,
		ILSTable,
		INRTable,
		JPYTable,
		MXNTable,
		NOKTable,
		NZDTable,
		PKRTable,
		PLNTable,
		RUBTable,
		SARTable,
		SEKTable,
		TRYTable,
		TWDTable,
		UAHTable,
		USDTable,
		UsersTable,
		ZARTable,
	}
)

func init() {
	AEDTable.Annotation = &entsql.Annotation{
		Table: "AED",
	}
	ARSTable.Annotation = &entsql.Annotation{
		Table: "ARS",
	}
	AUDTable.Annotation = &entsql.Annotation{
		Table: "AUD",
	}
	BHDTable.Annotation = &entsql.Annotation{
		Table: "BHD",
	}
	BRLTable.Annotation = &entsql.Annotation{
		Table: "BRL",
	}
	BTCTable.Annotation = &entsql.Annotation{
		Table: "BTC",
	}
	CADTable.Annotation = &entsql.Annotation{
		Table: "CAD",
	}
	CHFTable.Annotation = &entsql.Annotation{
		Table: "CHF",
	}
	CLPTable.Annotation = &entsql.Annotation{
		Table: "CLP",
	}
	CNYTable.Annotation = &entsql.Annotation{
		Table: "CNY",
	}
	CZKTable.Annotation = &entsql.Annotation{
		Table: "CZK",
	}
	EUROTable.Annotation = &entsql.Annotation{
		Table: "EURO",
	}
	GBPTable.Annotation = &entsql.Annotation{
		Table: "GBP",
	}
	HKDTable.Annotation = &entsql.Annotation{
		Table: "HKD",
	}
	HUFTable.Annotation = &entsql.Annotation{
		Table: "HUF",
	}
	IDRTable.Annotation = &entsql.Annotation{
		Table: "IDR",
	}
	ILSTable.Annotation = &entsql.Annotation{
		Table: "ILS",
	}
	INRTable.Annotation = &entsql.Annotation{
		Table: "INR",
	}
	JPYTable.Annotation = &entsql.Annotation{
		Table: "JPY",
	}
	MXNTable.Annotation = &entsql.Annotation{
		Table: "MXN",
	}
	NOKTable.Annotation = &entsql.Annotation{
		Table: "NOK",
	}
	NZDTable.Annotation = &entsql.Annotation{
		Table: "NZD",
	}
	PKRTable.Annotation = &entsql.Annotation{
		Table: "PKR",
	}
	PLNTable.Annotation = &entsql.Annotation{
		Table: "PLN",
	}
	RUBTable.Annotation = &entsql.Annotation{
		Table: "RUB",
	}
	SARTable.Annotation = &entsql.Annotation{
		Table: "SAR",
	}
	SEKTable.Annotation = &entsql.Annotation{
		Table: "SEK",
	}
	TRYTable.Annotation = &entsql.Annotation{
		Table: "TRY",
	}
	TWDTable.Annotation = &entsql.Annotation{
		Table: "TWD",
	}
	UAHTable.Annotation = &entsql.Annotation{
		Table: "UAH",
	}
	USDTable.Annotation = &entsql.Annotation{
		Table: "USD",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "Users",
	}
	ZARTable.Annotation = &entsql.Annotation{
		Table: "ZAR",
	}
}
