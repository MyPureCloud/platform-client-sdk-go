package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Swaprowindexrequest
type Swaprowindexrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SourceRowId - Unique identifier of the source row to swap
	SourceRowId *string `json:"sourceRowId,omitempty"`

	// SourceRowIndex - The current index position of the source row. Must be positive, starting from 1 and less than or equal to the size of the table
	SourceRowIndex *int `json:"sourceRowIndex,omitempty"`

	// TargetRowId - Unique identifier of the target row to swap
	TargetRowId *string `json:"targetRowId,omitempty"`

	// TargetRowIndex - The current index position of the target row. Must be positive, starting from 1 and less than or equal to the size of the table
	TargetRowIndex *int `json:"targetRowIndex,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Swaprowindexrequest) SetField(field string, fieldValue interface{}) {
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

func (o Swaprowindexrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Swaprowindexrequest
	
	return json.Marshal(&struct { 
		SourceRowId *string `json:"sourceRowId,omitempty"`
		
		SourceRowIndex *int `json:"sourceRowIndex,omitempty"`
		
		TargetRowId *string `json:"targetRowId,omitempty"`
		
		TargetRowIndex *int `json:"targetRowIndex,omitempty"`
		Alias
	}{ 
		SourceRowId: o.SourceRowId,
		
		SourceRowIndex: o.SourceRowIndex,
		
		TargetRowId: o.TargetRowId,
		
		TargetRowIndex: o.TargetRowIndex,
		Alias:    (Alias)(o),
	})
}

func (o *Swaprowindexrequest) UnmarshalJSON(b []byte) error {
	var SwaprowindexrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SwaprowindexrequestMap)
	if err != nil {
		return err
	}
	
	if SourceRowId, ok := SwaprowindexrequestMap["sourceRowId"].(string); ok {
		o.SourceRowId = &SourceRowId
	}
    
	if SourceRowIndex, ok := SwaprowindexrequestMap["sourceRowIndex"].(float64); ok {
		SourceRowIndexInt := int(SourceRowIndex)
		o.SourceRowIndex = &SourceRowIndexInt
	}
	
	if TargetRowId, ok := SwaprowindexrequestMap["targetRowId"].(string); ok {
		o.TargetRowId = &TargetRowId
	}
    
	if TargetRowIndex, ok := SwaprowindexrequestMap["targetRowIndex"].(float64); ok {
		TargetRowIndexInt := int(TargetRowIndex)
		o.TargetRowIndex = &TargetRowIndexInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Swaprowindexrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
