package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Addadherenceadjustmentagentrequest
type Addadherenceadjustmentagentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ReasonCodeId - The ID of the reason code for this adherence adjustment
	ReasonCodeId *string `json:"reasonCodeId,omitempty"`

	// StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of the adherence adjustment in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// SubmitterNotes - Notes provided by the submitter for this adherence adjustment
	SubmitterNotes *string `json:"submitterNotes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Addadherenceadjustmentagentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Addadherenceadjustmentagentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Addadherenceadjustmentagentrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ReasonCodeId *string `json:"reasonCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		SubmitterNotes *string `json:"submitterNotes,omitempty"`
		Alias
	}{ 
		ReasonCodeId: o.ReasonCodeId,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		SubmitterNotes: o.SubmitterNotes,
		Alias:    (Alias)(o),
	})
}

func (o *Addadherenceadjustmentagentrequest) UnmarshalJSON(b []byte) error {
	var AddadherenceadjustmentagentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddadherenceadjustmentagentrequestMap)
	if err != nil {
		return err
	}
	
	if ReasonCodeId, ok := AddadherenceadjustmentagentrequestMap["reasonCodeId"].(string); ok {
		o.ReasonCodeId = &ReasonCodeId
	}
    
	if startDateString, ok := AddadherenceadjustmentagentrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AddadherenceadjustmentagentrequestMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if SubmitterNotes, ok := AddadherenceadjustmentagentrequestMap["submitterNotes"].(string); ok {
		o.SubmitterNotes = &SubmitterNotes
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Addadherenceadjustmentagentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
