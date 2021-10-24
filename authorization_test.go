package clickup

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonalTokenAuthorization_ApplyAuthorization(t *testing.T) {
	type fields struct {
		token string
	}

	type args struct {
		r *Request
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHeader http.Header
	}{
		{
			name: "If I ApplyAuthorization with a PersonalTokenAuthorization, the input request should get a valid Authorization header",
			fields: fields{
				token: "token",
			},
			args: args{
				r: &Request{Request: &http.Request{
					Header: http.Header{},
				}},
			},
			wantHeader: http.Header{
				header_authorization: []string{"token"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewPersonalTokenAuthorization(tt.fields.token)
			a.ApplyAuthorization(tt.args.r)
			assert.Equal(t, tt.wantHeader, tt.args.r.Header)
		})
	}
}
