package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewlink - A link between elements in a journey view
type Journeyviewlink struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The identifier of the element downstream
	Id *string `json:"id,omitempty"`

	// ConstraintWithin - A time constraint on this link, which requires a customer to complete the downstream element within this amount of time to be counted.
	ConstraintWithin *Journeyviewlinktimeconstraint `json:"constraintWithin,omitempty"`

	// ConstraintAfter - A time constraint on this link, which requires a customer must complete the downstream element after this amount of time to be counted.
	ConstraintAfter *Journeyviewlinktimeconstraint `json:"constraintAfter,omitempty"`

	// EventCountType - The type of events that will be counted. Note: Concurrent will override any JourneyViewLinkTimeConstraint. Default is Sequential.
	EventCountType *string `json:"eventCountType,omitempty"`

	// JoinAttributes - Other (secondary) attributes on which this link should join the customers being counted
	JoinAttributes *[]string `json:"joinAttributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewlink) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewlink) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewlink
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConstraintWithin *Journeyviewlinktimeconstraint `json:"constraintWithin,omitempty"`
		
		ConstraintAfter *Journeyviewlinktimeconstraint `json:"constraintAfter,omitempty"`
		
		EventCountType *string `json:"eventCountType,omitempty"`
		
		JoinAttributes *[]string `json:"joinAttributes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ConstraintWithin: o.ConstraintWithin,
		
		ConstraintAfter: o.ConstraintAfter,
		
		EventCountType: o.EventCountType,
		
		JoinAttributes: o.JoinAttributes,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewlink) UnmarshalJSON(b []byte) error {
	var JourneyviewlinkMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewlinkMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyviewlinkMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConstraintWithin, ok := JourneyviewlinkMap["constraintWithin"].(map[string]interface{}); ok {
		ConstraintWithinString, _ := json.Marshal(ConstraintWithin)
		json.Unmarshal(ConstraintWithinString, &o.ConstraintWithin)
	}
	
	if ConstraintAfter, ok := JourneyviewlinkMap["constraintAfter"].(map[string]interface{}); ok {
		ConstraintAfterString, _ := json.Marshal(ConstraintAfter)
		json.Unmarshal(ConstraintAfterString, &o.ConstraintAfter)
	}
	
	if EventCountType, ok := JourneyviewlinkMap["eventCountType"].(string); ok {
		o.EventCountType = &EventCountType
	}
    
	if JoinAttributes, ok := JourneyviewlinkMap["joinAttributes"].([]interface{}); ok {
		JoinAttributesString, _ := json.Marshal(JoinAttributes)
		json.Unmarshal(JoinAttributesString, &o.JoinAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewlink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
