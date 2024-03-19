package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewresultmetrics - The metrics of an element or a link in journey 
type Journeyviewresultmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ParticipantCount - number of external contacts who participated at the element
	ParticipantCount *int `json:"participantCount,omitempty"`

	// ActiveCount - number of external contacts who could still complete the journey
	ActiveCount *int `json:"activeCount,omitempty"`

	// CompletedCount - number of external contacts who completed the journey forward
	CompletedCount *int `json:"completedCount,omitempty"`

	// DropoutCount - number of external contacts who did not make it to the next element
	DropoutCount *int `json:"dropoutCount,omitempty"`

	// FlowCount - number of external contacts who moved from one element to next element but did not complete the journey
	FlowCount *int `json:"flowCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewresultmetrics) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewresultmetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewresultmetrics
	
	return json.Marshal(&struct { 
		ParticipantCount *int `json:"participantCount,omitempty"`
		
		ActiveCount *int `json:"activeCount,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		DropoutCount *int `json:"dropoutCount,omitempty"`
		
		FlowCount *int `json:"flowCount,omitempty"`
		Alias
	}{ 
		ParticipantCount: o.ParticipantCount,
		
		ActiveCount: o.ActiveCount,
		
		CompletedCount: o.CompletedCount,
		
		DropoutCount: o.DropoutCount,
		
		FlowCount: o.FlowCount,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewresultmetrics) UnmarshalJSON(b []byte) error {
	var JourneyviewresultmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewresultmetricsMap)
	if err != nil {
		return err
	}
	
	if ParticipantCount, ok := JourneyviewresultmetricsMap["participantCount"].(float64); ok {
		ParticipantCountInt := int(ParticipantCount)
		o.ParticipantCount = &ParticipantCountInt
	}
	
	if ActiveCount, ok := JourneyviewresultmetricsMap["activeCount"].(float64); ok {
		ActiveCountInt := int(ActiveCount)
		o.ActiveCount = &ActiveCountInt
	}
	
	if CompletedCount, ok := JourneyviewresultmetricsMap["completedCount"].(float64); ok {
		CompletedCountInt := int(CompletedCount)
		o.CompletedCount = &CompletedCountInt
	}
	
	if DropoutCount, ok := JourneyviewresultmetricsMap["dropoutCount"].(float64); ok {
		DropoutCountInt := int(DropoutCount)
		o.DropoutCount = &DropoutCountInt
	}
	
	if FlowCount, ok := JourneyviewresultmetricsMap["flowCount"].(float64); ok {
		FlowCountInt := int(FlowCount)
		o.FlowCount = &FlowCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewresultmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
