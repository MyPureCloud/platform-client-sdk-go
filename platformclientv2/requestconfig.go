package platformclientv2
import (
	"encoding/json"
)

// Requestconfig - Defines response components of the Action Request.
type Requestconfig struct { 
	// RequestUrlTemplate - URL that may include placeholders for requests to 3rd party service
	RequestUrlTemplate *string `json:"requestUrlTemplate,omitempty"`


	// RequestTemplate - Velocity template to define request body sent to 3rd party service.
	RequestTemplate *string `json:"requestTemplate,omitempty"`


	// RequestTemplateUri - URI to retrieve requestTemplate
	RequestTemplateUri *string `json:"requestTemplateUri,omitempty"`


	// RequestType - HTTP method to use for request
	RequestType *string `json:"requestType,omitempty"`


	// Headers - Headers to include in request in (Header Name, Value) pairs.
	Headers *map[string]string `json:"headers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Requestconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
