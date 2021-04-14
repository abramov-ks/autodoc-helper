package autodoc

import (
	"reflect"
	"testing"
)

func Test_parseAuthResponse(t *testing.T) {
	type args struct {
		response []byte
	}
	tests := []struct {
		name    string
		args    args
		want    AuthResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test auth response",
			args: args{response: []byte(`{
    "scope": "offline_access",
    "token_type": "Bearer",
    "access_token": "123",
    "expires_in": 1200,
    "refresh_token": "321"
}`)},
			want:    AuthResult{Scope: "offline_access", TokenType: "Bearer", AccessToken: "123", ExpiresIn: 1200, RefreshToken: "321"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAuthResponse(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAuthResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAuthResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
