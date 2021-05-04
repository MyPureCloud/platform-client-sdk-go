package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsresolution
type Analyticsresolution struct { 
	// EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// QueueId - The ID of the last queue on which the conversation was handled.
	QueueId *string `json:"queueId,omitempty"`


	// UserId - The ID of the last user who handled the conversation.
	UserId *string `json:"userId,omitempty"`


	// GetnNextContactAvoided - The number of interactions for which next contact was avoided.
	GetnNextContactAvoided *int `json:"getnNextContactAvoided,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsresolution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
