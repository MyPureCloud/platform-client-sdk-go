package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyform
type Surveyform struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The survey form name
	Name *string `json:"name,omitempty"`


	// ModifiedDate - Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// Published - Is this form published
	Published *bool `json:"published,omitempty"`


	// Disabled - Is this form disabled
	Disabled *bool `json:"disabled,omitempty"`


	// ContextId - Unique Id for all versions of this form
	ContextId *string `json:"contextId,omitempty"`


	// Language - Language for survey viewer localization. Currently localized languages: da, de, en-US, es, fi, fr, it, ja, ko, nl, no, pl, pt-BR, sv, th, tr, zh-CH, zh-TW
	Language *string `json:"language,omitempty"`


	// Header - Markdown text for the top of the form.
	Header *string `json:"header,omitempty"`


	// Footer - Markdown text for the bottom of the form.
	Footer *string `json:"footer,omitempty"`


	// QuestionGroups - A list of question groups
	QuestionGroups *[]Surveyquestiongroup `json:"questionGroups,omitempty"`


	// PublishedVersions - List of published version of this form
	PublishedVersions *Domainentitylistingsurveyform `json:"publishedVersions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Surveyform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyform

	
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
		
		Disabled *bool `json:"disabled,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *string `json:"header,omitempty"`
		
		Footer *string `json:"footer,omitempty"`
		
		QuestionGroups *[]Surveyquestiongroup `json:"questionGroups,omitempty"`
		
		PublishedVersions *Domainentitylistingsurveyform `json:"publishedVersions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ModifiedDate: ModifiedDate,
		
		Published: u.Published,
		
		Disabled: u.Disabled,
		
		ContextId: u.ContextId,
		
		Language: u.Language,
		
		Header: u.Header,
		
		Footer: u.Footer,
		
		QuestionGroups: u.QuestionGroups,
		
		PublishedVersions: u.PublishedVersions,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
