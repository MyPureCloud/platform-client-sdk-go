package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningscheduleslotsqueryrequest
type Learningscheduleslotsqueryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Range of time to get slots for scheduling learning activities. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// LengthInMinutes - The duration of Learning Assignment to schedule in 15 minutes granularity
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// UserIds - The user IDs for which to fetch schedules. Must be only 1.
	UserIds *[]string `json:"userIds,omitempty"`

	// InterruptibleAssignmentId - Assignment ID to exclude from consideration when determining blocked slots
	InterruptibleAssignmentId *string `json:"interruptibleAssignmentId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningscheduleslotsqueryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Learningscheduleslotsqueryrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Learningscheduleslotsqueryrequest
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		InterruptibleAssignmentId *string `json:"interruptibleAssignmentId,omitempty"`
		Alias
	}{ 
		Interval: o.Interval,
		
		LengthInMinutes: o.LengthInMinutes,
		
		UserIds: o.UserIds,
		
		InterruptibleAssignmentId: o.InterruptibleAssignmentId,
		Alias:    (Alias)(o),
	})
}

func (o *Learningscheduleslotsqueryrequest) UnmarshalJSON(b []byte) error {
	var LearningscheduleslotsqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningscheduleslotsqueryrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := LearningscheduleslotsqueryrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if LengthInMinutes, ok := LearningscheduleslotsqueryrequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if UserIds, ok := LearningscheduleslotsqueryrequestMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if InterruptibleAssignmentId, ok := LearningscheduleslotsqueryrequestMap["interruptibleAssignmentId"].(string); ok {
		o.InterruptibleAssignmentId = &InterruptibleAssignmentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
