package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewelementfilterrange - the range of attribute values to filter on. At least one comparator must be defined
type Journeyviewelementfilterrange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Lt - comparator: less than
	Lt *Journeyviewelementfilterrangedata `json:"lt,omitempty"`

	// Lte - comparator: less than or equal
	Lte *Journeyviewelementfilterrangedata `json:"lte,omitempty"`

	// Gt - comparator: greater than
	Gt *Journeyviewelementfilterrangedata `json:"gt,omitempty"`

	// Gte - comparator: greater than or equal
	Gte *Journeyviewelementfilterrangedata `json:"gte,omitempty"`

	// Eq - comparator: is equal to
	Eq *Journeyviewelementfilterrangedata `json:"eq,omitempty"`

	// Neq - comparator: is not equal to
	Neq *Journeyviewelementfilterrangedata `json:"neq,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewelementfilterrange) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewelementfilterrange) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewelementfilterrange
	
	return json.Marshal(&struct { 
		Lt *Journeyviewelementfilterrangedata `json:"lt,omitempty"`
		
		Lte *Journeyviewelementfilterrangedata `json:"lte,omitempty"`
		
		Gt *Journeyviewelementfilterrangedata `json:"gt,omitempty"`
		
		Gte *Journeyviewelementfilterrangedata `json:"gte,omitempty"`
		
		Eq *Journeyviewelementfilterrangedata `json:"eq,omitempty"`
		
		Neq *Journeyviewelementfilterrangedata `json:"neq,omitempty"`
		Alias
	}{ 
		Lt: o.Lt,
		
		Lte: o.Lte,
		
		Gt: o.Gt,
		
		Gte: o.Gte,
		
		Eq: o.Eq,
		
		Neq: o.Neq,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewelementfilterrange) UnmarshalJSON(b []byte) error {
	var JourneyviewelementfilterrangeMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewelementfilterrangeMap)
	if err != nil {
		return err
	}
	
	if Lt, ok := JourneyviewelementfilterrangeMap["lt"].(map[string]interface{}); ok {
		LtString, _ := json.Marshal(Lt)
		json.Unmarshal(LtString, &o.Lt)
	}
	
	if Lte, ok := JourneyviewelementfilterrangeMap["lte"].(map[string]interface{}); ok {
		LteString, _ := json.Marshal(Lte)
		json.Unmarshal(LteString, &o.Lte)
	}
	
	if Gt, ok := JourneyviewelementfilterrangeMap["gt"].(map[string]interface{}); ok {
		GtString, _ := json.Marshal(Gt)
		json.Unmarshal(GtString, &o.Gt)
	}
	
	if Gte, ok := JourneyviewelementfilterrangeMap["gte"].(map[string]interface{}); ok {
		GteString, _ := json.Marshal(Gte)
		json.Unmarshal(GteString, &o.Gte)
	}
	
	if Eq, ok := JourneyviewelementfilterrangeMap["eq"].(map[string]interface{}); ok {
		EqString, _ := json.Marshal(Eq)
		json.Unmarshal(EqString, &o.Eq)
	}
	
	if Neq, ok := JourneyviewelementfilterrangeMap["neq"].(map[string]interface{}); ok {
		NeqString, _ := json.Marshal(Neq)
		json.Unmarshal(NeqString, &o.Neq)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilterrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
