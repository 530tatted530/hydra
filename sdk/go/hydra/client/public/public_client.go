// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DisconnectUser opens ID connect front backchannel enabled logout

This endpoint initiates and completes user logout at ORY Hydra and initiates OpenID Connect Front-/Back-channel logout:

https://openid.net/specs/openid-connect-frontchannel-1_0.html
https://openid.net/specs/openid-connect-backchannel-1_0.html
*/
func (a *Client) DisconnectUser(params *DisconnectUserParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisconnectUserParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disconnectUser",
		Method:             "GET",
		PathPattern:        "/oauth2/disconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisconnectUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DiscoverOpenIDConfiguration opens ID connect discovery

The well known endpoint an be used to retrieve information for OpenID Connect clients. We encourage you to not roll
your own OpenID Connect client but to use an OpenID Connect client library instead. You can learn more on this
flow at https://openid.net/specs/openid-connect-discovery-1_0.html
*/
func (a *Client) DiscoverOpenIDConfiguration(params *DiscoverOpenIDConfigurationParams) (*DiscoverOpenIDConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoverOpenIDConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "discoverOpenIDConfiguration",
		Method:             "GET",
		PathPattern:        "/.well-known/openid-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DiscoverOpenIDConfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DiscoverOpenIDConfigurationOK), nil

}

/*
Oauth2Token thes o auth 2 0 token endpoint

The client makes a request to the token endpoint by sending the
following parameters using the "application/x-www-form-urlencoded" HTTP
request entity-body.

> Do not implement a client for this endpoint yourself. Use a library. There are many libraries
> available for any programming language. You can find a list of libraries here: https://oauth.net/code/
>
> Do not the the Hydra SDK does not implement this endpoint properly. Use one of the libraries listed above!
*/
func (a *Client) Oauth2Token(params *Oauth2TokenParams, authInfo runtime.ClientAuthInfoWriter) (*Oauth2TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOauth2TokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "oauth2Token",
		Method:             "POST",
		PathPattern:        "/oauth2/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Oauth2TokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Oauth2TokenOK), nil

}

/*
OauthAuth thes o auth 2 0 authorize endpoint

This endpoint is not documented here because you should never use your own implementation to perform OAuth2 flows.
OAuth2 is a very popular protocol and a library for your programming language will exists.

To learn more about this flow please refer to the specification: https://tools.ietf.org/html/rfc6749
*/
func (a *Client) OauthAuth(params *OauthAuthParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOauthAuthParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "oauthAuth",
		Method:             "GET",
		PathPattern:        "/oauth2/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OauthAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RevokeOAuth2Token revokes o auth2 tokens

Revoking a token (both access and refresh) means that the tokens will be invalid. A revoked access token can no
longer be used to make access requests, and a revoked refresh token can no longer be used to refresh an access token.
Revoking a refresh token also invalidates the access token that was created with it. A token may only be revoked by
the client the token was generated for.
*/
func (a *Client) RevokeOAuth2Token(params *RevokeOAuth2TokenParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeOAuth2TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeOAuth2TokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeOAuth2Token",
		Method:             "POST",
		PathPattern:        "/oauth2/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RevokeOAuth2TokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RevokeOAuth2TokenOK), nil

}

/*
Userinfo opens ID connect userinfo

This endpoint returns the payload of the ID Token, including the idTokenExtra values, of the provided OAuth 2.0 access token.
The endpoint implements http://openid.net/specs/openid-connect-core-1_0.html#UserInfo .
*/
func (a *Client) Userinfo(params *UserinfoParams, authInfo runtime.ClientAuthInfoWriter) (*UserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userinfo",
		Method:             "GET",
		PathPattern:        "/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserinfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserinfoOK), nil

}

/*
WellKnown JSONs web keys discovery

This endpoint returns JSON Web Keys to be used as public keys for verifying OpenID Connect ID Tokens and,
if enabled, OAuth 2.0 JWT Access Tokens. This endpoint can be used with client libraries like
[node-jwks-rsa](https://github.com/auth0/node-jwks-rsa) among others.
*/
func (a *Client) WellKnown(params *WellKnownParams) (*WellKnownOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWellKnownParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "wellKnown",
		Method:             "GET",
		PathPattern:        "/.well-known/jwks.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WellKnownReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WellKnownOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}