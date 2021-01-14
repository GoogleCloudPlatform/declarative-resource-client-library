// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dcl

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	// glog aliased import is necessary since these packages will be open-sourced
	// and that is the public name of the google logging package.
	glog "github.com/golang/glog"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const ua = "DeclarativeClientLib/0.0.1"

// Config is used to enclose the credentials and http client used to make
// requests to GCP APIs.
type Config struct {
	Retry         Retry
	header        http.Header
	clientOptions []option.ClientOption
	userAgent     string
	Logger        Logger
	BasePath      string
}

// UserAgent returns the user agent for the config, which will always include the
// declarative SDK name and version.
func (c *Config) UserAgent() string {
	return fmt.Sprintf("%s %s", c.userAgent, ua)
}

// NewConfig creates a Config object with the specified user agent and client.
func NewConfig(o ...ClientOption) *Config {
	return &Config{
		Retry:         fetchRetry(o),
		header:        fetchHeader(o),
		userAgent:     fetchUserAgent(o),
		Logger:        fetchLogger(o),
		BasePath:      fetchBasePath(o),
		clientOptions: fetchClientOptions(o),
	}
}

type loggingTransport struct {
	underlyingTransport http.RoundTripper
	logger              Logger
}

func (t loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		t.logger.Infof("Google API Request:\n-----------[REQUEST]----------\n%s\n-------[END REQUEST]--------", reqDump)
	} else {
		t.logger.Warningf("Failed to make request: %s", err)
	}
	resp, err := t.underlyingTransport.RoundTrip(req)
	if err == nil {
		respDump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			t.logger.Infof("Google API Response:\n-----------[RESPONSE]----------\n%s\n-------[END RESPONSE]--------", respDump)
		} else {
			t.logger.Warningf("Failed to parse response: %s", err)
		}
	} else {
		t.logger.Warningf("Failed to get response: %s", err)
	}
	return resp, err
}

// ApplyOption is an option that is accepted by Apply() functions.
type ApplyOption interface {
	apply(*applyOpts)
}

// ClientOption is an option that is accepted in client creation.
type ClientOption interface {
	apply(*clientOpts)
}

// applyOpts refers to options that are taken in the apply function.
type applyOpts struct {
	params    []LifecycleParam
	stateHint Resource
}

type lifecycleParamOption struct {
	param LifecycleParam
}

func (l lifecycleParamOption) apply(o *applyOpts) {
	o.params = append(o.params, l.param)
}

// WithLifecycleParam allows a user to specify the proper lifecycle params.
func WithLifecycleParam(d LifecycleParam) ApplyOption {
	return lifecycleParamOption{param: d}
}

// FetchLifecycleParams returns the list of lifecycle params.
func FetchLifecycleParams(c []ApplyOption) []LifecycleParam {
	var o applyOpts
	for _, p := range c {
		p.apply(&o)
	}
	return o.params
}

type stateHint struct {
	state Resource
}

func (s stateHint) apply(o *applyOpts) {
	o.stateHint = s.state
}

// WithStateHint takes in a resource which will be used in place of the applied
// resource any time the current configuration of the resource is relevant.
// For instance, if an identity field will change, passing a state hint will ensure
// that the current resource is fetched (and possibly deleted).
func WithStateHint(r Resource) ApplyOption {
	return stateHint{state: r}
}

// FetchStateHint returns either nil or a dcl.Resource representing the pre-apply state.
func FetchStateHint(c []ApplyOption) Resource {
	var o applyOpts
	for _, p := range c {
		p.apply(&o)
	}
	return o.stateHint
}

// clientOpts refers to options passed into the client.
type clientOpts struct {
	backoff       Retry
	header        http.Header
	userAgent     string
	logger        Logger
	basepath      string
	clientOptions []option.ClientOption
}

type timeoutOption struct {
	backoff Retry
}

func (t timeoutOption) apply(o *clientOpts) {
	o.backoff = t.backoff
}

type headerOption struct {
	header string
	value  string
}

func (h headerOption) apply(o *clientOpts) {
	if o.header == nil {
		o.header = make(http.Header)
	}
	o.header.Add(h.header, h.value)
}

type userAgentOption struct {
	userAgent string
}

func (u userAgentOption) apply(o *clientOpts) {
	o.userAgent = u.userAgent
}

type loggerOption struct {
	logger Logger
}

func (l loggerOption) apply(o *clientOpts) {
	o.logger = l.logger
}

type basepathOption struct {
	basepath string
}

func (l basepathOption) apply(o *clientOpts) {
	o.basepath = l.basepath
}

type apiClientOption struct {
	o option.ClientOption
}

func (l apiClientOption) apply(o *clientOpts) {
	o.clientOptions = append(o.clientOptions, l.o)
}

// WithRetry allows a user to specify the proper backoff for a Apply()
func WithRetry(b Retry) ClientOption {
	return timeoutOption{backoff: b}
}

