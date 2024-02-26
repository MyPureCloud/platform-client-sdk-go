package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewelement - An element within a journey view
type Journeyviewelement struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The unique identifier of the element within the elements list
	Id *string `json:"id,omitempty"`

	// Name - The unique name of the element within the view
	Name *string `json:"name,omitempty"`

	// Attributes - Required attributes of the element
	Attributes *Journeyviewelementattributes `json:"attributes,omitempty"`

	// Filter - Any filters applied to this element
	Filter *Journeyviewelementfilter `json:"filter,omitempty"`

	// FollowedBy - A list of JourneyViewLink objects, listing the elements downstream of this element
	FollowedBy *[]Journeyviewlink `json:"followedBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewelement) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewelement) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewelement
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Attributes *Journeyviewelementattributes `json:"attributes,omitempty"`
		
		Filter *Journeyviewelementfilter `json:"filter,omitempty"`
		
		FollowedBy *[]Journeyviewlink `json:"followedBy,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Attributes: o.Attributes,
		
		Filter: o.Filter,
		
		FollowedBy: o.FollowedBy,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewelement) UnmarshalJSON(b []byte) error {
	var JourneyviewelementMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewelementMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyviewelementMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := JourneyviewelementMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Attributes, ok := JourneyviewelementMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Filter, ok := JourneyviewelementMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if FollowedBy, ok := JourneyviewelementMap["followedBy"].([]interface{}); ok {
		FollowedByString, _ := json.Marshal(FollowedBy)
		json.Unmarshal(FollowedByString, &o.FollowedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewelement) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
