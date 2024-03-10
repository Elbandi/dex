// Code generated by ent, DO NOT EDIT.

package password

import (
	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldEmail, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v []byte) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldHash, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldUsername, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldUserID, v))
}

// Groups applies equality check predicate on the "groups" field. It's identical to GroupsEQ.
func Groups(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldGroups, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Password {
	return predicate.Password(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Password {
	return predicate.Password(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Password {
	return predicate.Password(sql.FieldContainsFold(FieldEmail, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v []byte) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v []byte) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...[]byte) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...[]byte) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v []byte) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v []byte) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v []byte) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v []byte) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldHash, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Password {
	return predicate.Password(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Password {
	return predicate.Password(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Password {
	return predicate.Password(sql.FieldContainsFold(FieldUsername, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Password {
	return predicate.Password(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Password {
	return predicate.Password(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Password {
	return predicate.Password(sql.FieldContainsFold(FieldUserID, v))
}

// GroupsEQ applies the EQ predicate on the "groups" field.
func GroupsEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldEQ(FieldGroups, v))
}

// GroupsNEQ applies the NEQ predicate on the "groups" field.
func GroupsNEQ(v string) predicate.Password {
	return predicate.Password(sql.FieldNEQ(FieldGroups, v))
}

// GroupsIn applies the In predicate on the "groups" field.
func GroupsIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldIn(FieldGroups, vs...))
}

// GroupsNotIn applies the NotIn predicate on the "groups" field.
func GroupsNotIn(vs ...string) predicate.Password {
	return predicate.Password(sql.FieldNotIn(FieldGroups, vs...))
}

// GroupsGT applies the GT predicate on the "groups" field.
func GroupsGT(v string) predicate.Password {
	return predicate.Password(sql.FieldGT(FieldGroups, v))
}

// GroupsGTE applies the GTE predicate on the "groups" field.
func GroupsGTE(v string) predicate.Password {
	return predicate.Password(sql.FieldGTE(FieldGroups, v))
}

// GroupsLT applies the LT predicate on the "groups" field.
func GroupsLT(v string) predicate.Password {
	return predicate.Password(sql.FieldLT(FieldGroups, v))
}

// GroupsLTE applies the LTE predicate on the "groups" field.
func GroupsLTE(v string) predicate.Password {
	return predicate.Password(sql.FieldLTE(FieldGroups, v))
}

// GroupsContains applies the Contains predicate on the "groups" field.
func GroupsContains(v string) predicate.Password {
	return predicate.Password(sql.FieldContains(FieldGroups, v))
}

// GroupsHasPrefix applies the HasPrefix predicate on the "groups" field.
func GroupsHasPrefix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasPrefix(FieldGroups, v))
}

// GroupsHasSuffix applies the HasSuffix predicate on the "groups" field.
func GroupsHasSuffix(v string) predicate.Password {
	return predicate.Password(sql.FieldHasSuffix(FieldGroups, v))
}

// GroupsEqualFold applies the EqualFold predicate on the "groups" field.
func GroupsEqualFold(v string) predicate.Password {
	return predicate.Password(sql.FieldEqualFold(FieldGroups, v))
}

// GroupsContainsFold applies the ContainsFold predicate on the "groups" field.
func GroupsContainsFold(v string) predicate.Password {
	return predicate.Password(sql.FieldContainsFold(FieldGroups, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Password) predicate.Password {
	return predicate.Password(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Password) predicate.Password {
	return predicate.Password(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Password) predicate.Password {
	return predicate.Password(sql.NotPredicates(p))
}
