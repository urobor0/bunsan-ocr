package ocr

import (
	"bunsan-ocr/kit/projectpath"
	"fmt"
	"testing"
)

var resourcesPath = fmt.Sprintf("%s/resources", projectpath.RootDir())

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
