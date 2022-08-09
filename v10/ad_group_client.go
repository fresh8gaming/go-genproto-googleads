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
	servicespb "github.com/fresh8gaming/go-genproto-googleads/pb/v10/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAdGroupClientHook clientHook

// AdGroupCallOptions contains the retry settings for each method of AdGroupClient.
type AdGroupCallOptions struct {
	MutateAdGroups []gax.CallOption
}

func defaultAdGroupGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupCallOptions() *AdGroupCallOptions {
	return &AdGroupCallOptions{
		MutateAdGroups: []gax.CallOption{
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

// internalAdGroupClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroups(context.Context, *servicespb.MutateAdGroupsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupsResponse, error)
}

// AdGroupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad groups.
type AdGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupClient

	// The call options for this service.
	CallOptions *AdGroupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroups creates, updates, or removes ad groups. Operation statuses are returned.
//
// List of thrown errors:
// AdGroupError (at )
// AdxError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// BiddingError (at )
// BiddingStrategyError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// ListOperationError (at )
// MultiplierError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SettingError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdGroupClient) MutateAdGroups(ctx context.Context, req *servicespb.MutateAdGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupsResponse, error) {
	return c.internalClient.MutateAdGroups(ctx, req, opts...)
}

// adGroupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupClient
	CallOptions **AdGroupCallOptions

	// The gRPC API client.
	adGroupClient servicespb.AdGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupClient creates a new ad group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad groups.
func NewAdGroupClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupClient, error) {
	clientOpts := defaultAdGroupGRPCClientOptions()
	if newAdGroupClientHook != nil {
		hookOpts, err := newAdGroupClientHook(ctx, clientHookParams{})
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
	client := AdGroupClient{CallOptions: defaultAdGroupCallOptions()}

	c := &adGroupGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		adGroupClient:    servicespb.NewAdGroupServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupGRPCClient) MutateAdGroups(ctx context.Context, req *servicespb.MutateAdGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdGroups[0:len((*c.CallOptions).MutateAdGroups):len((*c.CallOptions).MutateAdGroups)], opts...)
	var resp *servicespb.MutateAdGroupsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupClient.MutateAdGroups(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
