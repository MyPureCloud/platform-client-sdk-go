package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatetriggerrequest
type Updatetriggerrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - Version of this trigger
	Version *int `json:"version,omitempty"`

	// Enabled - Boolean indicating if Trigger is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Target - The target to invoke when a matching event is received
	Target *Triggertarget `json:"target,omitempty"`

	// MatchCriteria - The configuration for when a trigger is considered to be a match for an event
	MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`

	// Name - The name of the trigger
	Name *string `json:"name,omitempty"`

	// TopicName - The topic that will cause the trigger to be invoked. Must match existing trigger topicName.
	TopicName *string `json:"topicName,omitempty"`

	// EventTTLSeconds - Optional length of time that events are meaningful after origination. Events older than this threshold may be dropped if the platform is delayed in processing events. Unset means events are valid indefinitely, otherwise must be set to at least 10 seconds. Only one of eventTTLSeconds or delayBySeconds can be set.
	EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`

	// DelayBySeconds - Optional delay invoking target after trigger fires. Must be in the range of 60 to 900 seconds. Only one of eventTTLSeconds or delayBySeconds can be set. Until delayed triggers are released supplying this attribute will cause a failure.
	DelayBySeconds *int `json:"delayBySeconds,omitempty"`

	// Description - Description of the trigger. Can be up to 512 characters in length.
	Description *string `json:"description,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatetriggerrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updatetriggerrequest) MarshalJSON() ([]byte, error) {
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
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Updatetriggerrequest
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Target *Triggertarget `json:"target,omitempty"`
		
		MatchCriteria *[]Matchcriteria `json:"matchCriteria,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TopicName *string `json:"topicName,omitempty"`
		
		EventTTLSeconds *int `json:"eventTTLSeconds,omitempty"`
		
		DelayBySeconds *int `json:"delayBySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		Enabled: o.Enabled,
		
		Target: o.Target,
		
		MatchCriteria: o.MatchCriteria,
		
		Name: o.Name,
		
		TopicName: o.TopicName,
		
		EventTTLSeconds: o.EventTTLSeconds,
		
		DelayBySeconds: o.DelayBySeconds,
		
		Description: o.Description,
		Alias:    (Alias)(o),
	})
}

func (o *Updatetriggerrequest) UnmarshalJSON(b []byte) error {
	var UpdatetriggerrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatetriggerrequestMap)
	if err != nil {
		return err
	}
	
	if Version, ok := UpdatetriggerrequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Enabled, ok := UpdatetriggerrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Target, ok := UpdatetriggerrequestMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if MatchCriteria, ok := UpdatetriggerrequestMap["matchCriteria"].([]interface{}); ok {
		MatchCriteriaString, _ := json.Marshal(MatchCriteria)
		json.Unmarshal(MatchCriteriaString, &o.MatchCriteria)
	}
	
	if Name, ok := UpdatetriggerrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TopicName, ok := UpdatetriggerrequestMap["topicName"].(string); ok {
		o.TopicName = &TopicName
	}
    
	if EventTTLSeconds, ok := UpdatetriggerrequestMap["eventTTLSeconds"].(float64); ok {
		EventTTLSecondsInt := int(EventTTLSeconds)
		o.EventTTLSeconds = &EventTTLSecondsInt
	}
	
	if DelayBySeconds, ok := UpdatetriggerrequestMap["delayBySeconds"].(float64); ok {
		DelayBySecondsInt := int(DelayBySeconds)
		o.DelayBySeconds = &DelayBySecondsInt
	}
	
	if Description, ok := UpdatetriggerrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updatetriggerrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
