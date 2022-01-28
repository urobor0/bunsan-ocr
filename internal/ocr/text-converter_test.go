package ocr

import (
	"bunsan-ocr/kit/identifier"
	"bunsan-ocr/kit/projectpath"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resourcesPath = fmt.Sprintf("%s/resources", projectpath.RootDir())

func TestTextConverter(t *testing.T) {
	id, err := identifier.New()
	assert.NoError(t, err)

	jobID, err := NewJobID(id)
	assert.NoError(t, err)

	type args struct {
		filePath string
		id JobID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid file", args: args{
			filePath: fmt.Sprintf("%s/input-example.txt", resourcesPath),
			id: jobID,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TextConverter(tt.args.filePath, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TextConverter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
