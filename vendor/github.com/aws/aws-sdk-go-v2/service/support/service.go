// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Support provides the API operation methods for making requests to
// AWS Support. See this package's package overview docs
// for details on the service.
//
// Support methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Support struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Support)

// Used for custom request initialization logic
var initRequest func(*Support, *aws.Request)

// Service information constants
const (
	ServiceName = "support"   // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Support client with a config.
//
// Example:
//     // Create a Support client from just a config.
//     svc := support.New(myConfig)
func New(config aws.Config) *Support {
	var signingName string
	signingRegion := config.Region

	svc := &Support{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2013-04-15",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSSupport_20130415",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Support operation and runs any
// custom request initialization.
func (c *Support) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
