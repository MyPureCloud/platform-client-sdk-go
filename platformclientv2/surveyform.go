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

func (o *Surveyform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyform
	
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
		Id: o.Id,
		
		Name: o.Name,
		
		ModifiedDate: ModifiedDate,
		
		Published: o.Published,
		
		Disabled: o.Disabled,
		
		ContextId: o.ContextId,
		
		Language: o.Language,
		
		Header: o.Header,
		
		Footer: o.Footer,
		
		QuestionGroups: o.QuestionGroups,
		
		PublishedVersions: o.PublishedVersions,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyform) UnmarshalJSON(b []byte) error {
	var SurveyformMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyformMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SurveyformMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SurveyformMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if modifiedDateString, ok := SurveyformMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if Published, ok := SurveyformMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if Disabled, ok := SurveyformMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
	
	if ContextId, ok := SurveyformMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if Language, ok := SurveyformMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Header, ok := SurveyformMap["header"].(string); ok {
		o.Header = &Header
	}
	
	if Footer, ok := SurveyformMap["footer"].(string); ok {
		o.Footer = &Footer
	}
	
	if QuestionGroups, ok := SurveyformMap["questionGroups"].([]interface{}); ok {
		QuestionGroupsString, _ := json.Marshal(QuestionGroups)
		json.Unmarshal(QuestionGroupsString, &o.QuestionGroups)
	}
	
	if PublishedVersions, ok := SurveyformMap["publishedVersions"].(map[string]interface{}); ok {
		PublishedVersionsString, _ := json.Marshal(PublishedVersions)
		json.Unmarshal(PublishedVersionsString, &o.PublishedVersions)
	}
	
	if SelfUri, ok := SurveyformMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
