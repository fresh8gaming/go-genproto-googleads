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
	resourcespb "github.com/fresh8gaming/go-genproto-googleads/pb/v9/resources"
	servicespb "github.com/fresh8gaming/go-genproto-googleads/pb/v9/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newConversionValueRuleSetClientHook clientHook

// ConversionValueRuleSetCallOptions contains the retry settings for each method of ConversionValueRuleSetClient.
type ConversionValueRuleSetCallOptions struct {
	GetConversionValueRuleSet     []gax.CallOption
	MutateConversionValueRuleSets []gax.CallOption
}

func defaultConversionValueRuleSetGRPCClientOptions() []option.ClientOption {
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

func defaultConversionValueRuleSetCallOptions() *ConversionValueRuleSetCallOptions {
	return &ConversionValueRuleSetCallOptions{
		GetConversionValueRuleSet: []gax.CallOption{
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
		MutateConversionValueRuleSets: []gax.CallOption{
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

// internalConversionValueRuleSetClient is an interface that defines the methods available from Google Ads API.
type internalConversionValueRuleSetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetConversionValueRuleSet(context.Context, *servicespb.GetConversionValueRuleSetRequest, ...gax.CallOption) (*resourcespb.ConversionValueRuleSet, error)
	MutateConversionValueRuleSets(context.Context, *servicespb.MutateConversionValueRuleSetsRequest, ...gax.CallOption) (*servicespb.MutateConversionValueRuleSetsResponse, error)
}

// ConversionValueRuleSetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage conversion value rule sets.
type ConversionValueRuleSetClient struct {
	// The internal transport-dependent client.
	internalClient internalConversionValueRuleSetClient

	// The call options for this service.
	CallOptions *ConversionValueRuleSetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionValueRuleSetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionValueRuleSetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConversionValueRuleSetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetConversionValueRuleSet returns the requested conversion value rule set.
func (c *ConversionValueRuleSetClient) GetConversionValueRuleSet(ctx context.Context, req *servicespb.GetConversionValueRuleSetRequest, opts ...gax.CallOption) (*resourcespb.ConversionValueRuleSet, error) {
	return c.internalClient.GetConversionValueRuleSet(ctx, req, opts...)
}

// MutateConversionValueRuleSets creates, updates or removes conversion value rule sets. Operation statuses
// are returned.
func (c *ConversionValueRuleSetClient) MutateConversionValueRuleSets(ctx context.Context, req *servicespb.MutateConversionValueRuleSetsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionValueRuleSetsResponse, error) {
	return c.internalClient.MutateConversionValueRuleSets(ctx, req, opts...)
}

// conversionValueRuleSetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionValueRuleSetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ConversionValueRuleSetClient
	CallOptions **ConversionValueRuleSetCallOptions

	// The gRPC API client.
	conversionValueRuleSetClient servicespb.ConversionValueRuleSetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConversionValueRuleSetClient creates a new conversion value rule set service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage conversion value rule sets.
func NewConversionValueRuleSetClient(ctx context.Context, opts ...option.ClientOption) (*ConversionValueRuleSetClient, error) {
	clientOpts := defaultConversionValueRuleSetGRPCClientOptions()
	if newConversionValueRuleSetClientHook != nil {
		hookOpts, err := newConversionValueRuleSetClientHook(ctx, clientHookParams{})
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
	client := ConversionValueRuleSetClient{CallOptions: defaultConversionValueRuleSetCallOptions()}

	c := &conversionValueRuleSetGRPCClient{
		connPool:                     connPool,
		disableDeadlines:             disableDeadlines,
		conversionValueRuleSetClient: servicespb.NewConversionValueRuleSetServiceClient(connPool),
		CallOptions:                  &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *conversionValueRuleSetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionValueRuleSetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionValueRuleSetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *conversionValueRuleSetGRPCClient) GetConversionValueRuleSet(ctx context.Context, req *servicespb.GetConversionValueRuleSetRequest, opts ...gax.CallOption) (*resourcespb.ConversionValueRuleSet, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetConversionValueRuleSet[0:len((*c.CallOptions).GetConversionValueRuleSet):len((*c.CallOptions).GetConversionValueRuleSet)], opts...)
	var resp *resourcespb.ConversionValueRuleSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionValueRuleSetClient.GetConversionValueRuleSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversionValueRuleSetGRPCClient) MutateConversionValueRuleSets(ctx context.Context, req *servicespb.MutateConversionValueRuleSetsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionValueRuleSetsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateConversionValueRuleSets[0:len((*c.CallOptions).MutateConversionValueRuleSets):len((*c.CallOptions).MutateConversionValueRuleSets)], opts...)
	var resp *servicespb.MutateConversionValueRuleSetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionValueRuleSetClient.MutateConversionValueRuleSets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
