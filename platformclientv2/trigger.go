package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trigger - Defines a process automation trigger.
type Trigger struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// EventTTLSeconds - Optional length of time that events are meaningful after origination. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely, otherwise must be set to at least 10 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
	EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`

	// DelayBySeconds - Optional delay invoking target after trigger fires. Must be in the range of 60 to 900 seconds. Only one of eventTTLSeconds or delayBySeconds can be set. Until delayed triggers are released supplying this attribute will cause a failure.
	DelayBySeconds *int `json:"delayBySeconds,omitempty"`

	// Description - Description of the trigger. Can be up to 512 characters in length.
	Description *string `json:"description,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trigger) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Trigger) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		DelayBySeconds *int `json:"delayBySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TopicName: o.TopicName,
		
		Target: o.Target,
		
		Version: o.Version,
		
		Enabled: o.Enabled,
		
		MatchCriteria: o.MatchCriteria,
		
		EventTTLSeconds: o.EventTTLSeconds,
		
		DelayBySeconds: o.DelayBySeconds,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
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
	
	if DelayBySeconds, ok := TriggerMap["delayBySeconds"].(float64); ok {
		DelayBySecondsInt := int(DelayBySeconds)
		o.DelayBySeconds = &DelayBySecondsInt
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
