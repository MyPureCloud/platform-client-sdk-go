package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignscheduleconfigchangeschedulerecurrence
type Dialercampaignscheduleconfigchangeschedulerecurrence struct { 
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
	VarRange *Dialercampaignscheduleconfigchangerecurrencerange `json:"range,omitempty"`

	// Pattern
	Pattern *Dialercampaignscheduleconfigchangerecurrencepattern `json:"pattern,omitempty"`

	// Alterations - modifications to the original recurrence schedule
	Alterations *[]Dialercampaignscheduleconfigchangealteration `json:"alterations,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercampaignscheduleconfigchangeschedulerecurrence) SetField(field string, fieldValue interface{}) {
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

func (o Dialercampaignscheduleconfigchangeschedulerecurrence) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercampaignscheduleconfigchangeschedulerecurrence
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		VarRange *Dialercampaignscheduleconfigchangerecurrencerange `json:"range,omitempty"`
		
		Pattern *Dialercampaignscheduleconfigchangerecurrencepattern `json:"pattern,omitempty"`
		
		Alterations *[]Dialercampaignscheduleconfigchangealteration `json:"alterations,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
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
		Alias:    (Alias)(o),
	})
}

func (o *Dialercampaignscheduleconfigchangeschedulerecurrence) UnmarshalJSON(b []byte) error {
	var DialercampaignscheduleconfigchangeschedulerecurrenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignscheduleconfigchangeschedulerecurrenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Start, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["end"].(string); ok {
		o.End = &End
	}
    
	if TimeZone, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if VarRange, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Pattern, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if Alterations, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["alterations"].([]interface{}); ok {
		AlterationsString, _ := json.Marshal(Alterations)
		json.Unmarshal(AlterationsString, &o.Alterations)
	}
	
	if AdditionalProperties, ok := DialercampaignscheduleconfigchangeschedulerecurrenceMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangeschedulerecurrence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
