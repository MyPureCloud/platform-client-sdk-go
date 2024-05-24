package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reoccurrence
type Reoccurrence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Start - The  start date time of the initial occurrence as an ISO-8601 string in UTC time, e.g: 2023-11-21T16:30:25.000Z
	Start *string `json:"start,omitempty"`

	// End - The end date time of the initial occurrence as an ISO-8601 string in UTC time, e.g: 2023-12-21T16:30:25.000Z
	End *string `json:"end,omitempty"`

	// TimeZone - The time zone of the schedule e.g.:  America/New_York
	TimeZone *string `json:"timeZone,omitempty"`

	// Pattern - The schedule pattern e.g.: Daily/Weekly
	Pattern *Pattern `json:"pattern,omitempty"`

	// VarRange - The schedule range e.g.: EndDate/NoEnd/Numbered
	VarRange *Range `json:"range,omitempty"`

	// Alterations - Modifications to the original recurrence schedule (Exclusions/Inclusions)
	Alterations *[]Alteration `json:"alterations,omitempty"`

	// NextOccurrenceDetails - The next occurrence details for the next start and end occurrences for the recurrence
	NextOccurrenceDetails *Nextoccurrencedetails `json:"nextOccurrenceDetails,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reoccurrence) SetField(field string, fieldValue interface{}) {
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

func (o Reoccurrence) MarshalJSON() ([]byte, error) {
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
	type Alias Reoccurrence
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Pattern *Pattern `json:"pattern,omitempty"`
		
		VarRange *Range `json:"range,omitempty"`
		
		Alterations *[]Alteration `json:"alterations,omitempty"`
		
		NextOccurrenceDetails *Nextoccurrencedetails `json:"nextOccurrenceDetails,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Start: o.Start,
		
		End: o.End,
		
		TimeZone: o.TimeZone,
		
		Pattern: o.Pattern,
		
		VarRange: o.VarRange,
		
		Alterations: o.Alterations,
		
		NextOccurrenceDetails: o.NextOccurrenceDetails,
		Alias:    (Alias)(o),
	})
}

func (o *Reoccurrence) UnmarshalJSON(b []byte) error {
	var ReoccurrenceMap map[string]interface{}
	err := json.Unmarshal(b, &ReoccurrenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReoccurrenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Start, ok := ReoccurrenceMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := ReoccurrenceMap["end"].(string); ok {
		o.End = &End
	}
    
	if TimeZone, ok := ReoccurrenceMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if Pattern, ok := ReoccurrenceMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if VarRange, ok := ReoccurrenceMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Alterations, ok := ReoccurrenceMap["alterations"].([]interface{}); ok {
		AlterationsString, _ := json.Marshal(Alterations)
		json.Unmarshal(AlterationsString, &o.Alterations)
	}
	
	if NextOccurrenceDetails, ok := ReoccurrenceMap["nextOccurrenceDetails"].(map[string]interface{}); ok {
		NextOccurrenceDetailsString, _ := json.Marshal(NextOccurrenceDetails)
		json.Unmarshal(NextOccurrenceDetailsString, &o.NextOccurrenceDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reoccurrence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
