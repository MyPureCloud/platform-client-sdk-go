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

func (u *Knowledgetraining) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgetraining

	
	DateTriggered := new(string)
	if u.DateTriggered != nil {
		
		*DateTriggered = timeutil.Strftime(u.DateTriggered, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTriggered = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DatePromoted := new(string)
	if u.DatePromoted != nil {
		
		*DatePromoted = timeutil.Strftime(u.DatePromoted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		DateTriggered: DateTriggered,
		
		DateCompleted: DateCompleted,
		
		Status: u.Status,
		
		LanguageCode: u.LanguageCode,
		
		KnowledgeBase: u.KnowledgeBase,
		
		ErrorMessage: u.ErrorMessage,
		
		KnowledgeDocumentsState: u.KnowledgeDocumentsState,
		
		DatePromoted: DatePromoted,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgetraining) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
