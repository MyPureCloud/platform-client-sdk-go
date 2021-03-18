package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Knowledgetraining) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
