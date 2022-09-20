package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trigger - Defines a process automation trigger.
type Trigger struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the trigger
	Name *string `json:"name,omitempty"`


	// TopicName - The topic that will cause the trigger to be invoked
	TopicName *string `json:"topicName,omitempty"`


	// Target - The target to invoke when a matching event is received
	Target *Triggertarget `json:"target,omitempty"`


	// Version - Version of this trigger
	Version *int `json:"version,omitempty"`


	// Enabled - Whether or not the trigger is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MatchCriteria - The configuration for when a trigger is considered to be a match for an event
	MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`


	// EventTTLSeconds - How long each event is meaningful after origination, in seconds. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely.
	EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`


	// Description - Description of the trigger. Can be up to 512 characters in length.
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Trigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trigger
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TopicName *string `json:"topicName,omitempty"`
		
		Target *Triggertarget `json:"target,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`
		
		EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TopicName: o.TopicName,
		
		Target: o.Target,
		
		Version: o.Version,
		
		Enabled: o.Enabled,
		
		MatchCriteria: o.MatchCriteria,
		
		EventTTLSeconds: o.EventTTLSeconds,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Trigger) UnmarshalJSON(b []byte) error {
	var TriggerMap map[string]interface{}
	err := json.Unmarshal(b, &TriggerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TriggerMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TriggerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TopicName, ok := TriggerMap["topicName"].(string); ok {
		o.TopicName = &TopicName
	}
    
	if Target, ok := TriggerMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if Version, ok := TriggerMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Enabled, ok := TriggerMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MatchCriteria, ok := TriggerMap["matchCriteria"].([]interface{}); ok {
		MatchCriteriaString, _ := json.Marshal(MatchCriteria)
		json.Unmarshal(MatchCriteriaString, &o.MatchCriteria)
	}
	
	if EventTTLSeconds, ok := TriggerMap["eventTTLSeconds"].(float64); ok {
		EventTTLSecondsInt := int(EventTTLSeconds)
		o.EventTTLSeconds = &EventTTLSecondsInt
	}
	
	if Description, ok := TriggerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SelfUri, ok := TriggerMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Trigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
