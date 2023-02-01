package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupssearchresponse
type Groupssearchresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Total - The total number of results found
	Total *int `json:"total,omitempty"`

	// PageCount - The total number of pages
	PageCount *int `json:"pageCount,omitempty"`

	// PageSize - The current page size
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - The current page number
	PageNumber *int `json:"pageNumber,omitempty"`

	// PreviousPage - Q64 value for the previous page of results
	PreviousPage *string `json:"previousPage,omitempty"`

	// CurrentPage - Q64 value for the current page of results
	CurrentPage *string `json:"currentPage,omitempty"`

	// NextPage - Q64 value for the next page of results
	NextPage *string `json:"nextPage,omitempty"`

	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`

	// Results - Search results
	Results *[]Group `json:"results,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Groupssearchresponse) SetField(field string, fieldValue interface{}) {
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

func (o Groupssearchresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Groupssearchresponse
	
	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PreviousPage *string `json:"previousPage,omitempty"`
		
		CurrentPage *string `json:"currentPage,omitempty"`
		
		NextPage *string `json:"nextPage,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Results *[]Group `json:"results,omitempty"`
		Alias
	}{ 
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		PreviousPage: o.PreviousPage,
		
		CurrentPage: o.CurrentPage,
		
		NextPage: o.NextPage,
		
		Types: o.Types,
		
		Results: o.Results,
		Alias:    (Alias)(o),
	})
}

func (o *Groupssearchresponse) UnmarshalJSON(b []byte) error {
	var GroupssearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &GroupssearchresponseMap)
	if err != nil {
		return err
	}
	
	if Total, ok := GroupssearchresponseMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := GroupssearchresponseMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if PageSize, ok := GroupssearchresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := GroupssearchresponseMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PreviousPage, ok := GroupssearchresponseMap["previousPage"].(string); ok {
		o.PreviousPage = &PreviousPage
	}
    
	if CurrentPage, ok := GroupssearchresponseMap["currentPage"].(string); ok {
		o.CurrentPage = &CurrentPage
	}
    
	if NextPage, ok := GroupssearchresponseMap["nextPage"].(string); ok {
		o.NextPage = &NextPage
	}
    
	if Types, ok := GroupssearchresponseMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Results, ok := GroupssearchresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Groupssearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
