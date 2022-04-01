package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcallbackrequest
type Patchcallbackrequest struct { 
	// ConversationId - The conversationId.
	ConversationId *string `json:"conversationId,omitempty"`


	// QueueId - The identifier of the queue to be used for the callback.
	QueueId *string `json:"queueId,omitempty"`


	// AgentId - The agentId.
	AgentId *string `json:"agentId,omitempty"`


	// CallbackScheduledTime - The scheduled date-time for the callback. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

}

func (o *Patchcallbackrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchcallbackrequest
	
	CallbackScheduledTime := new(string)
	if o.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(o.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		QueueId: o.QueueId,
		
		AgentId: o.AgentId,
		
		CallbackScheduledTime: CallbackScheduledTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchcallbackrequest) UnmarshalJSON(b []byte) error {
	var PatchcallbackrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchcallbackrequestMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := PatchcallbackrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if QueueId, ok := PatchcallbackrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if AgentId, ok := PatchcallbackrequestMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
	
	if callbackScheduledTimeString, ok := PatchcallbackrequestMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchcallbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
