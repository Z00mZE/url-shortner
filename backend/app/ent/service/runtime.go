// Code generated by entc, DO NOT EDIT.

package service

import (
	"time"

	"github.com/Z00mZE/url-shortner/ent/schema"
	"github.com/Z00mZE/url-shortner/ent/service/shorturl"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	shorturlFields := schema.ShortUrl{}.Fields()
	_ = shorturlFields
	// shorturlDescURL is the schema descriptor for url field.
	shorturlDescURL := shorturlFields[0].Descriptor()
	// shorturl.URLValidator is a validator for the "url" field. It is called by the builders before save.
	shorturl.URLValidator = func() func(string) error {
		validators := shorturlDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shorturlDescExpiredAt is the schema descriptor for expired_at field.
	shorturlDescExpiredAt := shorturlFields[1].Descriptor()
	// shorturl.DefaultExpiredAt holds the default value on creation for the expired_at field.
	shorturl.DefaultExpiredAt = shorturlDescExpiredAt.Default.(func() time.Time)
}
