package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
)

//ParseTemplate function to parse the body of the email with the given parameters
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("../../pkg/email/templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
