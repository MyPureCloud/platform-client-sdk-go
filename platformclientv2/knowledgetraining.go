package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgetraining
type Knowledgetraining struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateTriggered - Trigger date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTriggered *time.Time `json:"dateTriggered,omitempty"`


	// DateCompleted - Training completed date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// Status - Training status.
	Status *string `json:"status,omitempty"`


	// LanguageCode - Language of the documents that are trained.
	LanguageCode *string `json:"languageCode,omitempty"`


	// KnowledgeBase - Knowledge Base that the training belongs to.
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// ErrorMessage - Any error message during the Training or Promote action.
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// KnowledgeDocumentsState - State of the Trained Documents, which can be one of these Draft, Active, Discarded, Archived.
	KnowledgeDocumentsState *string `json:"knowledgeDocumentsState,omitempty"`


	// DatePromoted - Trained Documents Promoted date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePromoted *time.Time `json:"datePromoted,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgetraining) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgetraining
	
	DateTriggered := new(string)
	if o.DateTriggered != nil {
		
		*DateTriggered = timeutil.Strftime(o.DateTriggered, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTriggered = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DatePromoted := new(string)
	if o.DatePromoted != nil {
		
		*DatePromoted = timeutil.Strftime(o.DatePromoted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePromoted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateTriggered *string `json:"dateTriggered,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		KnowledgeDocumentsState *string `json:"knowledgeDocumentsState,omitempty"`
		
		DatePromoted *string `json:"datePromoted,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DateTriggered: DateTriggered,
		
		DateCompleted: DateCompleted,
		
		Status: o.Status,
		
		LanguageCode: o.LanguageCode,
		
		KnowledgeBase: o.KnowledgeBase,
		
		ErrorMessage: o.ErrorMessage,
		
		KnowledgeDocumentsState: o.KnowledgeDocumentsState,
		
		DatePromoted: DatePromoted,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgetraining) UnmarshalJSON(b []byte) error {
	var KnowledgetrainingMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgetrainingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgetrainingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateTriggeredString, ok := KnowledgetrainingMap["dateTriggered"].(string); ok {
		DateTriggered, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateTriggeredString)
		o.DateTriggered = &DateTriggered
	}
	
	if dateCompletedString, ok := KnowledgetrainingMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if Status, ok := KnowledgetrainingMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if LanguageCode, ok := KnowledgetrainingMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if KnowledgeBase, ok := KnowledgetrainingMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if ErrorMessage, ok := KnowledgetrainingMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if KnowledgeDocumentsState, ok := KnowledgetrainingMap["knowledgeDocumentsState"].(string); ok {
		o.KnowledgeDocumentsState = &KnowledgeDocumentsState
	}
    
	if datePromotedString, ok := KnowledgetrainingMap["datePromoted"].(string); ok {
		DatePromoted, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePromotedString)
		o.DatePromoted = &DatePromoted
	}
	
	if SelfUri, ok := KnowledgetrainingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgetraining) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
