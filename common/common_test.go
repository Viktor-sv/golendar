package common

import "testing"

func TestParseToken(t *testing.T) {

	tests := []struct {
		name    string
		token   string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "mytest", token: "my_token", want: "dsfdf", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
