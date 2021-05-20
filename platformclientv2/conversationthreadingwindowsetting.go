package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Conversationthreadingwindowsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
