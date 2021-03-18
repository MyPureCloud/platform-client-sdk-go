package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Messageinfo
type Messageinfo struct { 
	// LocalizableMessageCode - Key that can be used to localize the message.
	LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`


	// Message - Description of the message.
	Message *string `json:"message,omitempty"`


	// MessageWithParams - Message with template fields for variable replacement.
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams - Map with fields for variable replacement.
	MessageParams *map[string]string `json:"messageParams,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messageinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
