// Code generated by entc, DO NOT EDIT.

package user

import (
	"main/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TelegramId applies equality check predicate on the "TelegramId" field. It's identical to TelegramIdEQ.
func TelegramId(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelegramId), v))
	})
}

// AllStables applies equality check predicate on the "AllStables" field. It's identical to AllStablesEQ.
func AllStables(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllStables), v))
	})
}

// RegTime applies equality check predicate on the "RegTime" field. It's identical to RegTimeEQ.
func RegTime(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegTime), v))
	})
}

// TelegramIdEQ applies the EQ predicate on the "TelegramId" field.
func TelegramIdEQ(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelegramId), v))
	})
}

// TelegramIdNEQ applies the NEQ predicate on the "TelegramId" field.
func TelegramIdNEQ(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelegramId), v))
	})
}

// TelegramIdIn applies the In predicate on the "TelegramId" field.
func TelegramIdIn(vs ...int64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelegramId), v...))
	})
}

// TelegramIdNotIn applies the NotIn predicate on the "TelegramId" field.
func TelegramIdNotIn(vs ...int64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelegramId), v...))
	})
}

// TelegramIdGT applies the GT predicate on the "TelegramId" field.
func TelegramIdGT(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelegramId), v))
	})
}

// TelegramIdGTE applies the GTE predicate on the "TelegramId" field.
func TelegramIdGTE(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelegramId), v))
	})
}

// TelegramIdLT applies the LT predicate on the "TelegramId" field.
func TelegramIdLT(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelegramId), v))
	})
}

// TelegramIdLTE applies the LTE predicate on the "TelegramId" field.
func TelegramIdLTE(v int64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelegramId), v))
	})
}

// AllStablesEQ applies the EQ predicate on the "AllStables" field.
func AllStablesEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllStables), v))
	})
}

// AllStablesNEQ applies the NEQ predicate on the "AllStables" field.
func AllStablesNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllStables), v))
	})
}

// AllStablesIn applies the In predicate on the "AllStables" field.
func AllStablesIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAllStables), v...))
	})
}

// AllStablesNotIn applies the NotIn predicate on the "AllStables" field.
func AllStablesNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAllStables), v...))
	})
}

// AllStablesGT applies the GT predicate on the "AllStables" field.
func AllStablesGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllStables), v))
	})
}

// AllStablesGTE applies the GTE predicate on the "AllStables" field.
func AllStablesGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllStables), v))
	})
}

// AllStablesLT applies the LT predicate on the "AllStables" field.
func AllStablesLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllStables), v))
	})
}

// AllStablesLTE applies the LTE predicate on the "AllStables" field.
func AllStablesLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllStables), v))
	})
}

// AllStablesContains applies the Contains predicate on the "AllStables" field.
func AllStablesContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAllStables), v))
	})
}

// AllStablesHasPrefix applies the HasPrefix predicate on the "AllStables" field.
func AllStablesHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAllStables), v))
	})
}

// AllStablesHasSuffix applies the HasSuffix predicate on the "AllStables" field.
func AllStablesHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAllStables), v))
	})
}

// AllStablesIsNil applies the IsNil predicate on the "AllStables" field.
func AllStablesIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAllStables)))
	})
}

// AllStablesNotNil applies the NotNil predicate on the "AllStables" field.
func AllStablesNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAllStables)))
	})
}

// AllStablesEqualFold applies the EqualFold predicate on the "AllStables" field.
func AllStablesEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAllStables), v))
	})
}

// AllStablesContainsFold applies the ContainsFold predicate on the "AllStables" field.
func AllStablesContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAllStables), v))
	})
}

// RegTimeEQ applies the EQ predicate on the "RegTime" field.
func RegTimeEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegTime), v))
	})
}

// RegTimeNEQ applies the NEQ predicate on the "RegTime" field.
func RegTimeNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegTime), v))
	})
}

// RegTimeIn applies the In predicate on the "RegTime" field.
func RegTimeIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegTime), v...))
	})
}

// RegTimeNotIn applies the NotIn predicate on the "RegTime" field.
func RegTimeNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegTime), v...))
	})
}

// RegTimeGT applies the GT predicate on the "RegTime" field.
func RegTimeGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegTime), v))
	})
}

// RegTimeGTE applies the GTE predicate on the "RegTime" field.
func RegTimeGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegTime), v))
	})
}

// RegTimeLT applies the LT predicate on the "RegTime" field.
func RegTimeLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegTime), v))
	})
}

// RegTimeLTE applies the LTE predicate on the "RegTime" field.
func RegTimeLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
