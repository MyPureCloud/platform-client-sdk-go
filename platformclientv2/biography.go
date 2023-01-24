package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Biography
type Biography struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Biography - Personal detailed description
	Biography *string `json:"biography,omitempty"`

	// Interests
	Interests *[]string `json:"interests,omitempty"`

	// Hobbies
	Hobbies *[]string `json:"hobbies,omitempty"`

	// Spouse
	Spouse *string `json:"spouse,omitempty"`

	// Education - User education details
	Education *[]Education `json:"education,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Biography) SetField(field string, fieldValue interface{}) {
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

func (o Biography) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Biography
	
	return json.Marshal(&struct { 
		Biography *string `json:"biography,omitempty"`
		
		Interests *[]string `json:"interests,omitempty"`
		
		Hobbies *[]string `json:"hobbies,omitempty"`
		
		Spouse *string `json:"spouse,omitempty"`
		
		Education *[]Education `json:"education,omitempty"`
		Alias
	}{ 
		Biography: o.Biography,
		
		Interests: o.Interests,
		
		Hobbies: o.Hobbies,
		
		Spouse: o.Spouse,
		
		Education: o.Education,
		Alias:    (Alias)(o),
	})
}

func (o *Biography) UnmarshalJSON(b []byte) error {
	var BiographyMap map[string]interface{}
	err := json.Unmarshal(b, &BiographyMap)
	if err != nil {
		return err
	}
	
	if Biography, ok := BiographyMap["biography"].(string); ok {
		o.Biography = &Biography
	}
    
	if Interests, ok := BiographyMap["interests"].([]interface{}); ok {
		InterestsString, _ := json.Marshal(Interests)
		json.Unmarshal(InterestsString, &o.Interests)
	}
	
	if Hobbies, ok := BiographyMap["hobbies"].([]interface{}); ok {
		HobbiesString, _ := json.Marshal(Hobbies)
		json.Unmarshal(HobbiesString, &o.Hobbies)
	}
	
	if Spouse, ok := BiographyMap["spouse"].(string); ok {
		o.Spouse = &Spouse
	}
    
	if Education, ok := BiographyMap["education"].([]interface{}); ok {
		EducationString, _ := json.Marshal(Education)
		json.Unmarshal(EducationString, &o.Education)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Biography) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
