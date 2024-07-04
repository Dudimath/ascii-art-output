package ascii

import "testing"

func TestPrintArt(t *testing.T) {
	type args struct {
		str     string
		asciMap [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Empty String", args: args{str: "", asciMap: [][]string{}}, want: "", wantErr: false},
		{name: "Unsupported Escape", args: args{str: "Unsupported\\a", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "Single Character", args: args{str: "A", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "String with Newline", args: args{str: "A\nB", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "String with Spaces", args: args{str: "A B", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "String with Unsupported Characters", args: args{str: "A#B", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "Multiple Lines", args: args{str: "A\nB\nC", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "Unsupported Character in Middle", args: args{str: "AB@C", asciMap: [][]string{}}, want: "", wantErr: true},
		{name: "All Supported Characters", args: args{str: "ABC", asciMap: [][]string{}}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintArt(tt.args.str, tt.args.asciMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrintArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
