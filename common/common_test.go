package common

import "testing"

func TestParseToken(t *testing.T) {

	tests := []struct {
		name    string
		token   string
		want    string
		wantErr bool
	}{
		{name: "Andy",
			token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnb2xlbmRhciIsInN1YiI6IkFuZHkifQ.TcqvtFjryWm2O9SoHhXjftgSW_K5WAdk9Fcxx6EEgVw",
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnb2xlbmRhciIsInN1YiI6IkFuZHkifQ.TcqvtFjryWm2O9SoHhXjftgSW_K5WAdk9Fcxx6EEgVw",
			wantErr: false},
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
