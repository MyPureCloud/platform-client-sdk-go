package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequest
type Queryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`

	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

	// FacetNameRequests
	FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`

	// Sort
	Sort *[]Sortitem `json:"sort,omitempty"`

	// Filters
	Filters *[]Contentfilteritem `json:"filters,omitempty"`

	// AttributeFilters
	AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`

	// IncludeShares
	IncludeShares *bool `json:"includeShares,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Queryrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Queryrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`
		
		Sort *[]Sortitem `json:"sort,omitempty"`
		
		Filters *[]Contentfilteritem `json:"filters,omitempty"`
		
		AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`
		
		IncludeShares *bool `json:"includeShares,omitempty"`
		Alias
	}{ 
		QueryPhrase: o.QueryPhrase,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		FacetNameRequests: o.FacetNameRequests,
		
		Sort: o.Sort,
		
		Filters: o.Filters,
		
		AttributeFilters: o.AttributeFilters,
		
		IncludeShares: o.IncludeShares,
		Alias:    (Alias)(o),
	})
}

func (o *Queryrequest) UnmarshalJSON(b []byte) error {
	var QueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &QueryrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := QueryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if PageNumber, ok := QueryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := QueryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if FacetNameRequests, ok := QueryrequestMap["facetNameRequests"].([]interface{}); ok {
		FacetNameRequestsString, _ := json.Marshal(FacetNameRequests)
		json.Unmarshal(FacetNameRequestsString, &o.FacetNameRequests)
	}
	
	if Sort, ok := QueryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Filters, ok := QueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if AttributeFilters, ok := QueryrequestMap["attributeFilters"].([]interface{}); ok {
		AttributeFiltersString, _ := json.Marshal(AttributeFilters)
		json.Unmarshal(AttributeFiltersString, &o.AttributeFilters)
	}
	
	if IncludeShares, ok := QueryrequestMap["includeShares"].(bool); ok {
		o.IncludeShares = &IncludeShares
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
