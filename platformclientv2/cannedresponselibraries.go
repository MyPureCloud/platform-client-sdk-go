package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Cannedresponselibraries
type Cannedresponselibraries struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// LibraryIds - Set of canned response library IDs associated with the queue only when mode is SelectedOnly.
	LibraryIds *[]string `json:"libraryIds,omitempty"`

	// Mode - The association mode of canned response libraries to queue
	Mode *string `json:"mode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Cannedresponselibraries) SetField(field string, fieldValue interface{}) {
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

func (o Cannedresponselibraries) MarshalJSON() ([]byte, error) {
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
	type Alias Cannedresponselibraries
	
	return json.Marshal(&struct { 
		LibraryIds *[]string `json:"libraryIds,omitempty"`
		
		Mode *string `json:"mode,omitempty"`
		Alias
	}{ 
		LibraryIds: o.LibraryIds,
		
		Mode: o.Mode,
		Alias:    (Alias)(o),
	})
}

func (o *Cannedresponselibraries) UnmarshalJSON(b []byte) error {
	var CannedresponselibrariesMap map[string]interface{}
	err := json.Unmarshal(b, &CannedresponselibrariesMap)
	if err != nil {
		return err
	}
	
	if LibraryIds, ok := CannedresponselibrariesMap["libraryIds"].([]interface{}); ok {
		LibraryIdsString, _ := json.Marshal(LibraryIds)
		json.Unmarshal(LibraryIdsString, &o.LibraryIds)
	}
	
	if Mode, ok := CannedresponselibrariesMap["mode"].(string); ok {
		o.Mode = &Mode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Cannedresponselibraries) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
