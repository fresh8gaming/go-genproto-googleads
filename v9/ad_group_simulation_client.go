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

var newAdGroupSimulationClientHook clientHook

// AdGroupSimulationCallOptions contains the retry settings for each method of AdGroupSimulationClient.
type AdGroupSimulationCallOptions struct {
	GetAdGroupSimulation []gax.CallOption
}

func defaultAdGroupSimulationGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupSimulationCallOptions() *AdGroupSimulationCallOptions {
	return &AdGroupSimulationCallOptions{
		GetAdGroupSimulation: []gax.CallOption{
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

// internalAdGroupSimulationClient is an interface that defines the methods availaible from Google Ads API.
type internalAdGroupSimulationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAdGroupSimulation(context.Context, *servicespb.GetAdGroupSimulationRequest, ...gax.CallOption) (*resourcespb.AdGroupSimulation, error)
}

// AdGroupSimulationClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch ad group simulations.
type AdGroupSimulationClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupSimulationClient

	// The call options for this service.
	CallOptions *AdGroupSimulationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupSimulationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupSimulationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupSimulationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAdGroupSimulation returns the requested ad group simulation in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupSimulationClient) GetAdGroupSimulation(ctx context.Context, req *servicespb.GetAdGroupSimulationRequest, opts ...gax.CallOption) (*resourcespb.AdGroupSimulation, error) {
	return c.internalClient.GetAdGroupSimulation(ctx, req, opts...)
}

// adGroupSimulationGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupSimulationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupSimulationClient
	CallOptions **AdGroupSimulationCallOptions

	// The gRPC API client.
	adGroupSimulationClient servicespb.AdGroupSimulationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupSimulationClient creates a new ad group simulation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch ad group simulations.
func NewAdGroupSimulationClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupSimulationClient, error) {
	clientOpts := defaultAdGroupSimulationGRPCClientOptions()
	if newAdGroupSimulationClientHook != nil {
		hookOpts, err := newAdGroupSimulationClientHook(ctx, clientHookParams{})
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
	client := AdGroupSimulationClient{CallOptions: defaultAdGroupSimulationCallOptions()}

	c := &adGroupSimulationGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		adGroupSimulationClient: servicespb.NewAdGroupSimulationServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adGroupSimulationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupSimulationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupSimulationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupSimulationGRPCClient) GetAdGroupSimulation(ctx context.Context, req *servicespb.GetAdGroupSimulationRequest, opts ...gax.CallOption) (*resourcespb.AdGroupSimulation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAdGroupSimulation[0:len((*c.CallOptions).GetAdGroupSimulation):len((*c.CallOptions).GetAdGroupSimulation)], opts...)
	var resp *resourcespb.AdGroupSimulation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupSimulationClient.GetAdGroupSimulation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
