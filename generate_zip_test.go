package filetools

import (
	"context"
	"errors"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateZip(t *testing.T) {
	output, err := GenerateFile(context.Background(), FileRequest{
		OutputPath:   "/ziptest/123/test.txt",
		TemplateName: "test",
		Data: struct {
			Name string
		}{Name: "nuzur"},
		Funcs:           nil,
		DisableGoFormat: true,
	})

	assert.NotNil(t, output)
	assert.NoError(t, err)
	rootDir := CurrentPath()
	finalOutput := path.Join(rootDir, "ziptest", "123", "test.txt")
	if _, err := os.Stat(finalOutput); errors.Is(err, os.ErrNotExist) {
		assert.NoError(t, err)
	}

	err = GenerateZip(context.Background(), ZipRequest{
		OutputPath: "ziptest",
		Identifier: "123",
	})
	assert.NoError(t, err)
	zipFile := path.Join(rootDir, "ziptest", "123.zip")
	if _, err := os.Stat(zipFile); errors.Is(err, os.ErrNotExist) {
		assert.NoError(t, err)
	}

	os.RemoveAll(path.Join(rootDir, "ziptest"))
}