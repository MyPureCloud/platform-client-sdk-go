package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogicalinterfaceschangetopicerrorinfo
type Edgelogicalinterfaceschangetopicerrorinfo struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`

}

func (u *Edgelogicalinterfaceschangetopicerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogicalinterfaceschangetopicerrorinfo

	

	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		Code *string `json:"code,omitempty"`
		*Alias
	}{ 
		Message: u.Message,
		
		MessageWithParams: u.MessageWithParams,
		
		MessageParams: u.MessageParams,
		
		Code: u.Code,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
