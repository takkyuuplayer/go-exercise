//go:generate go run ../../../../loilo-inc/logos/cmd/digen/main.go $GOFILE
package logosdi

import "time"

// Define struct named by "Manifest"
type Manifest struct {
	String string
	Time   time.Time
}
