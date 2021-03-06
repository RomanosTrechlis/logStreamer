package print

import (
	"fmt"

	t "github.com/RomanosTrechlis/go-scribe/internal/util/format/time"
)

const (
	layout string = "2006-01-02 15:04:05.999"
)

// Print logs to console some informative message
func Print(message string) {
	fmt.Printf("%s [INFO] %s\n", t.PrintTime(layout), message)
}
