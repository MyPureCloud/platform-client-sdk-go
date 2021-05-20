package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueusereventtopicqueuemember
type Queueusereventtopicqueuemember struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// User
	User *Queueusereventtopicuserreference `json:"user,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// Joined
	Joined *bool `json:"joined,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueusereventtopicqueuemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
