package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Requestconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Requestconfig

	

	return json.Marshal(&struct { 
		RequestUrlTemplate *string `json:"requestUrlTemplate,omitempty"`
		
		RequestTemplate *string `json:"requestTemplate,omitempty"`
		
		RequestTemplateUri *string `json:"requestTemplateUri,omitempty"`
		
		RequestType *string `json:"requestType,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		*Alias
	}{ 
		RequestUrlTemplate: u.RequestUrlTemplate,
		
		RequestTemplate: u.RequestTemplate,
		
		RequestTemplateUri: u.RequestTemplateUri,
		
		RequestType: u.RequestType,
		
		Headers: u.Headers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Requestconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
