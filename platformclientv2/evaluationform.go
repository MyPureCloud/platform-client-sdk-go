package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationform
type Evaluationform struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The evaluation form name
	Name *string `json:"name,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

func (u *Evaluationform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationform

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		QuestionGroups *[]Evaluationquestiongroup `json:"questionGroups,omitempty"`
		
		PublishedVersions *Domainentitylistingevaluationform `json:"publishedVersions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ModifiedDate: ModifiedDate,
		
		Published: u.Published,
		
		ContextId: u.ContextId,
		
		QuestionGroups: u.QuestionGroups,
		
		PublishedVersions: u.PublishedVersions,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
