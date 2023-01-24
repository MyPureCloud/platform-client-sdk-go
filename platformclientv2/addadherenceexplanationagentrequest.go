package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Addadherenceexplanationagentrequest
type Addadherenceexplanationagentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the adherence explanation
	VarType *string `json:"type,omitempty"`

	// StartDate - The start timestamp of the adherence explanation in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Notes - Notes about the adherence explanation
	Notes *string `json:"notes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Addadherenceexplanationagentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Addadherenceexplanationagentrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate", }
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
	type Alias Addadherenceexplanationagentrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Notes: o.Notes,
		Alias:    (Alias)(o),
	})
}

func (o *Addadherenceexplanationagentrequest) UnmarshalJSON(b []byte) error {
	var AddadherenceexplanationagentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddadherenceexplanationagentrequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AddadherenceexplanationagentrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if startDateString, ok := AddadherenceexplanationagentrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AddadherenceexplanationagentrequestMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Notes, ok := AddadherenceexplanationagentrequestMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Addadherenceexplanationagentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
