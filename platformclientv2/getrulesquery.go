package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Getrulesquery
type Getrulesquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RuleType - The rule type of the alerts the query will return
	RuleType *string `json:"ruleType,omitempty"`

	// QueryType - The type of query being performed.
	QueryType *string `json:"queryType,omitempty"`

	// EnabledType - The state of the rule the query will return.  The accepted choices are Enabled, Disabled, or All
	EnabledType *string `json:"enabledType,omitempty"`

	// PageNumber - The page number of the queried response
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - The number of entities to return of the queried response.  The max is 25
	PageSize *int `json:"pageSize,omitempty"`

	// SortBy - The field to sort responses by.  The accepted choices are Name and DateStart
	SortBy *string `json:"sortBy,omitempty"`

	// SortOrder - The order in which response will be sorted.  The accepted choices are Asc and Desc
	SortOrder *string `json:"sortOrder,omitempty"`

	// RuleName - The name of the rule being queries.
	RuleName *string `json:"ruleName,omitempty"`

	// NameSearchType - Specifies how strict the name search needs to be. Expected values are Exact and Contains if querying by name.
	NameSearchType *string `json:"nameSearchType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Getrulesquery) SetField(field string, fieldValue interface{}) {
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

func (o Getrulesquery) MarshalJSON() ([]byte, error) {
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
	type Alias Getrulesquery
	
	return json.Marshal(&struct { 
		RuleType *string `json:"ruleType,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		EnabledType *string `json:"enabledType,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		RuleName *string `json:"ruleName,omitempty"`
		
		NameSearchType *string `json:"nameSearchType,omitempty"`
		Alias
	}{ 
		RuleType: o.RuleType,
		
		QueryType: o.QueryType,
		
		EnabledType: o.EnabledType,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		SortBy: o.SortBy,
		
		SortOrder: o.SortOrder,
		
		RuleName: o.RuleName,
		
		NameSearchType: o.NameSearchType,
		Alias:    (Alias)(o),
	})
}

func (o *Getrulesquery) UnmarshalJSON(b []byte) error {
	var GetrulesqueryMap map[string]interface{}
	err := json.Unmarshal(b, &GetrulesqueryMap)
	if err != nil {
		return err
	}
	
	if RuleType, ok := GetrulesqueryMap["ruleType"].(string); ok {
		o.RuleType = &RuleType
	}
    
	if QueryType, ok := GetrulesqueryMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if EnabledType, ok := GetrulesqueryMap["enabledType"].(string); ok {
		o.EnabledType = &EnabledType
	}
    
	if PageNumber, ok := GetrulesqueryMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := GetrulesqueryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if SortBy, ok := GetrulesqueryMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if SortOrder, ok := GetrulesqueryMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if RuleName, ok := GetrulesqueryMap["ruleName"].(string); ok {
		o.RuleName = &RuleName
	}
    
	if NameSearchType, ok := GetrulesqueryMap["nameSearchType"].(string); ok {
		o.NameSearchType = &NameSearchType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Getrulesquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
