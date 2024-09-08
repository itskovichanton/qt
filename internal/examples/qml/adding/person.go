package main

import (
	"github.com/itskovichanton/qt/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
