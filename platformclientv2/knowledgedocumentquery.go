package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentquery
type Knowledgedocumentquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results.
	PageNumber *int `json:"pageNumber,omitempty"`

	// IncludeDraftDocuments - Indicates whether the results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

	// Interval - Retrieves the documents created/modified/published in specified date and time range.
	Interval *Documentqueryinterval `json:"interval,omitempty"`

	// Filter - Filter for the document query.
	Filter *Documentquery `json:"filter,omitempty"`

	// IncludeVariations - Indicates which document variations to include in returned documents. All: all variations regardless of the filter expression; AllMatching: all variations that match the filter expression; SingleMostRelevant: single variation that matches the filter expression and has the highest priority. The default is All.
	IncludeVariations *string `json:"includeVariations,omitempty"`

	// SortOrder - The sort order for results.
	SortOrder *string `json:"sortOrder,omitempty"`

	// SortBy - The field in the documents that you want to sort the results by.
	SortBy *string `json:"sortBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentquery) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentquery) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentquery
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		
		Interval *Documentqueryinterval `json:"interval,omitempty"`
		
		Filter *Documentquery `json:"filter,omitempty"`
		
		IncludeVariations *string `json:"includeVariations,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		Alias
	}{ 
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		
		Interval: o.Interval,
		
		Filter: o.Filter,
		
		IncludeVariations: o.IncludeVariations,
		
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentquery) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentqueryMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentqueryMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := KnowledgedocumentqueryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentqueryMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if IncludeDraftDocuments, ok := KnowledgedocumentqueryMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    
	if Interval, ok := KnowledgedocumentqueryMap["interval"].(map[string]interface{}); ok {
		IntervalString, _ := json.Marshal(Interval)
		json.Unmarshal(IntervalString, &o.Interval)
	}
	
	if Filter, ok := KnowledgedocumentqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if IncludeVariations, ok := KnowledgedocumentqueryMap["includeVariations"].(string); ok {
		o.IncludeVariations = &IncludeVariations
	}
    
	if SortOrder, ok := KnowledgedocumentqueryMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortBy, ok := KnowledgedocumentqueryMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
