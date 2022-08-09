// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/fresh8/go-genproto-googleads/pb/v11/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCustomerExtensionSettingClientHook clientHook

// CustomerExtensionSettingCallOptions contains the retry settings for each method of CustomerExtensionSettingClient.
type CustomerExtensionSettingCallOptions struct {
	MutateCustomerExtensionSettings []gax.CallOption
}

func defaultCustomerExtensionSettingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCustomerExtensionSettingCallOptions() *CustomerExtensionSettingCallOptions {
	return &CustomerExtensionSettingCallOptions{
		MutateCustomerExtensionSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalCustomerExtensionSettingClient is an interface that defines the methods available from Google Ads API.
type internalCustomerExtensionSettingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomerExtensionSettings(context.Context, *servicespb.MutateCustomerExtensionSettingsRequest, ...gax.CallOption) (*servicespb.MutateCustomerExtensionSettingsResponse, error)
}

// CustomerExtensionSettingClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customer extension settings.
type CustomerExtensionSettingClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerExtensionSettingClient

	// The call options for this service.
	CallOptions *CustomerExtensionSettingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerExtensionSettingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerExtensionSettingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CustomerExtensionSettingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomerExtensionSettings creates, updates, or removes customer extension settings. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// CriterionError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// ExtensionSettingError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// ListOperationError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *CustomerExtensionSettingClient) MutateCustomerExtensionSettings(ctx context.Context, req *servicespb.MutateCustomerExtensionSettingsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerExtensionSettingsResponse, error) {
	return c.internalClient.MutateCustomerExtensionSettings(ctx, req, opts...)
}

// customerExtensionSettingGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerExtensionSettingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CustomerExtensionSettingClient
	CallOptions **CustomerExtensionSettingCallOptions

	// The gRPC API client.
	customerExtensionSettingClient servicespb.CustomerExtensionSettingServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomerExtensionSettingClient creates a new customer extension setting service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customer extension settings.
func NewCustomerExtensionSettingClient(ctx context.Context, opts ...option.ClientOption) (*CustomerExtensionSettingClient, error) {
	clientOpts := defaultCustomerExtensionSettingGRPCClientOptions()
	if newCustomerExtensionSettingClientHook != nil {
		hookOpts, err := newCustomerExtensionSettingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomerExtensionSettingClient{CallOptions: defaultCustomerExtensionSettingCallOptions()}

	c := &customerExtensionSettingGRPCClient{
		connPool:                       connPool,
		disableDeadlines:               disableDeadlines,
		customerExtensionSettingClient: servicespb.NewCustomerExtensionSettingServiceClient(connPool),
		CallOptions:                    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *customerExtensionSettingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerExtensionSettingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerExtensionSettingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerExtensionSettingGRPCClient) MutateCustomerExtensionSettings(ctx context.Context, req *servicespb.MutateCustomerExtensionSettingsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerExtensionSettingsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomerExtensionSettings[0:len((*c.CallOptions).MutateCustomerExtensionSettings):len((*c.CallOptions).MutateCustomerExtensionSettings)], opts...)
	var resp *servicespb.MutateCustomerExtensionSettingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerExtensionSettingClient.MutateCustomerExtensionSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
