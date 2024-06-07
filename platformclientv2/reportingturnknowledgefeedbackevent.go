package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledgefeedbackevent
type Reportingturnknowledgefeedbackevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SearchId - The ID of this knowledge search.
	SearchId *string `json:"searchId,omitempty"`

	// KnowledgeBaseId - The Knowledge Base ID that the captured knowledge data relates to.
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty"`

	// Documents - The list of search documents that the feedback applies to.
	Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`

	// FeedbackRating - The feedback rating for the search (1.0 - 5.0). 1 = Negative, 5 = Positive.
	FeedbackRating *int `json:"feedbackRating,omitempty"`

	// DocumentVariationId - The variation of the document.
	DocumentVariationId *string `json:"documentVariationId,omitempty"`

	// DocumentVersionId - The version of the document.
	DocumentVersionId *string `json:"documentVersionId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingturnknowledgefeedbackevent) SetField(field string, fieldValue interface{}) {
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

func (o Reportingturnknowledgefeedbackevent) MarshalJSON() ([]byte, error) {
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
	type Alias Reportingturnknowledgefeedbackevent
	
	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty"`
		
		Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`
		
		FeedbackRating *int `json:"feedbackRating,omitempty"`
		
		DocumentVariationId *string `json:"documentVariationId,omitempty"`
		
		DocumentVersionId *string `json:"documentVersionId,omitempty"`
		Alias
	}{ 
		SearchId: o.SearchId,
		
		KnowledgeBaseId: o.KnowledgeBaseId,
		
		Documents: o.Documents,
		
		FeedbackRating: o.FeedbackRating,
		
		DocumentVariationId: o.DocumentVariationId,
		
		DocumentVersionId: o.DocumentVersionId,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingturnknowledgefeedbackevent) UnmarshalJSON(b []byte) error {
	var ReportingturnknowledgefeedbackeventMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnknowledgefeedbackeventMap)
	if err != nil {
		return err
	}
	
	if SearchId, ok := ReportingturnknowledgefeedbackeventMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if KnowledgeBaseId, ok := ReportingturnknowledgefeedbackeventMap["knowledgeBaseId"].(string); ok {
		o.KnowledgeBaseId = &KnowledgeBaseId
	}
    
	if Documents, ok := ReportingturnknowledgefeedbackeventMap["documents"].([]interface{}); ok {
		DocumentsString, _ := json.Marshal(Documents)
		json.Unmarshal(DocumentsString, &o.Documents)
	}
	
	if FeedbackRating, ok := ReportingturnknowledgefeedbackeventMap["feedbackRating"].(float64); ok {
		FeedbackRatingInt := int(FeedbackRating)
		o.FeedbackRating = &FeedbackRatingInt
	}
	
	if DocumentVariationId, ok := ReportingturnknowledgefeedbackeventMap["documentVariationId"].(string); ok {
		o.DocumentVariationId = &DocumentVariationId
	}
    
	if DocumentVersionId, ok := ReportingturnknowledgefeedbackeventMap["documentVersionId"].(string); ok {
		o.DocumentVersionId = &DocumentVersionId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgefeedbackevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
