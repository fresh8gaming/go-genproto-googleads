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

var newKeywordPlanClientHook clientHook

// KeywordPlanCallOptions contains the retry settings for each method of KeywordPlanClient.
type KeywordPlanCallOptions struct {
	GetKeywordPlan             []gax.CallOption
	MutateKeywordPlans         []gax.CallOption
	GenerateForecastCurve      []gax.CallOption
	GenerateForecastTimeSeries []gax.CallOption
	GenerateForecastMetrics    []gax.CallOption
	GenerateHistoricalMetrics  []gax.CallOption
}

func defaultKeywordPlanGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordPlanCallOptions() *KeywordPlanCallOptions {
	return &KeywordPlanCallOptions{
		GetKeywordPlan: []gax.CallOption{
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
		MutateKeywordPlans: []gax.CallOption{
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
		GenerateForecastCurve: []gax.CallOption{
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
		GenerateForecastTimeSeries: []gax.CallOption{
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
		GenerateForecastMetrics: []gax.CallOption{
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
		GenerateHistoricalMetrics: []gax.CallOption{
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

// internalKeywordPlanClient is an interface that defines the methods availaible from Google Ads API.
type internalKeywordPlanClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetKeywordPlan(context.Context, *servicespb.GetKeywordPlanRequest, ...gax.CallOption) (*resourcespb.KeywordPlan, error)
	MutateKeywordPlans(context.Context, *servicespb.MutateKeywordPlansRequest, ...gax.CallOption) (*servicespb.MutateKeywordPlansResponse, error)
	GenerateForecastCurve(context.Context, *servicespb.GenerateForecastCurveRequest, ...gax.CallOption) (*servicespb.GenerateForecastCurveResponse, error)
	GenerateForecastTimeSeries(context.Context, *servicespb.GenerateForecastTimeSeriesRequest, ...gax.CallOption) (*servicespb.GenerateForecastTimeSeriesResponse, error)
	GenerateForecastMetrics(context.Context, *servicespb.GenerateForecastMetricsRequest, ...gax.CallOption) (*servicespb.GenerateForecastMetricsResponse, error)
	GenerateHistoricalMetrics(context.Context, *servicespb.GenerateHistoricalMetricsRequest, ...gax.CallOption) (*servicespb.GenerateHistoricalMetricsResponse, error)
}

// KeywordPlanClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage keyword plans.
type KeywordPlanClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordPlanClient

	// The call options for this service.
	CallOptions *KeywordPlanCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *KeywordPlanClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetKeywordPlan returns the requested plan in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanClient) GetKeywordPlan(ctx context.Context, req *servicespb.GetKeywordPlanRequest, opts ...gax.CallOption) (*resourcespb.KeywordPlan, error) {
	return c.internalClient.GetKeywordPlan(ctx, req, opts...)
}

// MutateKeywordPlans creates, updates, or removes keyword plans. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanError (at )
// MutateError (at )
// NewResourceCreationError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// StringLengthError (at )
func (c *KeywordPlanClient) MutateKeywordPlans(ctx context.Context, req *servicespb.MutateKeywordPlansRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlansResponse, error) {
	return c.internalClient.MutateKeywordPlans(ctx, req, opts...)
}

// GenerateForecastCurve returns the requested Keyword Plan forecast curve.
// Only the bidding strategy is considered for generating forecast curve.
// The bidding strategy value specified in the plan is ignored.
//
// To generate a forecast at a value specified in the plan, use
// KeywordPlanService.GenerateForecastMetrics.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanClient) GenerateForecastCurve(ctx context.Context, req *servicespb.GenerateForecastCurveRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastCurveResponse, error) {
	return c.internalClient.GenerateForecastCurve(ctx, req, opts...)
}

// GenerateForecastTimeSeries returns a forecast in the form of a time series for the Keyword Plan over
// the next 52 weeks.
// (1) Forecasts closer to the current date are generally more accurate than
// further out.
//
// (2) The forecast reflects seasonal trends using current and
// prior traffic patterns. The forecast period of the plan is ignored.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanClient) GenerateForecastTimeSeries(ctx context.Context, req *servicespb.GenerateForecastTimeSeriesRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastTimeSeriesResponse, error) {
	return c.internalClient.GenerateForecastTimeSeries(ctx, req, opts...)
}

// GenerateForecastMetrics returns the requested Keyword Plan forecasts.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanClient) GenerateForecastMetrics(ctx context.Context, req *servicespb.GenerateForecastMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastMetricsResponse, error) {
	return c.internalClient.GenerateForecastMetrics(ctx, req, opts...)
}

// GenerateHistoricalMetrics returns the requested Keyword Plan historical metrics.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanClient) GenerateHistoricalMetrics(ctx context.Context, req *servicespb.GenerateHistoricalMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateHistoricalMetricsResponse, error) {
	return c.internalClient.GenerateHistoricalMetrics(ctx, req, opts...)
}

// keywordPlanGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordPlanGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing KeywordPlanClient
	CallOptions **KeywordPlanCallOptions

	// The gRPC API client.
	keywordPlanClient servicespb.KeywordPlanServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewKeywordPlanClient creates a new keyword plan service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage keyword plans.
func NewKeywordPlanClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanClient, error) {
	clientOpts := defaultKeywordPlanGRPCClientOptions()
	if newKeywordPlanClientHook != nil {
		hookOpts, err := newKeywordPlanClientHook(ctx, clientHookParams{})
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
	client := KeywordPlanClient{CallOptions: defaultKeywordPlanCallOptions()}

	c := &keywordPlanGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		keywordPlanClient: servicespb.NewKeywordPlanServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *keywordPlanGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordPlanGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordPlanGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordPlanGRPCClient) GetKeywordPlan(ctx context.Context, req *servicespb.GetKeywordPlanRequest, opts ...gax.CallOption) (*resourcespb.KeywordPlan, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKeywordPlan[0:len((*c.CallOptions).GetKeywordPlan):len((*c.CallOptions).GetKeywordPlan)], opts...)
	var resp *resourcespb.KeywordPlan
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.GetKeywordPlan(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanGRPCClient) MutateKeywordPlans(ctx context.Context, req *servicespb.MutateKeywordPlansRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlansResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateKeywordPlans[0:len((*c.CallOptions).MutateKeywordPlans):len((*c.CallOptions).MutateKeywordPlans)], opts...)
	var resp *servicespb.MutateKeywordPlansResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.MutateKeywordPlans(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanGRPCClient) GenerateForecastCurve(ctx context.Context, req *servicespb.GenerateForecastCurveRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastCurveResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "keyword_plan", url.QueryEscape(req.GetKeywordPlan())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateForecastCurve[0:len((*c.CallOptions).GenerateForecastCurve):len((*c.CallOptions).GenerateForecastCurve)], opts...)
	var resp *servicespb.GenerateForecastCurveResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.GenerateForecastCurve(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanGRPCClient) GenerateForecastTimeSeries(ctx context.Context, req *servicespb.GenerateForecastTimeSeriesRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastTimeSeriesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "keyword_plan", url.QueryEscape(req.GetKeywordPlan())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateForecastTimeSeries[0:len((*c.CallOptions).GenerateForecastTimeSeries):len((*c.CallOptions).GenerateForecastTimeSeries)], opts...)
	var resp *servicespb.GenerateForecastTimeSeriesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.GenerateForecastTimeSeries(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanGRPCClient) GenerateForecastMetrics(ctx context.Context, req *servicespb.GenerateForecastMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateForecastMetricsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "keyword_plan", url.QueryEscape(req.GetKeywordPlan())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateForecastMetrics[0:len((*c.CallOptions).GenerateForecastMetrics):len((*c.CallOptions).GenerateForecastMetrics)], opts...)
	var resp *servicespb.GenerateForecastMetricsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.GenerateForecastMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanGRPCClient) GenerateHistoricalMetrics(ctx context.Context, req *servicespb.GenerateHistoricalMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateHistoricalMetricsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "keyword_plan", url.QueryEscape(req.GetKeywordPlan())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateHistoricalMetrics[0:len((*c.CallOptions).GenerateHistoricalMetrics):len((*c.CallOptions).GenerateHistoricalMetrics)], opts...)
	var resp *servicespb.GenerateHistoricalMetricsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanClient.GenerateHistoricalMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
