package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexv2slot
type Lexv2slot struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SlotName - The slot name
	SlotName *string `json:"slotName,omitempty"`

	// Description - The slot description
	Description *string `json:"description,omitempty"`

	// SlotId - The slot id
	SlotId *string `json:"slotId,omitempty"`

	// VarType - The slot type
	VarType *string `json:"type,omitempty"`

	// SlotTypeId - The slot type id
	SlotTypeId *string `json:"slotTypeId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Lexv2slot) SetField(field string, fieldValue interface{}) {
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

func (o Lexv2slot) MarshalJSON() ([]byte, error) {
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
	type Alias Lexv2slot
	
	return json.Marshal(&struct { 
		SlotName *string `json:"slotName,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SlotId *string `json:"slotId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		SlotTypeId *string `json:"slotTypeId,omitempty"`
		Alias
	}{ 
		SlotName: o.SlotName,
		
		Description: o.Description,
		
		SlotId: o.SlotId,
		
		VarType: o.VarType,
		
		SlotTypeId: o.SlotTypeId,
		Alias:    (Alias)(o),
	})
}

func (o *Lexv2slot) UnmarshalJSON(b []byte) error {
	var Lexv2slotMap map[string]interface{}
	err := json.Unmarshal(b, &Lexv2slotMap)
	if err != nil {
		return err
	}
	
	if SlotName, ok := Lexv2slotMap["slotName"].(string); ok {
		o.SlotName = &SlotName
	}
    
	if Description, ok := Lexv2slotMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SlotId, ok := Lexv2slotMap["slotId"].(string); ok {
		o.SlotId = &SlotId
	}
    
	if VarType, ok := Lexv2slotMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if SlotTypeId, ok := Lexv2slotMap["slotTypeId"].(string); ok {
		o.SlotTypeId = &SlotTypeId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lexv2slot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}