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

func (o *Queueusereventtopicqueuemember) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		User: o.User,
		
		QueueId: o.QueueId,
		
		Joined: o.Joined,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueusereventtopicqueuemember) UnmarshalJSON(b []byte) error {
	var QueueusereventtopicqueuememberMap map[string]interface{}
	err := json.Unmarshal(b, &QueueusereventtopicqueuememberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueusereventtopicqueuememberMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := QueueusereventtopicqueuememberMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if QueueId, ok := QueueusereventtopicqueuememberMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if Joined, ok := QueueusereventtopicqueuememberMap["joined"].(bool); ok {
		o.Joined = &Joined
	}
	
	if AdditionalProperties, ok := QueueusereventtopicqueuememberMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueusereventtopicqueuemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
