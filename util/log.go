package util

import (
	"github.com/gookit/color"
	"os"
)

func Fatalln(v ...interface{}) {
	color.Red.Println(v...)
	os.Exit(1)
}
