package ocr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b , _, _ = runtime.Caller(0)
	basePath = filepath.Join(filepath.Dir(b), "../../")
	resourcesPath = fmt.Sprintf("%s/resources", basePath)
)

func TestTextConverter(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid file", args: args{
			filePath: fmt.Sprintf("%s/input-example.txt", resourcesPath),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TextConverter(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("TextConverter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
