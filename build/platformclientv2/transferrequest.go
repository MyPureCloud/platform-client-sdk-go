package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transferrequest
type Transferrequest struct { 
	// UserId - The user ID of the transfer target.
	UserId *string `json:"userId,omitempty"`


	// Address - The phone number or address of the transfer target.
	Address *string `json:"address,omitempty"`


	// UserName - The user name of the transfer target.
	UserName *string `json:"userName,omitempty"`


	// QueueId - The queue ID of the transfer target.
	QueueId *string `json:"queueId,omitempty"`


	// Voicemail - If true, transfer to the voicemail inbox of the participant that is being replaced.
	Voicemail *bool `json:"voicemail,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transferrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
