package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Queueusereventtopicqueuemember) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueusereventtopicqueuemember

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Queueusereventtopicuserreference `json:"user,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		Joined *bool `json:"joined,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		User: u.User,
		
		QueueId: u.QueueId,
		
		Joined: u.Joined,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueusereventtopicqueuemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
