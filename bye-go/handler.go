package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprint("Bye, Go. You mentioned: %s", string(req))
}
