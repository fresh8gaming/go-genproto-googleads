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
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v7/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v7/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newFeedItemSetLinkClientHook clientHook

// FeedItemSetLinkCallOptions contains the retry settings for each method of FeedItemSetLinkClient.
type FeedItemSetLinkCallOptions struct {
	GetFeedItemSetLink     []gax.CallOption
	MutateFeedItemSetLinks []gax.CallOption
}

func defaultFeedItemSetLinkGRPCClientOptions() []option.ClientOption {
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

func defaultFeedItemSetLinkCallOptions() *FeedItemSetLinkCallOptions {
	return &FeedItemSetLinkCallOptions{
		GetFeedItemSetLink: []gax.CallOption{
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
		MutateFeedItemSetLinks: []gax.CallOption{
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

// internalFeedItemSetLinkClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedItemSetLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedItemSetLink(context.Context, *servicespb.GetFeedItemSetLinkRequest, ...gax.CallOption) (*resourcespb.FeedItemSetLink, error)
	MutateFeedItemSetLinks(context.Context, *servicespb.MutateFeedItemSetLinksRequest, ...gax.CallOption) (*servicespb.MutateFeedItemSetLinksResponse, error)
}

// FeedItemSetLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage feed item set links.
type FeedItemSetLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedItemSetLinkClient

	// The call options for this service.
	CallOptions *FeedItemSetLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedItemSetLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedItemSetLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedItemSetLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedItemSetLink returns the requested feed item set link in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemSetLinkClient) GetFeedItemSetLink(ctx context.Context, req *servicespb.GetFeedItemSetLinkRequest, opts ...gax.CallOption) (*resourcespb.FeedItemSetLink, error) {
	return c.internalClient.GetFeedItemSetLink(ctx, req, opts...)
}

// MutateFeedItemSetLinks creates, updates, or removes feed item set links.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemSetLinkClient) MutateFeedItemSetLinks(ctx context.Context, req *servicespb.MutateFeedItemSetLinksRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemSetLinksResponse, error) {
	return c.internalClient.MutateFeedItemSetLinks(ctx, req, opts...)
}

// feedItemSetLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedItemSetLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedItemSetLinkClient
	CallOptions **FeedItemSetLinkCallOptions

	// The gRPC API client.
	feedItemSetLinkClient servicespb.FeedItemSetLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedItemSetLinkClient creates a new feed item set link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage feed item set links.
func NewFeedItemSetLinkClient(ctx context.Context, opts ...option.ClientOption) (*FeedItemSetLinkClient, error) {
	clientOpts := defaultFeedItemSetLinkGRPCClientOptions()
	if newFeedItemSetLinkClientHook != nil {
		hookOpts, err := newFeedItemSetLinkClientHook(ctx, clientHookParams{})
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
	client := FeedItemSetLinkClient{CallOptions: defaultFeedItemSetLinkCallOptions()}

	c := &feedItemSetLinkGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		feedItemSetLinkClient: servicespb.NewFeedItemSetLinkServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedItemSetLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedItemSetLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedItemSetLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedItemSetLinkGRPCClient) GetFeedItemSetLink(ctx context.Context, req *servicespb.GetFeedItemSetLinkRequest, opts ...gax.CallOption) (*resourcespb.FeedItemSetLink, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedItemSetLink[0:len((*c.CallOptions).GetFeedItemSetLink):len((*c.CallOptions).GetFeedItemSetLink)], opts...)
	var resp *resourcespb.FeedItemSetLink
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemSetLinkClient.GetFeedItemSetLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *feedItemSetLinkGRPCClient) MutateFeedItemSetLinks(ctx context.Context, req *servicespb.MutateFeedItemSetLinksRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemSetLinksResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateFeedItemSetLinks[0:len((*c.CallOptions).MutateFeedItemSetLinks):len((*c.CallOptions).MutateFeedItemSetLinks)], opts...)
	var resp *servicespb.MutateFeedItemSetLinksResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemSetLinkClient.MutateFeedItemSetLinks(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
