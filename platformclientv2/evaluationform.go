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

func (o *Evaluationform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationform
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		ModifiedDate: ModifiedDate,
		
		Published: o.Published,
		
		ContextId: o.ContextId,
		
		QuestionGroups: o.QuestionGroups,
		
		PublishedVersions: o.PublishedVersions,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationform) UnmarshalJSON(b []byte) error {
	var EvaluationformMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationformMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationformMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluationformMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if modifiedDateString, ok := EvaluationformMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if Published, ok := EvaluationformMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ContextId, ok := EvaluationformMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if QuestionGroups, ok := EvaluationformMap["questionGroups"].([]interface{}); ok {
		QuestionGroupsString, _ := json.Marshal(QuestionGroups)
		json.Unmarshal(QuestionGroupsString, &o.QuestionGroups)
	}
	
	if PublishedVersions, ok := EvaluationformMap["publishedVersions"].(map[string]interface{}); ok {
		PublishedVersionsString, _ := json.Marshal(PublishedVersions)
		json.Unmarshal(PublishedVersionsString, &o.PublishedVersions)
	}
	
	if SelfUri, ok := EvaluationformMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
