package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2enterpriseuser - Defines a SCIM enterprise user.
type Scimv2enterpriseuser struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Division - The division that the user belongs to.
	Division *string `json:"division,omitempty"`

	// Department - The department that the user belongs to.
	Department *string `json:"department,omitempty"`

	// Manager - The user's manager.
	Manager *Manager `json:"manager,omitempty"`

	// EmployeeNumber - The user's employee number.
	EmployeeNumber *string `json:"employeeNumber,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimv2enterpriseuser) SetField(field string, fieldValue interface{}) {
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

func (o Scimv2enterpriseuser) MarshalJSON() ([]byte, error) {
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
	type Alias Scimv2enterpriseuser
	
	return json.Marshal(&struct { 
		Division *string `json:"division,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Manager *Manager `json:"manager,omitempty"`
		
		EmployeeNumber *string `json:"employeeNumber,omitempty"`
		Alias
	}{ 
		Division: o.Division,
		
		Department: o.Department,
		
		Manager: o.Manager,
		
		EmployeeNumber: o.EmployeeNumber,
		Alias:    (Alias)(o),
	})
}

func (o *Scimv2enterpriseuser) UnmarshalJSON(b []byte) error {
	var Scimv2enterpriseuserMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2enterpriseuserMap)
	if err != nil {
		return err
	}
	
	if Division, ok := Scimv2enterpriseuserMap["division"].(string); ok {
		o.Division = &Division
	}
    
	if Department, ok := Scimv2enterpriseuserMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Manager, ok := Scimv2enterpriseuserMap["manager"].(map[string]interface{}); ok {
		ManagerString, _ := json.Marshal(Manager)
		json.Unmarshal(ManagerString, &o.Manager)
	}
	
	if EmployeeNumber, ok := Scimv2enterpriseuserMap["employeeNumber"].(string); ok {
		o.EmployeeNumber = &EmployeeNumber
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2enterpriseuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
