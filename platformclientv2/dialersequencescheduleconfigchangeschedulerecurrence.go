package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequencescheduleconfigchangeschedulerecurrence
type Dialersequencescheduleconfigchangeschedulerecurrence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - the recurrence id
	Id *string `json:"id,omitempty"`

	// Start - scheduled start time represented as an ISO-8601 string; for example, yyyy-MM-ddTHH:mm:ss.SSSZ
	Start *string `json:"start,omitempty"`

	// End - scheduled end time represented as an ISO-8601 string; for example, yyyy-MM-ddTHH:mm:ss.SSSZ
	End *string `json:"end,omitempty"`

	// TimeZone - the timezone the recurrence will use
	TimeZone *string `json:"timeZone,omitempty"`

	// VarRange
	VarRange *Dialersequencescheduleconfigchangerecurrencerange `json:"range,omitempty"`

	// Pattern
	Pattern *Dialersequencescheduleconfigchangerecurrencepattern `json:"pattern,omitempty"`

	// Alterations - modifications to the original recurrence schedule
	Alterations *[]Dialersequencescheduleconfigchangealteration `json:"alterations,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// GetAdditionalProperties
	GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialersequencescheduleconfigchangeschedulerecurrence) SetField(field string, fieldValue interface{}) {
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

func (o Dialersequencescheduleconfigchangeschedulerecurrence) MarshalJSON() ([]byte, error) {
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
	type Alias Dialersequencescheduleconfigchangeschedulerecurrence
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		VarRange *Dialersequencescheduleconfigchangerecurrencerange `json:"range,omitempty"`
		
		Pattern *Dialersequencescheduleconfigchangerecurrencepattern `json:"pattern,omitempty"`
		
		Alterations *[]Dialersequencescheduleconfigchangealteration `json:"alterations,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Start: o.Start,
		
		End: o.End,
		
		TimeZone: o.TimeZone,
		
		VarRange: o.VarRange,
		
		Pattern: o.Pattern,
		
		Alterations: o.Alterations,
		
		AdditionalProperties: o.AdditionalProperties,
		
		GetAdditionalProperties: o.GetAdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialersequencescheduleconfigchangeschedulerecurrence) UnmarshalJSON(b []byte) error {
	var DialersequencescheduleconfigchangeschedulerecurrenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequencescheduleconfigchangeschedulerecurrenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Start, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["end"].(string); ok {
		o.End = &End
	}
    
	if TimeZone, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if VarRange, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Pattern, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if Alterations, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["alterations"].([]interface{}); ok {
		AlterationsString, _ := json.Marshal(Alterations)
		json.Unmarshal(AlterationsString, &o.Alterations)
	}
	
	if AdditionalProperties, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if GetAdditionalProperties, ok := DialersequencescheduleconfigchangeschedulerecurrenceMap["getAdditionalProperties"].(map[string]interface{}); ok {
		GetAdditionalPropertiesString, _ := json.Marshal(GetAdditionalProperties)
		json.Unmarshal(GetAdditionalPropertiesString, &o.GetAdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangeschedulerecurrence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
