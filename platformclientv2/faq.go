package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Faq
type Faq struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Question - Question from the knowledge base that was matched to user request.
	Question *string `json:"question,omitempty"`

	// Answer - Answer from the knowledge base corresponding to the identified question.
	Answer *string `json:"answer,omitempty"`

	// SourceUri - A URI uniquely identifying the document, e.g. projects/acme-inc/knowledgeBases/MTAyNjgxNDU1Nzc3NTM1NzU0MjQ/documents/MTI5ODc3NzQzOTQ5MTc5NzgxMTI.
	SourceUri *string `json:"sourceUri,omitempty"`

	// DocumentUrl - URL pointing to a web page if document was sourced from a URL.
	DocumentUrl *string `json:"documentUrl,omitempty"`

	// DocumentDisplayName - A human-readable description of the document, e.g. 'Sample store FAQ'
	DocumentDisplayName *string `json:"documentDisplayName,omitempty"`

	// Confidence - Value between 0 and 1. 1 corresponds to very confident, 0 to not confident at all
	Confidence *float32 `json:"confidence,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Faq) SetField(field string, fieldValue interface{}) {
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

func (o Faq) MarshalJSON() ([]byte, error) {
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
	type Alias Faq
	
	return json.Marshal(&struct { 
		Question *string `json:"question,omitempty"`
		
		Answer *string `json:"answer,omitempty"`
		
		SourceUri *string `json:"sourceUri,omitempty"`
		
		DocumentUrl *string `json:"documentUrl,omitempty"`
		
		DocumentDisplayName *string `json:"documentDisplayName,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		Alias
	}{ 
		Question: o.Question,
		
		Answer: o.Answer,
		
		SourceUri: o.SourceUri,
		
		DocumentUrl: o.DocumentUrl,
		
		DocumentDisplayName: o.DocumentDisplayName,
		
		Confidence: o.Confidence,
		Alias:    (Alias)(o),
	})
}

func (o *Faq) UnmarshalJSON(b []byte) error {
	var FaqMap map[string]interface{}
	err := json.Unmarshal(b, &FaqMap)
	if err != nil {
		return err
	}
	
	if Question, ok := FaqMap["question"].(string); ok {
		o.Question = &Question
	}
    
	if Answer, ok := FaqMap["answer"].(string); ok {
		o.Answer = &Answer
	}
    
	if SourceUri, ok := FaqMap["sourceUri"].(string); ok {
		o.SourceUri = &SourceUri
	}
    
	if DocumentUrl, ok := FaqMap["documentUrl"].(string); ok {
		o.DocumentUrl = &DocumentUrl
	}
    
	if DocumentDisplayName, ok := FaqMap["documentDisplayName"].(string); ok {
		o.DocumentDisplayName = &DocumentDisplayName
	}
    
	if Confidence, ok := FaqMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Faq) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
