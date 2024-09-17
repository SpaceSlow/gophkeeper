package sensitive_records

import (
	"bytes"
	"fmt"
	"net/http"
)

type TextFileStrategy struct{}

func (s *TextFileStrategy) Upload(req *http.Request) (string, error) {
	file, _, err := req.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Text file content: %s", buf.String()), nil
}
