package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
