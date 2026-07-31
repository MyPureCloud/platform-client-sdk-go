package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledgemetadata
type Reportingturnknowledgemetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// KnowledgeId - The ID of the knowledge setting or knowledge base
	KnowledgeId *string `json:"knowledgeId,omitempty"`

	// KnowledgeName - The name of the knowledge setting or knowledge base
	KnowledgeName *string `json:"knowledgeName,omitempty"`

	// SearchId - SearchID used in the attempted search
	SearchId *string `json:"searchId,omitempty"`

	// Query - The query used in the knowledge query
	Query *string `json:"query,omitempty"`

	// RetrievalStatus - The result of the knowledge search
	RetrievalStatus *string `json:"retrievalStatus,omitempty"`

	// AnswerGenerationStatus - The result of the knowledge generation
	AnswerGenerationStatus *string `json:"answerGenerationStatus,omitempty"`

	// GeneratedAnswer - The generated answer
	GeneratedAnswer *string `json:"generatedAnswer,omitempty"`

	// FailureReason - Failure reason if knowledge query failed
	FailureReason *string `json:"failureReason,omitempty"`

	// TopConfidence - Highest confidence score of returned knowledgeSources
	TopConfidence *float64 `json:"topConfidence,omitempty"`

	// RetrievedSources - List of the sources retrieved by the knowledge search
	RetrievedSources *[]Knowledgesource `json:"retrievedSources,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingturnknowledgemetadata) SetField(field string, fieldValue interface{}) {
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

func (o Reportingturnknowledgemetadata) MarshalJSON() ([]byte, error) {
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
	type Alias Reportingturnknowledgemetadata
	
	return json.Marshal(&struct { 
		KnowledgeId *string `json:"knowledgeId,omitempty"`
		
		KnowledgeName *string `json:"knowledgeName,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Query *string `json:"query,omitempty"`
		
		RetrievalStatus *string `json:"retrievalStatus,omitempty"`
		
		AnswerGenerationStatus *string `json:"answerGenerationStatus,omitempty"`
		
		GeneratedAnswer *string `json:"generatedAnswer,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		TopConfidence *float64 `json:"topConfidence,omitempty"`
		
		RetrievedSources *[]Knowledgesource `json:"retrievedSources,omitempty"`
		Alias
	}{ 
		KnowledgeId: o.KnowledgeId,
		
		KnowledgeName: o.KnowledgeName,
		
		SearchId: o.SearchId,
		
		Query: o.Query,
		
		RetrievalStatus: o.RetrievalStatus,
		
		AnswerGenerationStatus: o.AnswerGenerationStatus,
		
		GeneratedAnswer: o.GeneratedAnswer,
		
		FailureReason: o.FailureReason,
		
		TopConfidence: o.TopConfidence,
		
		RetrievedSources: o.RetrievedSources,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingturnknowledgemetadata) UnmarshalJSON(b []byte) error {
	var ReportingturnknowledgemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnknowledgemetadataMap)
	if err != nil {
		return err
	}
	
	if KnowledgeId, ok := ReportingturnknowledgemetadataMap["knowledgeId"].(string); ok {
		o.KnowledgeId = &KnowledgeId
	}
    
	if KnowledgeName, ok := ReportingturnknowledgemetadataMap["knowledgeName"].(string); ok {
		o.KnowledgeName = &KnowledgeName
	}
    
	if SearchId, ok := ReportingturnknowledgemetadataMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Query, ok := ReportingturnknowledgemetadataMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if RetrievalStatus, ok := ReportingturnknowledgemetadataMap["retrievalStatus"].(string); ok {
		o.RetrievalStatus = &RetrievalStatus
	}
    
	if AnswerGenerationStatus, ok := ReportingturnknowledgemetadataMap["answerGenerationStatus"].(string); ok {
		o.AnswerGenerationStatus = &AnswerGenerationStatus
	}
    
	if GeneratedAnswer, ok := ReportingturnknowledgemetadataMap["generatedAnswer"].(string); ok {
		o.GeneratedAnswer = &GeneratedAnswer
	}
    
	if FailureReason, ok := ReportingturnknowledgemetadataMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    
	if TopConfidence, ok := ReportingturnknowledgemetadataMap["topConfidence"].(float64); ok {
		o.TopConfidence = &TopConfidence
	}
    
	if RetrievedSources, ok := ReportingturnknowledgemetadataMap["retrievedSources"].([]interface{}); ok {
		RetrievedSourcesString, _ := json.Marshal(RetrievedSources)
		json.Unmarshal(RetrievedSourcesString, &o.RetrievedSources)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
