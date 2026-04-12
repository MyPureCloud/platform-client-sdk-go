package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatescheduledtriggerrequest
type Updatescheduledtriggerrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - Version of this scheduled trigger
	Version *int `json:"version,omitempty"`

	// Enabled - Boolean indicating if scheduled trigger is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Target - The target to invoke when the scheduled trigger fires
	Target *Triggertarget `json:"target,omitempty"`

	// Name - The name of the scheduled trigger. Can be up to 162 characters in length.
	Name *string `json:"name,omitempty"`

	// Schedule - The schedule configuration for when this trigger should fire
	Schedule *Triggerschedule `json:"schedule,omitempty"`

	// Description - Description of the trigger. Can be up to 512 characters in length.
	Description *string `json:"description,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatescheduledtriggerrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updatescheduledtriggerrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Updatescheduledtriggerrequest
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Target *Triggertarget `json:"target,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Schedule *Triggerschedule `json:"schedule,omitempty"`
		
		Description *string `json:"description,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		Enabled: o.Enabled,
		
		Target: o.Target,
		
		Name: o.Name,
		
		Schedule: o.Schedule,
		
		Description: o.Description,
		Alias:    (Alias)(o),
	})
}

func (o *Updatescheduledtriggerrequest) UnmarshalJSON(b []byte) error {
	var UpdatescheduledtriggerrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatescheduledtriggerrequestMap)
	if err != nil {
		return err
	}
	
	if Version, ok := UpdatescheduledtriggerrequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Enabled, ok := UpdatescheduledtriggerrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Target, ok := UpdatescheduledtriggerrequestMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if Name, ok := UpdatescheduledtriggerrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Schedule, ok := UpdatescheduledtriggerrequestMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if Description, ok := UpdatescheduledtriggerrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updatescheduledtriggerrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
