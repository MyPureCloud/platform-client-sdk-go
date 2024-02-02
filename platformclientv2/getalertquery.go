package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Getalertquery
type Getalertquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RuleType - The rule type of the alerts the query will return
	RuleType *string `json:"ruleType,omitempty"`

	// QueryType - The type of query being performed.
	QueryType *string `json:"queryType,omitempty"`

	// AlertStatus - The status of the alerts the query will return.
	AlertStatus *string `json:"alertStatus,omitempty"`

	// ViewedStatus - The view status of the alerts the query will return.
	ViewedStatus *string `json:"viewedStatus,omitempty"`

	// PageNumber - The page number of the queried response
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - The number of entities to return of the queried response.  The max is 25
	PageSize *int `json:"pageSize,omitempty"`

	// SortBy - The field to sort responses by.  The accepted choices are Name and DateStart
	SortBy *string `json:"sortBy,omitempty"`

	// SortOrder - The order in which response will be sorted.  The accepted choices are Asc and Desc
	SortOrder *string `json:"sortOrder,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Getalertquery) SetField(field string, fieldValue interface{}) {
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

func (o Getalertquery) MarshalJSON() ([]byte, error) {
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
	type Alias Getalertquery
	
	return json.Marshal(&struct { 
		RuleType *string `json:"ruleType,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		AlertStatus *string `json:"alertStatus,omitempty"`
		
		ViewedStatus *string `json:"viewedStatus,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		Alias
	}{ 
		RuleType: o.RuleType,
		
		QueryType: o.QueryType,
		
		AlertStatus: o.AlertStatus,
		
		ViewedStatus: o.ViewedStatus,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		SortBy: o.SortBy,
		
		SortOrder: o.SortOrder,
		Alias:    (Alias)(o),
	})
}

func (o *Getalertquery) UnmarshalJSON(b []byte) error {
	var GetalertqueryMap map[string]interface{}
	err := json.Unmarshal(b, &GetalertqueryMap)
	if err != nil {
		return err
	}
	
	if RuleType, ok := GetalertqueryMap["ruleType"].(string); ok {
		o.RuleType = &RuleType
	}
    
	if QueryType, ok := GetalertqueryMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if AlertStatus, ok := GetalertqueryMap["alertStatus"].(string); ok {
		o.AlertStatus = &AlertStatus
	}
    
	if ViewedStatus, ok := GetalertqueryMap["viewedStatus"].(string); ok {
		o.ViewedStatus = &ViewedStatus
	}
    
	if PageNumber, ok := GetalertqueryMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := GetalertqueryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if SortBy, ok := GetalertqueryMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if SortOrder, ok := GetalertqueryMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Getalertquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
