//go:build tools

package lib

import (
	_ "golang.org/x/tools/cmd/stringer"
)

// stringer: //go:generate go run golang.org/x/tools/cmd/stringer -type=TypeName1,TypeName2 -output enums_string.go
