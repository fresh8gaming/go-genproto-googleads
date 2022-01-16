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
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v9/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v9/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCustomerNegativeCriterionClientHook clientHook

// CustomerNegativeCriterionCallOptions contains the retry settings for each method of CustomerNegativeCriterionClient.
type CustomerNegativeCriterionCallOptions struct {
	GetCustomerNegativeCriterion   []gax.CallOption
	MutateCustomerNegativeCriteria []gax.CallOption
}

func defaultCustomerNegativeCriterionGRPCClientOptions() []option.ClientOption {
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

func defaultCustomerNegativeCriterionCallOptions() *CustomerNegativeCriterionCallOptions {
	return &CustomerNegativeCriterionCallOptions{
		GetCustomerNegativeCriterion: []gax.CallOption{
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
		MutateCustomerNegativeCriteria: []gax.CallOption{
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

// internalCustomerNegativeCriterionClient is an interface that defines the methods availaible from Google Ads API.
type internalCustomerNegativeCriterionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCustomerNegativeCriterion(context.Context, *servicespb.GetCustomerNegativeCriterionRequest, ...gax.CallOption) (*resourcespb.CustomerNegativeCriterion, error)
	MutateCustomerNegativeCriteria(context.Context, *servicespb.MutateCustomerNegativeCriteriaRequest, ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error)
}

// CustomerNegativeCriterionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customer negative criteria.
type CustomerNegativeCriterionClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerNegativeCriterionClient

	// The call options for this service.
	CallOptions *CustomerNegativeCriterionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerNegativeCriterionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerNegativeCriterionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CustomerNegativeCriterionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCustomerNegativeCriterion returns the requested criterion in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerNegativeCriterionClient) GetCustomerNegativeCriterion(ctx context.Context, req *servicespb.GetCustomerNegativeCriterionRequest, opts ...gax.CallOption) (*resourcespb.CustomerNegativeCriterion, error) {
	return c.internalClient.GetCustomerNegativeCriterion(ctx, req, opts...)
}

// MutateCustomerNegativeCriteria creates or removes criteria. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CriterionError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerNegativeCriterionClient) MutateCustomerNegativeCriteria(ctx context.Context, req *servicespb.MutateCustomerNegativeCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error) {
	return c.internalClient.MutateCustomerNegativeCriteria(ctx, req, opts...)
}

// customerNegativeCriterionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerNegativeCriterionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CustomerNegativeCriterionClient
	CallOptions **CustomerNegativeCriterionCallOptions

	// The gRPC API client.
	customerNegativeCriterionClient servicespb.CustomerNegativeCriterionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomerNegativeCriterionClient creates a new customer negative criterion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customer negative criteria.
func NewCustomerNegativeCriterionClient(ctx context.Context, opts ...option.ClientOption) (*CustomerNegativeCriterionClient, error) {
	clientOpts := defaultCustomerNegativeCriterionGRPCClientOptions()
	if newCustomerNegativeCriterionClientHook != nil {
		hookOpts, err := newCustomerNegativeCriterionClientHook(ctx, clientHookParams{})
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
	client := CustomerNegativeCriterionClient{CallOptions: defaultCustomerNegativeCriterionCallOptions()}

	c := &customerNegativeCriterionGRPCClient{
		connPool:                        connPool,
		disableDeadlines:                disableDeadlines,
		customerNegativeCriterionClient: servicespb.NewCustomerNegativeCriterionServiceClient(connPool),
		CallOptions:                     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *customerNegativeCriterionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerNegativeCriterionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerNegativeCriterionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerNegativeCriterionGRPCClient) GetCustomerNegativeCriterion(ctx context.Context, req *servicespb.GetCustomerNegativeCriterionRequest, opts ...gax.CallOption) (*resourcespb.CustomerNegativeCriterion, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCustomerNegativeCriterion[0:len((*c.CallOptions).GetCustomerNegativeCriterion):len((*c.CallOptions).GetCustomerNegativeCriterion)], opts...)
	var resp *resourcespb.CustomerNegativeCriterion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerNegativeCriterionClient.GetCustomerNegativeCriterion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customerNegativeCriterionGRPCClient) MutateCustomerNegativeCriteria(ctx context.Context, req *servicespb.MutateCustomerNegativeCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomerNegativeCriteria[0:len((*c.CallOptions).MutateCustomerNegativeCriteria):len((*c.CallOptions).MutateCustomerNegativeCriteria)], opts...)
	var resp *servicespb.MutateCustomerNegativeCriteriaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerNegativeCriterionClient.MutateCustomerNegativeCriteria(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
