package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Policyconditions
type Policyconditions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ForUsers
	ForUsers *[]User `json:"forUsers,omitempty"`

	// Directions
	Directions *[]string `json:"directions,omitempty"`

	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`

	// MediaTypes
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`

	// Duration
	Duration *Durationcondition `json:"duration,omitempty"`

	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`

	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`

	// CustomerParticipation - This condition is to filter out conversation with and without customer participation.
	CustomerParticipation *string `json:"customerParticipation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Policyconditions) SetField(field string, fieldValue interface{}) {
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

func (o Policyconditions) MarshalJSON() ([]byte, error) {
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
	type Alias Policyconditions
	
	return json.Marshal(&struct { 
		ForUsers *[]User `json:"forUsers,omitempty"`
		
		Directions *[]string `json:"directions,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		Duration *Durationcondition `json:"duration,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		
		CustomerParticipation *string `json:"customerParticipation,omitempty"`
		Alias
	}{ 
		ForUsers: o.ForUsers,
		
		Directions: o.Directions,
		
		DateRanges: o.DateRanges,
		
		MediaTypes: o.MediaTypes,
		
		ForQueues: o.ForQueues,
		
		Duration: o.Duration,
		
		WrapupCodes: o.WrapupCodes,
		
		TimeAllowed: o.TimeAllowed,
		
		CustomerParticipation: o.CustomerParticipation,
		Alias:    (Alias)(o),
	})
}

func (o *Policyconditions) UnmarshalJSON(b []byte) error {
	var PolicyconditionsMap map[string]interface{}
	err := json.Unmarshal(b, &PolicyconditionsMap)
	if err != nil {
		return err
	}
	
	if ForUsers, ok := PolicyconditionsMap["forUsers"].([]interface{}); ok {
		ForUsersString, _ := json.Marshal(ForUsers)
		json.Unmarshal(ForUsersString, &o.ForUsers)
	}
	
	if Directions, ok := PolicyconditionsMap["directions"].([]interface{}); ok {
		DirectionsString, _ := json.Marshal(Directions)
		json.Unmarshal(DirectionsString, &o.Directions)
	}
	
	if DateRanges, ok := PolicyconditionsMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	
	if MediaTypes, ok := PolicyconditionsMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if ForQueues, ok := PolicyconditionsMap["forQueues"].([]interface{}); ok {
		ForQueuesString, _ := json.Marshal(ForQueues)
		json.Unmarshal(ForQueuesString, &o.ForQueues)
	}
	
	if Duration, ok := PolicyconditionsMap["duration"].(map[string]interface{}); ok {
		DurationString, _ := json.Marshal(Duration)
		json.Unmarshal(DurationString, &o.Duration)
	}
	
	if WrapupCodes, ok := PolicyconditionsMap["wrapupCodes"].([]interface{}); ok {
		WrapupCodesString, _ := json.Marshal(WrapupCodes)
		json.Unmarshal(WrapupCodesString, &o.WrapupCodes)
	}
	
	if TimeAllowed, ok := PolicyconditionsMap["timeAllowed"].(map[string]interface{}); ok {
		TimeAllowedString, _ := json.Marshal(TimeAllowed)
		json.Unmarshal(TimeAllowedString, &o.TimeAllowed)
	}
	
	if CustomerParticipation, ok := PolicyconditionsMap["customerParticipation"].(string); ok {
		o.CustomerParticipation = &CustomerParticipation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Policyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
