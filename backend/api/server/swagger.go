package server

import (
	"bytes"
	"github.com/jonatascabral/pokedex/backend/pkg/utils"
	"io"
	"os"
)

var (
	swagger string
)

type SwaggerDoc struct{}

func NewSwaggerDoc() *SwaggerDoc {
	return &SwaggerDoc{}
}

func (s *SwaggerDoc) ReadDoc() string {
	if swagger != "" {
		return swagger
	}

	buf := bytes.NewBuffer(nil)

	f, err := os.Open("swagger.json")
	utils.PanicOnError(err)

	io.Copy(buf, f)
	f.Close()

	swagger = string(buf.Bytes())

	return swagger
}
