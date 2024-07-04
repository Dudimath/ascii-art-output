package ascii

import (
	"reflect"
	"testing"
)

func TestReadASCIIMapFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name:    "Unsupported file name",
			args:    args{filename: "unsupported.txt"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Unsupported file format",
			args:    args{filename: "standard.pdf"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Missing file",
			args:    args{filename: "missing.txt"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Incorrect file size - temporary file",
			args:    args{filename: "incorrect_size.txt"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Incorrect number of characters - temporary file",
			args:    args{filename: "incorrect_chars.txt"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadASCIIMapFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadASCIIMapFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadASCIIMapFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
