package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []bye) string {
	return fmt.Sprintf("Bye, Go. You said: %s", string(req))
}
