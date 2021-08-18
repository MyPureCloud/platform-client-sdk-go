package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Errordetails
type Errordetails struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Nested
	Nested **Errordetails `json:"nested,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`

}

func (u *Errordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Errordetails

	

	return json.Marshal(&struct { 
		Status *int `json:"status,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Nested **Errordetails `json:"nested,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		Message: u.Message,
		
		MessageWithParams: u.MessageWithParams,
		
		MessageParams: u.MessageParams,
		
		Code: u.Code,
		
		ContextId: u.ContextId,
		
		Nested: u.Nested,
		
		Details: u.Details,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Errordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
