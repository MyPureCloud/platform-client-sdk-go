package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchrequest
type Knowledgesearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Input query to search content in the knowledge base
	Query *string `json:"query,omitempty"`

	// PageSize - Page size of the returned results
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results
	PageNumber *int `json:"pageNumber,omitempty"`

	// DocumentType - Document type to be used while searching
	DocumentType *string `json:"documentType,omitempty"`

	// LanguageCode - query search for specific languageCode
	LanguageCode *string `json:"languageCode,omitempty"`

	// SearchOnDraftDocuments - If true the search query will be executed on draft documents, else it will be on active documents
	SearchOnDraftDocuments *bool `json:"searchOnDraftDocuments,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesearchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesearchrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		DocumentType *string `json:"documentType,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		SearchOnDraftDocuments *bool `json:"searchOnDraftDocuments,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		DocumentType: o.DocumentType,
		
		LanguageCode: o.LanguageCode,
		
		SearchOnDraftDocuments: o.SearchOnDraftDocuments,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesearchrequest) UnmarshalJSON(b []byte) error {
	var KnowledgesearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgesearchrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgesearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgesearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if DocumentType, ok := KnowledgesearchrequestMap["documentType"].(string); ok {
		o.DocumentType = &DocumentType
	}
    
	if LanguageCode, ok := KnowledgesearchrequestMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if SearchOnDraftDocuments, ok := KnowledgesearchrequestMap["searchOnDraftDocuments"].(bool); ok {
		o.SearchOnDraftDocuments = &SearchOnDraftDocuments
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
