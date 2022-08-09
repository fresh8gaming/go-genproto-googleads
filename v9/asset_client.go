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

var newAssetClientHook clientHook

// AssetCallOptions contains the retry settings for each method of AssetClient.
type AssetCallOptions struct {
	GetAsset     []gax.CallOption
	MutateAssets []gax.CallOption
}

func defaultAssetGRPCClientOptions() []option.ClientOption {
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

func defaultAssetCallOptions() *AssetCallOptions {
	return &AssetCallOptions{
		GetAsset: []gax.CallOption{
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
		MutateAssets: []gax.CallOption{
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

// internalAssetClient is an interface that defines the methods available from Google Ads API.
type internalAssetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAsset(context.Context, *servicespb.GetAssetRequest, ...gax.CallOption) (*resourcespb.Asset, error)
	MutateAssets(context.Context, *servicespb.MutateAssetsRequest, ...gax.CallOption) (*servicespb.MutateAssetsResponse, error)
}

// AssetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage assets. Asset types can be created with AssetService are
// YoutubeVideoAsset, MediaBundleAsset and ImageAsset. TextAsset should be
// created with Ad inline.
type AssetClient struct {
	// The internal transport-dependent client.
	internalClient internalAssetClient

	// The call options for this service.
	CallOptions *AssetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AssetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AssetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AssetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAsset returns the requested asset in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AssetClient) GetAsset(ctx context.Context, req *servicespb.GetAssetRequest, opts ...gax.CallOption) (*resourcespb.Asset, error) {
	return c.internalClient.GetAsset(ctx, req, opts...)
}

// MutateAssets creates assets. Operation statuses are returned.
//
// List of thrown errors:
// AssetError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// CurrencyCodeError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// ListOperationError (at )
// MediaUploadError (at )
// MutateError (at )
// NotAllowlistedError (at )
// NotEmptyError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
// YoutubeVideoRegistrationError (at )
func (c *AssetClient) MutateAssets(ctx context.Context, req *servicespb.MutateAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateAssetsResponse, error) {
	return c.internalClient.MutateAssets(ctx, req, opts...)
}

// assetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type assetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AssetClient
	CallOptions **AssetCallOptions

	// The gRPC API client.
	assetClient servicespb.AssetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAssetClient creates a new asset service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage assets. Asset types can be created with AssetService are
// YoutubeVideoAsset, MediaBundleAsset and ImageAsset. TextAsset should be
// created with Ad inline.
func NewAssetClient(ctx context.Context, opts ...option.ClientOption) (*AssetClient, error) {
	clientOpts := defaultAssetGRPCClientOptions()
	if newAssetClientHook != nil {
		hookOpts, err := newAssetClientHook(ctx, clientHookParams{})
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
	client := AssetClient{CallOptions: defaultAssetCallOptions()}

	c := &assetGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		assetClient:      servicespb.NewAssetServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *assetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *assetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *assetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *assetGRPCClient) GetAsset(ctx context.Context, req *servicespb.GetAssetRequest, opts ...gax.CallOption) (*resourcespb.Asset, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAsset[0:len((*c.CallOptions).GetAsset):len((*c.CallOptions).GetAsset)], opts...)
	var resp *resourcespb.Asset
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.assetClient.GetAsset(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *assetGRPCClient) MutateAssets(ctx context.Context, req *servicespb.MutateAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateAssetsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAssets[0:len((*c.CallOptions).MutateAssets):len((*c.CallOptions).MutateAssets)], opts...)
	var resp *servicespb.MutateAssetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.assetClient.MutateAssets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
