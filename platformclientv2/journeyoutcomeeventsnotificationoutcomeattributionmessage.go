package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcomeattributionmessage
type Journeyoutcomeeventsnotificationoutcomeattributionmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Outcome
	Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`

	// OutcomeTouchpoints
	OutcomeTouchpoints *[]Journeyoutcomeeventsnotificationoutcometouchpoint `json:"outcomeTouchpoints,omitempty"`

	// SegmentAssignments
	SegmentAssignments *[]Journeyoutcomeeventsnotificationsegment `json:"segmentAssignments,omitempty"`

	// AssociatedValue
	AssociatedValue *Journeyoutcomeeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeyoutcomeeventsnotificationoutcomeattributionmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyoutcomeeventsnotificationoutcomeattributionmessage
	
	return json.Marshal(&struct { 
		Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`
		
		OutcomeTouchpoints *[]Journeyoutcomeeventsnotificationoutcometouchpoint `json:"outcomeTouchpoints,omitempty"`
		
		SegmentAssignments *[]Journeyoutcomeeventsnotificationsegment `json:"segmentAssignments,omitempty"`
		
		AssociatedValue *Journeyoutcomeeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`
		Alias
	}{ 
		Outcome: o.Outcome,
		
		OutcomeTouchpoints: o.OutcomeTouchpoints,
		
		SegmentAssignments: o.SegmentAssignments,
		
		AssociatedValue: o.AssociatedValue,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcomeattributionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcomeattributionmessageMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if OutcomeTouchpoints, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["outcomeTouchpoints"].([]interface{}); ok {
		OutcomeTouchpointsString, _ := json.Marshal(OutcomeTouchpoints)
		json.Unmarshal(OutcomeTouchpointsString, &o.OutcomeTouchpoints)
	}
	
	if SegmentAssignments, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["segmentAssignments"].([]interface{}); ok {
		SegmentAssignmentsString, _ := json.Marshal(SegmentAssignments)
		json.Unmarshal(SegmentAssignmentsString, &o.SegmentAssignments)
	}
	
	if AssociatedValue, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["associatedValue"].(map[string]interface{}); ok {
		AssociatedValueString, _ := json.Marshal(AssociatedValue)
		json.Unmarshal(AssociatedValueString, &o.AssociatedValue)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
