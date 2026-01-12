package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Segmentassignment
type Segmentassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateAssigned - Date when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssigned *time.Time `json:"dateAssigned,omitempty"`

	// DateForUnassignment - Date indicating when a segment is scheduled to be unassigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateForUnassignment *time.Time `json:"dateForUnassignment,omitempty"`

	// Segment - The segment the assignment is for.
	Segment *Segmentassignmentsegment `json:"segment,omitempty"`

	// ExternalContact - External contact of the customer to which the segment is assigned.
	ExternalContact *Addressableentityref `json:"externalContact,omitempty"`

	// Session - For session-scoped segments, the session for which the segment was assigned.
	Session *Segmentassignmentsession `json:"session,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Segmentassignment) SetField(field string, fieldValue interface{}) {
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

func (o Segmentassignment) MarshalJSON() ([]byte, error) {
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
	type Alias Segmentassignment
	
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
		DateAssigned *string `json:"dateAssigned,omitempty"`
		
		DateForUnassignment *string `json:"dateForUnassignment,omitempty"`
		
		Segment *Segmentassignmentsegment `json:"segment,omitempty"`
		
		ExternalContact *Addressableentityref `json:"externalContact,omitempty"`
		
		Session *Segmentassignmentsession `json:"session,omitempty"`
		Alias
	}{ 
		DateAssigned: DateAssigned,
		
		DateForUnassignment: DateForUnassignment,
		
		Segment: o.Segment,
		
		ExternalContact: o.ExternalContact,
		
		Session: o.Session,
		Alias:    (Alias)(o),
	})
}

func (o *Segmentassignment) UnmarshalJSON(b []byte) error {
	var SegmentassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &SegmentassignmentMap)
	if err != nil {
		return err
	}
	
	if dateAssignedString, ok := SegmentassignmentMap["dateAssigned"].(string); ok {
		DateAssigned, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssignedString)
		o.DateAssigned = &DateAssigned
	}
	
	if dateForUnassignmentString, ok := SegmentassignmentMap["dateForUnassignment"].(string); ok {
		DateForUnassignment, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateForUnassignmentString)
		o.DateForUnassignment = &DateForUnassignment
	}
	
	if Segment, ok := SegmentassignmentMap["segment"].(map[string]interface{}); ok {
		SegmentString, _ := json.Marshal(Segment)
		json.Unmarshal(SegmentString, &o.Segment)
	}
	
	if ExternalContact, ok := SegmentassignmentMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if Session, ok := SegmentassignmentMap["session"].(map[string]interface{}); ok {
		SessionString, _ := json.Marshal(Session)
		json.Unmarshal(SessionString, &o.Session)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Segmentassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
