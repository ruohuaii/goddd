// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ruohuaii/goddd/infrastructure/db/ent/merchant"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	merchantFields := schema.Merchant{}.Fields()
	_ = merchantFields
	// merchantDescName is the schema descriptor for name field.
	merchantDescName := merchantFields[1].Descriptor()
	// merchant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	merchant.NameValidator = merchantDescName.Validators[0].(func(string) error)
	// merchantDescEmail is the schema descriptor for email field.
	merchantDescEmail := merchantFields[3].Descriptor()
	// merchant.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	merchant.EmailValidator = merchantDescEmail.Validators[0].(func(string) error)
	// merchantDescMobile is the schema descriptor for mobile field.
	merchantDescMobile := merchantFields[4].Descriptor()
	// merchant.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	merchant.MobileValidator = merchantDescMobile.Validators[0].(func(string) error)
	// merchantDescSalt is the schema descriptor for salt field.
	merchantDescSalt := merchantFields[5].Descriptor()
	// merchant.SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	merchant.SaltValidator = merchantDescSalt.Validators[0].(func(string) error)
	// merchantDescPassword is the schema descriptor for password field.
	merchantDescPassword := merchantFields[6].Descriptor()
	// merchant.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	merchant.PasswordValidator = merchantDescPassword.Validators[0].(func(string) error)
}
