package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestdocumentpresentation
type Knowledgeguestdocumentpresentation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Documents - The presented documents
	Documents *[]Presentedknowledgedocument `json:"documents,omitempty"`

	// SearchId - The search that surfaced the documents that were presented.
	SearchId *string `json:"searchId,omitempty"`

	// QueryType - The type of the query that surfaced the documents.
	QueryType *string `json:"queryType,omitempty"`

	// SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.
	SurfacingMethod *string `json:"surfacingMethod,omitempty"`

	// SessionId - Knowledge session ID.
	SessionId *string `json:"sessionId,omitempty"`

	// Application - The client application in which the documents were presented.
	Application *Knowledgeguestsearchclientapplication `json:"application,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeguestdocumentpresentation) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeguestdocumentpresentation) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeguestdocumentpresentation
	
	return json.Marshal(&struct { 
		Documents *[]Presentedknowledgedocument `json:"documents,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SurfacingMethod *string `json:"surfacingMethod,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		Application *Knowledgeguestsearchclientapplication `json:"application,omitempty"`
		Alias
	}{ 
		Documents: o.Documents,
		
		SearchId: o.SearchId,
		
		QueryType: o.QueryType,
		
		SurfacingMethod: o.SurfacingMethod,
		
		SessionId: o.SessionId,
		
		Application: o.Application,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeguestdocumentpresentation) UnmarshalJSON(b []byte) error {
	var KnowledgeguestdocumentpresentationMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestdocumentpresentationMap)
	if err != nil {
		return err
	}
	
	if Documents, ok := KnowledgeguestdocumentpresentationMap["documents"].([]interface{}); ok {
		DocumentsString, _ := json.Marshal(Documents)
		json.Unmarshal(DocumentsString, &o.Documents)
	}
	
	if SearchId, ok := KnowledgeguestdocumentpresentationMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if QueryType, ok := KnowledgeguestdocumentpresentationMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SurfacingMethod, ok := KnowledgeguestdocumentpresentationMap["surfacingMethod"].(string); ok {
		o.SurfacingMethod = &SurfacingMethod
	}
    
	if SessionId, ok := KnowledgeguestdocumentpresentationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if Application, ok := KnowledgeguestdocumentpresentationMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentpresentation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
