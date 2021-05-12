package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Destination
type Destination struct { 
	// Address - Address or phone number.
	Address *string `json:"address,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Destination) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
