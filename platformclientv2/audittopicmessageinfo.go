package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Audittopicmessageinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicmessageinfo

	

	return json.Marshal(&struct { 
		LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		*Alias
	}{ 
		LocalizableMessageCode: u.LocalizableMessageCode,
		
		Message: u.Message,
		
		MessageWithParams: u.MessageWithParams,
		
		MessageParams: u.MessageParams,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Audittopicmessageinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
