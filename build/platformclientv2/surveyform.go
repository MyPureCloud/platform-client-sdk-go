package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Surveyform) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
