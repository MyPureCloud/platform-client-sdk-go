package platformclientv2
import (
	"time"
	"encoding/json"
)

// Evaluationform
type Evaluationform struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The evaluation form name
	Name *string `json:"name,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// QuestionGroups - A list of question groups
	QuestionGroups *[]Evaluationquestiongroup `json:"questionGroups,omitempty"`


	// PublishedVersions
	PublishedVersions *Domainentitylistingevaluationform `json:"publishedVersions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationform) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
