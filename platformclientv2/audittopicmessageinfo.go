package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicmessageinfo
type Audittopicmessageinfo struct { 
	// LocalizableMessageCode
	LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`

}

// String returns a JSON representation of the model
func (o *Audittopicmessageinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
