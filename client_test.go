package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientOptions_Validate(t *testing.T) {
	personalTokenAuth := &PersonalTokenAuthorization{token: "token"}

	type fields struct {
		Authorization AuthorizationMethod
		APITargeting  *ClientAPITargetingOptions
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all required fields are filled with valid values, no errors should be returned",
			fields: fields{
				Authorization: personalTokenAuth,
				APITargeting:  _defaultOptions.APITargeting,
			},
			wantErr: false,
		},
		{
			name: "If Authorization is empty, an error should be returned",
			fields: fields{
				APITargeting: _defaultOptions.APITargeting,
			},
			wantErr: true,
		},
		{
			name: "If APITargeting is empty, an error should be returned",
			fields: fields{
				Authorization: personalTokenAuth,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := ClientOptions{
				Authorization: tt.fields.Authorization,
				APITargeting:  tt.fields.APITargeting,
			}

			err := o.Validate()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClientAPITargetingOptions_Validate(t *testing.T) {
	type fields struct {
		Scheme string
		Host   string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all required fields are filled with valid values, no errors should be returned",
			fields: fields{
				Scheme: _defaultOptions.APITargeting.Scheme,
				Host:   _defaultOptions.APITargeting.Host,
			},
			wantErr: false,
		},
		{
			name: "If Scheme is empty, an error should be returned",
			fields: fields{
				Host: _defaultOptions.APITargeting.Host,
			},
			wantErr: true,
		},
		{
			name: "If Host is empty, an error should be returned",
			fields: fields{
				Scheme: _defaultOptions.APITargeting.Scheme,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := ClientAPITargetingOptions{
				Scheme: tt.fields.Scheme,
				Host:   tt.fields.Host,
			}

			err := o.Validate()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestSetAuthorizationMethod(t *testing.T) {
	auth := NewPersonalTokenAuthorization("token")

	type args struct {
		authorization AuthorizationMethod
	}

	tests := []struct {
		name              string
		args              args
		wantErr           bool
		wantAuthorization AuthorizationMethod
	}{
		{
			name: "If I set a valid AuthorizationMethod, no errors should be returned",
			args: args{
				authorization: auth,
			},
			wantErr:           false,
			wantAuthorization: auth,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var options ClientOptions
			err := SetAuthorizationMethod(tt.args.authorization).apply(&options)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantAuthorization, options.Authorization)
		})
	}
}

func TestSetAPIURL(t *testing.T) {
	type args struct {
		apiURL string
	}

	tests := []struct {
		name                    string
		args                    args
		wantErr                 bool
		wantAPITargetingOptions *ClientAPITargetingOptions
	}{
		{
			name: "If I set a valid API URL, no errors should be returned",
			args: args{
				apiURL: "https://api.clickup.com",
			},
			wantErr: false,
			wantAPITargetingOptions: &ClientAPITargetingOptions{
				Scheme: "https",
				Host:   "api.clickup.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var options ClientOptions
			err := SetAPIURL(tt.args.apiURL).apply(&options)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantAPITargetingOptions, options.APITargeting)
		})
	}
}
