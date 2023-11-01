package utils

import (
	"github.com/pkg/errors"
	"log"
)

func SafeGO(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%+v", errors.Errorf("%+v", r))
			}
		}()
		f()
	}()
}
