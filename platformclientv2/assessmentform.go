package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentform
type Assessmentform struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateModified - Last modified date of the assessment form. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContextId - The unique Id for all versions of this assessment form
	ContextId *string `json:"contextId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Published - If true, assessment form is published
	Published *bool `json:"published,omitempty"`


	// PassPercent - The pass percent for the assessment form
	PassPercent *int `json:"passPercent,omitempty"`


	// QuestionGroups - A list of question groups
	QuestionGroups *[]Assessmentformquestiongroup `json:"questionGroups,omitempty"`

}

func (o *Assessmentform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentform
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		PassPercent *int `json:"passPercent,omitempty"`
		
		QuestionGroups *[]Assessmentformquestiongroup `json:"questionGroups,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DateModified: DateModified,
		
		ContextId: o.ContextId,
		
		SelfUri: o.SelfUri,
		
		Published: o.Published,
		
		PassPercent: o.PassPercent,
		
		QuestionGroups: o.QuestionGroups,
		Alias:    (*Alias)(o),
	})
}

func (o *Assessmentform) UnmarshalJSON(b []byte) error {
	var AssessmentformMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentformMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssessmentformMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateModifiedString, ok := AssessmentformMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ContextId, ok := AssessmentformMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if SelfUri, ok := AssessmentformMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Published, ok := AssessmentformMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if PassPercent, ok := AssessmentformMap["passPercent"].(float64); ok {
		PassPercentInt := int(PassPercent)
		o.PassPercent = &PassPercentInt
	}
	
	if QuestionGroups, ok := AssessmentformMap["questionGroups"].([]interface{}); ok {
		QuestionGroupsString, _ := json.Marshal(QuestionGroups)
		json.Unmarshal(QuestionGroupsString, &o.QuestionGroups)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
