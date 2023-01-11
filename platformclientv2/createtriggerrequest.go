package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createtriggerrequest
type Createtriggerrequest struct { 
	// Target - The target to invoke when a matching event is received
	Target *Triggertarget `json:"target,omitempty"`


	// Enabled - Boolean indicating if Trigger is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MatchCriteria - The configuration for when a trigger is considered to be a match for an event. When not provided, all events will fire the trigger
	MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`


	// Name - The name of the trigger
	Name *string `json:"name,omitempty"`


	// TopicName - The topic that will cause the trigger to be invoked. Cannot be updated after creation. Valid topics can be found at /processautomation/triggers/topics 
	TopicName *string `json:"topicName,omitempty"`


	// EventTTLSeconds - Optional length of time that events are meaningful after origination. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely, otherwise must be set to at least 10 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
	EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`


	// DelayBySeconds - Optional delay invoking target after trigger fires. Must be in the range of 60 to 900 seconds. Only one of eventTTLSeconds or delayBySeconds can be set. Until delayed triggers are released supplying this attribute will cause a failure.
	DelayBySeconds *int `json:"delayBySeconds,omitempty"`


	// Description - Description of the trigger. Can be up to 512 characters in length.
	Description *string `json:"description,omitempty"`

}

func (o *Createtriggerrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createtriggerrequest
	
	return json.Marshal(&struct { 
		Target *Triggertarget `json:"target,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TopicName *string `json:"topicName,omitempty"`
		
		EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`
		
		DelayBySeconds *int `json:"delayBySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Target: o.Target,
		
		Enabled: o.Enabled,
		
		MatchCriteria: o.MatchCriteria,
		
		Name: o.Name,
		
		TopicName: o.TopicName,
		
		EventTTLSeconds: o.EventTTLSeconds,
		
		DelayBySeconds: o.DelayBySeconds,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Createtriggerrequest) UnmarshalJSON(b []byte) error {
	var CreatetriggerrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatetriggerrequestMap)
	if err != nil {
		return err
	}
	
	if Target, ok := CreatetriggerrequestMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if Enabled, ok := CreatetriggerrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MatchCriteria, ok := CreatetriggerrequestMap["matchCriteria"].([]interface{}); ok {
		MatchCriteriaString, _ := json.Marshal(MatchCriteria)
		json.Unmarshal(MatchCriteriaString, &o.MatchCriteria)
	}
	
	if Name, ok := CreatetriggerrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TopicName, ok := CreatetriggerrequestMap["topicName"].(string); ok {
		o.TopicName = &TopicName
	}
    
	if EventTTLSeconds, ok := CreatetriggerrequestMap["eventTTLSeconds"].(float64); ok {
		EventTTLSecondsInt := int(EventTTLSeconds)
		o.EventTTLSeconds = &EventTTLSecondsInt
	}
	
	if DelayBySeconds, ok := CreatetriggerrequestMap["delayBySeconds"].(float64); ok {
		DelayBySecondsInt := int(DelayBySeconds)
		o.DelayBySeconds = &DelayBySecondsInt
	}
	
	if Description, ok := CreatetriggerrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createtriggerrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
