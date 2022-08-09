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
	servicespb "github.com/fresh8gaming/go-genproto-googleads/pb/v11/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAudienceInsightsClientHook clientHook

// AudienceInsightsCallOptions contains the retry settings for each method of AudienceInsightsClient.
type AudienceInsightsCallOptions struct {
	GenerateInsightsFinderReport   []gax.CallOption
	ListAudienceInsightsAttributes []gax.CallOption
}

func defaultAudienceInsightsGRPCClientOptions() []option.ClientOption {
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

func defaultAudienceInsightsCallOptions() *AudienceInsightsCallOptions {
	return &AudienceInsightsCallOptions{
		GenerateInsightsFinderReport: []gax.CallOption{
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
		ListAudienceInsightsAttributes: []gax.CallOption{
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

// internalAudienceInsightsClient is an interface that defines the methods available from Google Ads API.
type internalAudienceInsightsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateInsightsFinderReport(context.Context, *servicespb.GenerateInsightsFinderReportRequest, ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error)
	ListAudienceInsightsAttributes(context.Context, *servicespb.ListAudienceInsightsAttributesRequest, ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error)
}

// AudienceInsightsClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads.
type AudienceInsightsClient struct {
	// The internal transport-dependent client.
	internalClient internalAudienceInsightsClient

	// The call options for this service.
	CallOptions *AudienceInsightsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AudienceInsightsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AudienceInsightsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AudienceInsightsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateInsightsFinderReport creates a saved report that can be viewed in the Insights Finder tool.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) GenerateInsightsFinderReport(ctx context.Context, req *servicespb.GenerateInsightsFinderReportRequest, opts ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error) {
	return c.internalClient.GenerateInsightsFinderReport(ctx, req, opts...)
}

// ListAudienceInsightsAttributes searches for audience attributes that can be used to generate insights.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) ListAudienceInsightsAttributes(ctx context.Context, req *servicespb.ListAudienceInsightsAttributesRequest, opts ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error) {
	return c.internalClient.ListAudienceInsightsAttributes(ctx, req, opts...)
}

// audienceInsightsGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type audienceInsightsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AudienceInsightsClient
	CallOptions **AudienceInsightsCallOptions

	// The gRPC API client.
	audienceInsightsClient servicespb.AudienceInsightsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAudienceInsightsClient creates a new audience insights service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads.
func NewAudienceInsightsClient(ctx context.Context, opts ...option.ClientOption) (*AudienceInsightsClient, error) {
	clientOpts := defaultAudienceInsightsGRPCClientOptions()
	if newAudienceInsightsClientHook != nil {
		hookOpts, err := newAudienceInsightsClientHook(ctx, clientHookParams{})
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
	client := AudienceInsightsClient{CallOptions: defaultAudienceInsightsCallOptions()}

	c := &audienceInsightsGRPCClient{
		connPool:               connPool,
		disableDeadlines:       disableDeadlines,
		audienceInsightsClient: servicespb.NewAudienceInsightsServiceClient(connPool),
		CallOptions:            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *audienceInsightsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *audienceInsightsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *audienceInsightsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *audienceInsightsGRPCClient) GenerateInsightsFinderReport(ctx context.Context, req *servicespb.GenerateInsightsFinderReportRequest, opts ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateInsightsFinderReport[0:len((*c.CallOptions).GenerateInsightsFinderReport):len((*c.CallOptions).GenerateInsightsFinderReport)], opts...)
	var resp *servicespb.GenerateInsightsFinderReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.GenerateInsightsFinderReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *audienceInsightsGRPCClient) ListAudienceInsightsAttributes(ctx context.Context, req *servicespb.ListAudienceInsightsAttributesRequest, opts ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAudienceInsightsAttributes[0:len((*c.CallOptions).ListAudienceInsightsAttributes):len((*c.CallOptions).ListAudienceInsightsAttributes)], opts...)
	var resp *servicespb.ListAudienceInsightsAttributesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.ListAudienceInsightsAttributes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
