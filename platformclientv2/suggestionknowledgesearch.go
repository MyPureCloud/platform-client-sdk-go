package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestionknowledgesearch
type Suggestionknowledgesearch struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - The article title.
	Title *string `json:"title,omitempty"`

	// Snippets - Snippets of text from the article matching the query.
	Snippets *[]string `json:"snippets,omitempty"`

	// Confidence - Value between 0 and 1. 1 corresponds to very confident, 0 to not confident at all.
	Confidence *float32 `json:"confidence,omitempty"`

	// SearchId - The search id.
	SearchId *string `json:"searchId,omitempty"`

	// Document - The article matching the query.
	Document *Addressableentityref `json:"document,omitempty"`

	// Version - The version of the article.
	Version *Addressableentityref `json:"version,omitempty"`

	// KnowledgeAnswer - The most relevant answer within a searched article for the searched query
	KnowledgeAnswer *Suggestionknowledgeanswer `json:"knowledgeAnswer,omitempty"`

	// Variations - Variations of the article.
	Variations *[]Addressableentityref `json:"variations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Suggestionknowledgesearch) SetField(field string, fieldValue interface{}) {
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

func (o Suggestionknowledgesearch) MarshalJSON() ([]byte, error) {
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
	type Alias Suggestionknowledgesearch
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Snippets *[]string `json:"snippets,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Document *Addressableentityref `json:"document,omitempty"`
		
		Version *Addressableentityref `json:"version,omitempty"`
		
		KnowledgeAnswer *Suggestionknowledgeanswer `json:"knowledgeAnswer,omitempty"`
		
		Variations *[]Addressableentityref `json:"variations,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Snippets: o.Snippets,
		
		Confidence: o.Confidence,
		
		SearchId: o.SearchId,
		
		Document: o.Document,
		
		Version: o.Version,
		
		KnowledgeAnswer: o.KnowledgeAnswer,
		
		Variations: o.Variations,
		Alias:    (Alias)(o),
	})
}

func (o *Suggestionknowledgesearch) UnmarshalJSON(b []byte) error {
	var SuggestionknowledgesearchMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestionknowledgesearchMap)
	if err != nil {
		return err
	}
	
	if Title, ok := SuggestionknowledgesearchMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Snippets, ok := SuggestionknowledgesearchMap["snippets"].([]interface{}); ok {
		SnippetsString, _ := json.Marshal(Snippets)
		json.Unmarshal(SnippetsString, &o.Snippets)
	}
	
	if Confidence, ok := SuggestionknowledgesearchMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if SearchId, ok := SuggestionknowledgesearchMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Document, ok := SuggestionknowledgesearchMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if Version, ok := SuggestionknowledgesearchMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if KnowledgeAnswer, ok := SuggestionknowledgesearchMap["knowledgeAnswer"].(map[string]interface{}); ok {
		KnowledgeAnswerString, _ := json.Marshal(KnowledgeAnswer)
		json.Unmarshal(KnowledgeAnswerString, &o.KnowledgeAnswer)
	}
	
	if Variations, ok := SuggestionknowledgesearchMap["variations"].([]interface{}); ok {
		VariationsString, _ := json.Marshal(Variations)
		json.Unmarshal(VariationsString, &o.Variations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestionknowledgesearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
