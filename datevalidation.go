package datevalidation

import (
	"fmt"
	"time"
)

var timeNow = time.Now

func ValidateDateAllowdFuture(f, d string) error {
	return validateDate(f, d, false)
}

func ValidateDateDeniedFuture(f, d string) error {
	return validateDate(f, d, true)
}

func validateDate(f, d string, allowFuture bool) error {
	t, err := time.Parse(f, d)
	if err != nil {
		return fmt.Errorf(`can not parse:"%s"`, d)
	}
	if d != t.Format(f) {
		return fmt.Errorf(`not exists target date:"%s"`, d)
	}
	if !allowFuture {
		now := timeNow()
		if t.After(now) {
			return fmt.Errorf(`future date:"%s"`, d)
		}
	}

	return nil
}
