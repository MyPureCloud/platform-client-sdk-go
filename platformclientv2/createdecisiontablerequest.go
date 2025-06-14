package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createdecisiontablerequest
type Createdecisiontablerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The decision table name.
	Name *string `json:"name,omitempty"`

	// Description - The decision table description.
	Description *string `json:"description,omitempty"`

	// DivisionId - The ID of the division the decision table belongs to.
	DivisionId *string `json:"divisionId,omitempty"`

	// SchemaId - The ID of the rules schema used by the decision table.
	SchemaId *string `json:"schemaId,omitempty"`

	// Columns - The column definitions for this decision table.
	Columns *Createdecisiontablecolumnsrequest `json:"columns,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createdecisiontablerequest) SetField(field string, fieldValue interface{}) {
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

func (o Createdecisiontablerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createdecisiontablerequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		SchemaId *string `json:"schemaId,omitempty"`
		
		Columns *Createdecisiontablecolumnsrequest `json:"columns,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		DivisionId: o.DivisionId,
		
		SchemaId: o.SchemaId,
		
		Columns: o.Columns,
		Alias:    (Alias)(o),
	})
}

func (o *Createdecisiontablerequest) UnmarshalJSON(b []byte) error {
	var CreatedecisiontablerequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatedecisiontablerequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatedecisiontablerequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CreatedecisiontablerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DivisionId, ok := CreatedecisiontablerequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if SchemaId, ok := CreatedecisiontablerequestMap["schemaId"].(string); ok {
		o.SchemaId = &SchemaId
	}
    
	if Columns, ok := CreatedecisiontablerequestMap["columns"].(map[string]interface{}); ok {
		ColumnsString, _ := json.Marshal(Columns)
		json.Unmarshal(ColumnsString, &o.Columns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createdecisiontablerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
