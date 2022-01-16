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
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v8/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v8/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCampaignLabelClientHook clientHook

// CampaignLabelCallOptions contains the retry settings for each method of CampaignLabelClient.
type CampaignLabelCallOptions struct {
	GetCampaignLabel     []gax.CallOption
	MutateCampaignLabels []gax.CallOption
}

func defaultCampaignLabelGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignLabelCallOptions() *CampaignLabelCallOptions {
	return &CampaignLabelCallOptions{
		GetCampaignLabel: []gax.CallOption{
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
		MutateCampaignLabels: []gax.CallOption{
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

// internalCampaignLabelClient is an interface that defines the methods availaible from Google Ads API.
type internalCampaignLabelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCampaignLabel(context.Context, *servicespb.GetCampaignLabelRequest, ...gax.CallOption) (*resourcespb.CampaignLabel, error)
	MutateCampaignLabels(context.Context, *servicespb.MutateCampaignLabelsRequest, ...gax.CallOption) (*servicespb.MutateCampaignLabelsResponse, error)
}

// CampaignLabelClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage labels on campaigns.
type CampaignLabelClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignLabelClient

	// The call options for this service.
	CallOptions *CampaignLabelCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignLabelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignLabelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CampaignLabelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCampaignLabel returns the requested campaign-label relationship in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignLabelClient) GetCampaignLabel(ctx context.Context, req *servicespb.GetCampaignLabelRequest, opts ...gax.CallOption) (*resourcespb.CampaignLabel, error) {
	return c.internalClient.GetCampaignLabel(ctx, req, opts...)
}

// MutateCampaignLabels creates and removes campaign-label relationships.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// LabelError (at )
// MutateError (at )
// NewResourceCreationError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignLabelClient) MutateCampaignLabels(ctx context.Context, req *servicespb.MutateCampaignLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignLabelsResponse, error) {
	return c.internalClient.MutateCampaignLabels(ctx, req, opts...)
}

// campaignLabelGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignLabelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CampaignLabelClient
	CallOptions **CampaignLabelCallOptions

	// The gRPC API client.
	campaignLabelClient servicespb.CampaignLabelServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCampaignLabelClient creates a new campaign label service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage labels on campaigns.
func NewCampaignLabelClient(ctx context.Context, opts ...option.ClientOption) (*CampaignLabelClient, error) {
	clientOpts := defaultCampaignLabelGRPCClientOptions()
	if newCampaignLabelClientHook != nil {
		hookOpts, err := newCampaignLabelClientHook(ctx, clientHookParams{})
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
	client := CampaignLabelClient{CallOptions: defaultCampaignLabelCallOptions()}

	c := &campaignLabelGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		campaignLabelClient: servicespb.NewCampaignLabelServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *campaignLabelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignLabelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignLabelGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignLabelGRPCClient) GetCampaignLabel(ctx context.Context, req *servicespb.GetCampaignLabelRequest, opts ...gax.CallOption) (*resourcespb.CampaignLabel, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCampaignLabel[0:len((*c.CallOptions).GetCampaignLabel):len((*c.CallOptions).GetCampaignLabel)], opts...)
	var resp *resourcespb.CampaignLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignLabelClient.GetCampaignLabel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignLabelGRPCClient) MutateCampaignLabels(ctx context.Context, req *servicespb.MutateCampaignLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignLabelsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCampaignLabels[0:len((*c.CallOptions).MutateCampaignLabels):len((*c.CallOptions).MutateCampaignLabels)], opts...)
	var resp *servicespb.MutateCampaignLabelsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignLabelClient.MutateCampaignLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
