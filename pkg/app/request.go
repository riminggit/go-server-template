package app

import (
	"github.com/astaxie/beego/validation"
	"go-server-template/pkg/log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(err.Message)
		// logging.Info(err.Key, err.Message)
	}

	return
}
