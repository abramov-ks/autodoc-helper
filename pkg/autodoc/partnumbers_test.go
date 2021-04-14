package autodoc

import "testing"

func TestAutodocSession_getPartnumbersUrl(t *testing.T) {
	type fields struct {
		AuthData       AuthResult
		BaseUrl        string
		AuthUrl        string
		PartnumbersUrl string
		Username       string
		Password       string
	}
	type args struct {
		partnumber string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				PartnumbersUrl: "https://webapi.autodoc.ru/api/spareparts/511/%s/2?isrecross=false",
			},
			args: args{
				partnumber: "123",
			},
			want: "https://webapi.autodoc.ru/api/spareparts/511/123/2?isrecross=false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := &AutodocSession{
				AuthData:       tt.fields.AuthData,
				BaseUrl:        tt.fields.BaseUrl,
				AuthUrl:        tt.fields.AuthUrl,
				PartnumbersUrl: tt.fields.PartnumbersUrl,
				Username:       tt.fields.Username,
				Password:       tt.fields.Password,
			}
			if got := session.getPartnumbersUrl(tt.args.partnumber); got != tt.want {
				t.Errorf("getPartnumbersUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