// WithLogger allows a user to specify a custom logger.
func WithLogger(l Logger) ClientOption {
	return loggerOption{logger: l}
}

// WithHeader allows aribitrary HTTP headers to be addded to requests. Not all headers
// (e.g., "Content-Type") can be overridden. To set the User-Agent header, use WithUserAgent().
func WithHeader(header, value string) ClientOption {
	return headerOption{
		header: header,
		value:  value,
	}
}

// WithUserAgent allows a user to specify a custom user-agent.
func WithUserAgent(ua string) ClientOption {
	return userAgentOption{userAgent: ua}
}

// WithAPIKey returns a ClientOption that specifies an API key to be used as the basis for authentication.
func WithAPIKey(apiKey string) ClientOption {
	return apiClientOption{o: option.WithAPIKey(apiKey)}
}

// WithClientCertSource returns a ClientOption that specifies a callback function for obtaining a TLS client certificate.
func WithClientCertSource(s option.ClientCertSource) ClientOption {
	return apiClientOption{o: option.WithClientCertSource(s)}
}

// WithCredentials returns a ClientOption that authenticates API calls using a caller-supplier Credentials struct.
func WithCredentials(creds *google.Credentials) ClientOption {
	return apiClientOption{o: option.WithCredentials(creds)}
}

// WithCredentialsFile returns a ClientOption that authenticates API calls with the given service account or refresh token JSON credentials file.
func WithCredentialsFile(filename string) ClientOption {
	return apiClientOption{o: option.WithCredentialsFile(filename)}
}

// WithCredentialsJSON returns a ClientOption that authenticates API calls with the given service account or refresh token JSON credentials.
func WithCredentialsJSON(p []byte) ClientOption {
	return apiClientOption{o: option.WithCredentialsJSON(p)}
}

// WithHTTPClient returns a ClientOption that specifies the HTTP client to use as the basis of communications.
// When used, the WithHTTPClient option takes precedent over all other supplied authentication options.
func WithHTTPClient(client *http.Client) ClientOption {
	return apiClientOption{o: option.WithHTTPClient(client)}
}

// fetchRetry returns either a new Retry a user-specified retry.
func fetchRetry(c []ClientOption) Retry {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}

	if o.backoff != nil {
		return o.backoff
	}
	return NewBackoff()
}

// fetchLogger returns either a new Logger or a user-specified logger.
func fetchLogger(c []ClientOption) Logger {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}
	if o.logger == nil {
		return DefaultLogger()
	}
	return o.logger
}

// fetchHeader returns user-specified HTTP request headers.
func fetchHeader(c []ClientOption) http.Header {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}
	return o.header
}

// fetchUserAgent returns a user-specified or default user-agent string.
func fetchUserAgent(c []ClientOption) string {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}

	if o.userAgent == "" {
		return "declarative-client-lib"
	}
	return o.userAgent
}

// fetchBasePath returns a user-specified base path or empty string.
func fetchBasePath(c []ClientOption) string {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}
	return o.basepath
}

// fetchClientOptions returns the API options to be used for network endpoints.
func fetchClientOptions(c []ClientOption) []option.ClientOption {
	var o clientOpts
	for _, p := range c {
		p.apply(&o)
	}
	return o.clientOptions
}

// Logger is an interface for logging requests and responses.
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Warning(args ...interface{})
}

// DefaultLogger returns the default logger for the Declarative Client Library.
func DefaultLogger() Logger {
	return glogger{}
}

type glogger struct{}

// Fatal records Fatal errors.
func (l glogger) Fatal(args ...interface{}) {
	glog.Fatal(args)
}

// Fatalf records Fatal errors with added arguments.
func (l glogger) Fatalf(format string, args ...interface{}) {
	a := make([]interface{}, len(args))
	for i, v := range args {
		if s, ok := v.(*string); ok && s != nil {
			a[i] = *s
		} else {
			a[i] = v
		}
	}

	glog.Fatalf(format, a...)
}

// Info records Info errors.
func (l glogger) Info(args ...interface{}) {
	glog.Info(args)
}

// Infof records Info errors with added arguments.
func (l glogger) Infof(format string, args ...interface{}) {
	a := make([]interface{}, len(args))
	for i, v := range args {
		if s, ok := v.(*string); ok && s != nil {
			a[i] = *s
		} else {
			a[i] = v
		}
	}
	glog.Infof(format, a...)
}

// Warningf records Warning errors with added arguments.
func (l glogger) Warningf(format string, args ...interface{}) {
	a := make([]interface{}, len(args))
	for i, v := range args {
		if s, ok := v.(*string); ok && s != nil {
			a[i] = *s
		} else {
			a[i] = v
		}
	}

	glog.Warningf(format, a...)
}

// Warning records Warning errors.
func (l glogger) Warning(args ...interface{}) {
	glog.Warning(args...)
}
