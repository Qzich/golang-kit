package errors

import (
	"fmt"
	"os"
	"strings"
)

//
// ReportStartupErrorAndExit reports the startup error and exits the application.
//
func ReportStartupErrorAndExit(err error) {
	separationString := strings.Repeat("*", 80)
	message := fmt.Sprintf(`
%s
unable to start the application due to:
%s
%s
%+v
`, separationString, err.Error(), separationString, err)

	os.Stderr.Write([]byte(message))
	os.Exit(1)
}
