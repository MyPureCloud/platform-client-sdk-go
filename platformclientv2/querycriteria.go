package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Querycriteria - A criteria type that can be used in tandem with other criteria type to create queries of executionData
type Querycriteria struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CriteriaKey - The is the name of the criteria that can be queried.
	CriteriaKey *string `json:"criteriaKey,omitempty"`

	// CriteriaGroups - The executionData type that this criteria item can be used on.
	CriteriaGroups *[]string `json:"criteriaGroups,omitempty"`

	// Description - The is the description of the criteria.
	Description *string `json:"description,omitempty"`

	// Operators - A list of operators that can be used on this criteria.
	Operators *[]string `json:"operators,omitempty"`

	// DataType - The type of data for the criteria (string, int, etc).
	DataType *string `json:"dataType,omitempty"`

	// CategoryInfo - A logical grouping and display order for this item.
	CategoryInfo *Criteriacategoryinfo `json:"categoryInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Querycriteria) SetField(field string, fieldValue interface{}) {
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

func (o Querycriteria) MarshalJSON() ([]byte, error) {
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
	type Alias Querycriteria
	
	return json.Marshal(&struct { 
		CriteriaKey *string `json:"criteriaKey,omitempty"`
		
		CriteriaGroups *[]string `json:"criteriaGroups,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Operators *[]string `json:"operators,omitempty"`
		
		DataType *string `json:"dataType,omitempty"`
		
		CategoryInfo *Criteriacategoryinfo `json:"categoryInfo,omitempty"`
		Alias
	}{ 
		CriteriaKey: o.CriteriaKey,
		
		CriteriaGroups: o.CriteriaGroups,
		
		Description: o.Description,
		
		Operators: o.Operators,
		
		DataType: o.DataType,
		
		CategoryInfo: o.CategoryInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Querycriteria) UnmarshalJSON(b []byte) error {
	var QuerycriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &QuerycriteriaMap)
	if err != nil {
		return err
	}
	
	if CriteriaKey, ok := QuerycriteriaMap["criteriaKey"].(string); ok {
		o.CriteriaKey = &CriteriaKey
	}
    
	if CriteriaGroups, ok := QuerycriteriaMap["criteriaGroups"].([]interface{}); ok {
		CriteriaGroupsString, _ := json.Marshal(CriteriaGroups)
		json.Unmarshal(CriteriaGroupsString, &o.CriteriaGroups)
	}
	
	if Description, ok := QuerycriteriaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Operators, ok := QuerycriteriaMap["operators"].([]interface{}); ok {
		OperatorsString, _ := json.Marshal(Operators)
		json.Unmarshal(OperatorsString, &o.Operators)
	}
	
	if DataType, ok := QuerycriteriaMap["dataType"].(string); ok {
		o.DataType = &DataType
	}
    
	if CategoryInfo, ok := QuerycriteriaMap["categoryInfo"].(map[string]interface{}); ok {
		CategoryInfoString, _ := json.Marshal(CategoryInfo)
		json.Unmarshal(CategoryInfoString, &o.CategoryInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querycriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
