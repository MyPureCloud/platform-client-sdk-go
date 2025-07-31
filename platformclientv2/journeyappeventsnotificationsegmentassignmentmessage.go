package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyappeventsnotificationsegmentassignmentmessage
type Journeyappeventsnotificationsegmentassignmentmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Segment
	Segment *Journeyappeventsnotificationsegment `json:"segment,omitempty"`

	// AssignmentState
	AssignmentState *string `json:"assignmentState,omitempty"`

	// DateAssigned
	DateAssigned *time.Time `json:"dateAssigned,omitempty"`

	// DateForUnassignment
	DateForUnassignment *time.Time `json:"dateForUnassignment,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyappeventsnotificationsegmentassignmentmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeyappeventsnotificationsegmentassignmentmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAssigned","DateForUnassignment", }
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
	type Alias Journeyappeventsnotificationsegmentassignmentmessage
	
	DateAssigned := new(string)
	if o.DateAssigned != nil {
		
		*DateAssigned = timeutil.Strftime(o.DateAssigned, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssigned = nil
	}
	
	DateForUnassignment := new(string)
	if o.DateForUnassignment != nil {
		
		*DateForUnassignment = timeutil.Strftime(o.DateForUnassignment, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateForUnassignment = nil
	}
	
	return json.Marshal(&struct { 
		Segment *Journeyappeventsnotificationsegment `json:"segment,omitempty"`
		
		AssignmentState *string `json:"assignmentState,omitempty"`
		
		DateAssigned *string `json:"dateAssigned,omitempty"`
		
		DateForUnassignment *string `json:"dateForUnassignment,omitempty"`
		Alias
	}{ 
		Segment: o.Segment,
		
		AssignmentState: o.AssignmentState,
		
		DateAssigned: DateAssigned,
		
		DateForUnassignment: DateForUnassignment,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyappeventsnotificationsegmentassignmentmessage) UnmarshalJSON(b []byte) error {
	var JourneyappeventsnotificationsegmentassignmentmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappeventsnotificationsegmentassignmentmessageMap)
	if err != nil {
		return err
	}
	
	if Segment, ok := JourneyappeventsnotificationsegmentassignmentmessageMap["segment"].(map[string]interface{}); ok {
		SegmentString, _ := json.Marshal(Segment)
		json.Unmarshal(SegmentString, &o.Segment)
	}
	
	if AssignmentState, ok := JourneyappeventsnotificationsegmentassignmentmessageMap["assignmentState"].(string); ok {
		o.AssignmentState = &AssignmentState
	}
    
	if dateAssignedString, ok := JourneyappeventsnotificationsegmentassignmentmessageMap["dateAssigned"].(string); ok {
		DateAssigned, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssignedString)
		o.DateAssigned = &DateAssigned
	}
	
	if dateForUnassignmentString, ok := JourneyappeventsnotificationsegmentassignmentmessageMap["dateForUnassignment"].(string); ok {
		DateForUnassignment, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateForUnassignmentString)
		o.DateForUnassignment = &DateForUnassignment
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyappeventsnotificationsegmentassignmentmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
