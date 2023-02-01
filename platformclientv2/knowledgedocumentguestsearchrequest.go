package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentguestsearchrequest
type Knowledgedocumentguestsearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
	Query *string `json:"query,omitempty"`

	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results.
	PageNumber *int `json:"pageNumber,omitempty"`

	// SearchId - The globally unique identifier for the search.
	SearchId *string `json:"searchId,omitempty"`

	// Total - The total number of documents matching the query.
	Total *int `json:"total,omitempty"`

	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int `json:"pageCount,omitempty"`

	// QueryType - The type of the query that initiates the search.
	QueryType *string `json:"queryType,omitempty"`

	// SessionId - Session ID of the search.
	SessionId *string `json:"sessionId,omitempty"`

	// IncludeDraftDocuments - Indicates whether the search results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

	// App - The app where the session is started.
	App *Knowledgeguestsessionapp `json:"app,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentguestsearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentguestsearchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentguestsearchrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		
		App *Knowledgeguestsessionapp `json:"app,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		QueryType: o.QueryType,
		
		SessionId: o.SessionId,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		
		App: o.App,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentguestsearchrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentguestsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentguestsearchrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentguestsearchrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentguestsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentguestsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if SearchId, ok := KnowledgedocumentguestsearchrequestMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Total, ok := KnowledgedocumentguestsearchrequestMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgedocumentguestsearchrequestMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if QueryType, ok := KnowledgedocumentguestsearchrequestMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SessionId, ok := KnowledgedocumentguestsearchrequestMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if IncludeDraftDocuments, ok := KnowledgedocumentguestsearchrequestMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    
	if App, ok := KnowledgedocumentguestsearchrequestMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
