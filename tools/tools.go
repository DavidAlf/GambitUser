package tools

import (
	"fmt"
	"time"
)

func FechaMySQL() string {

	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02dt%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func FechaSQLServer() string {

	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
