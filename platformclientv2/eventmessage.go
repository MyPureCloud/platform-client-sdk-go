package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventmessage
type Eventmessage struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]interface{} `json:"messageParams,omitempty"`


	// DocumentationUri
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// ResourceURIs
	ResourceURIs *[]string `json:"resourceURIs,omitempty"`

}

func (u *Eventmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventmessage

	

	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]interface{} `json:"messageParams,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		ResourceURIs *[]string `json:"resourceURIs,omitempty"`
		*Alias
	}{ 
		Code: u.Code,
		
		Message: u.Message,
		
		MessageWithParams: u.MessageWithParams,
		
		MessageParams: u.MessageParams,
		
		DocumentationUri: u.DocumentationUri,
		
		ResourceURIs: u.ResourceURIs,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Eventmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
