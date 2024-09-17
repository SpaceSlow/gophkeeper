package sensitive_records

import (
	"bytes"
	"fmt"
	"net/http"
)

type BinaryFileStrategy struct{}

func (s *BinaryFileStrategy) Upload(req *http.Request) (string, error) {
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
	return fmt.Sprintf("Binary file of size %d bytes", buf.Len()), nil
}
