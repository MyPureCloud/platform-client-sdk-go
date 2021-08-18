package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationthreadingwindowsetting
type Conversationthreadingwindowsetting struct { 
	// MessengerType - The type of messenger
	MessengerType *string `json:"messengerType,omitempty"`


	// TimeoutInMinutes - The conversation threading window timeout (Minutes) of specified messenger type
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`

}

func (u *Conversationthreadingwindowsetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationthreadingwindowsetting

	

	return json.Marshal(&struct { 
		MessengerType *string `json:"messengerType,omitempty"`
		
		TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
		*Alias
	}{ 
		MessengerType: u.MessengerType,
		
		TimeoutInMinutes: u.TimeoutInMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindowsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
