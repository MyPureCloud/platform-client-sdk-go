package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffsettingsrequest
type Timeoffsettingsrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
	SubmissionRangeEnforced *bool `json:"submissionRangeEnforced,omitempty"`

	// SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
	SubmissionEarliestDaysFromNow *int `json:"submissionEarliestDaysFromNow,omitempty"`

	// SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
	SubmissionLatestDaysFromNow *int `json:"submissionLatestDaysFromNow,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffsettingsrequest) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffsettingsrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Timeoffsettingsrequest
	
	return json.Marshal(&struct { 
		SubmissionRangeEnforced *bool `json:"submissionRangeEnforced,omitempty"`
		
		SubmissionEarliestDaysFromNow *int `json:"submissionEarliestDaysFromNow,omitempty"`
		
		SubmissionLatestDaysFromNow *int `json:"submissionLatestDaysFromNow,omitempty"`
		Alias
	}{ 
		SubmissionRangeEnforced: o.SubmissionRangeEnforced,
		
		SubmissionEarliestDaysFromNow: o.SubmissionEarliestDaysFromNow,
		
		SubmissionLatestDaysFromNow: o.SubmissionLatestDaysFromNow,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffsettingsrequest) UnmarshalJSON(b []byte) error {
	var TimeoffsettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffsettingsrequestMap)
	if err != nil {
		return err
	}
	
	if SubmissionRangeEnforced, ok := TimeoffsettingsrequestMap["submissionRangeEnforced"].(bool); ok {
		o.SubmissionRangeEnforced = &SubmissionRangeEnforced
	}
    
	if SubmissionEarliestDaysFromNow, ok := TimeoffsettingsrequestMap["submissionEarliestDaysFromNow"].(float64); ok {
		SubmissionEarliestDaysFromNowInt := int(SubmissionEarliestDaysFromNow)
		o.SubmissionEarliestDaysFromNow = &SubmissionEarliestDaysFromNowInt
	}
	
	if SubmissionLatestDaysFromNow, ok := TimeoffsettingsrequestMap["submissionLatestDaysFromNow"].(float64); ok {
		SubmissionLatestDaysFromNowInt := int(SubmissionLatestDaysFromNow)
		o.SubmissionLatestDaysFromNow = &SubmissionLatestDaysFromNowInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
