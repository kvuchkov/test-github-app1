package function

// Handle a serverless request
func Handle(req []byte) string {
	return fm.Sprintf("Bye, Go. You mentioned: %s", string(req))
}
