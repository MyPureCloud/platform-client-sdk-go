package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationresult
type Schedulegenerationresult struct { 
	// Failed - Whether the schedule generation run failed
	Failed *bool `json:"failed,omitempty"`


	// RunId - The ID of the schedule generation run. Reference this when requesting support
	RunId *string `json:"runId,omitempty"`


	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`


	// Messages - User facing messages related to the schedule generation run
	Messages *[]Schedulegenerationmessage `json:"messages,omitempty"`


	// MessageSeverities - The list of messages by severity in this schedule generation run
	MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`

}

func (o *Schedulegenerationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegenerationresult
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		
		Messages *[]Schedulegenerationmessage `json:"messages,omitempty"`
		
		MessageSeverities *[]Schedulermessagetypeseverity `json:"messageSeverities,omitempty"`
		*Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		MessageCount: o.MessageCount,
		
		Messages: o.Messages,
		
		MessageSeverities: o.MessageSeverities,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulegenerationresult) UnmarshalJSON(b []byte) error {
	var SchedulegenerationresultMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegenerationresultMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := SchedulegenerationresultMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
    
	if RunId, ok := SchedulegenerationresultMap["runId"].(string); ok {
		o.RunId = &RunId
	}
    
	if MessageCount, ok := SchedulegenerationresultMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if Messages, ok := SchedulegenerationresultMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if MessageSeverities, ok := SchedulegenerationresultMap["messageSeverities"].([]interface{}); ok {
		MessageSeveritiesString, _ := json.Marshal(MessageSeverities)
		json.Unmarshal(MessageSeveritiesString, &o.MessageSeverities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegenerationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
