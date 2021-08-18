package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Systempromptasset
type Systempromptasset struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PromptId
	PromptId *string `json:"promptId,omitempty"`


	// Language - The asset resource language
	Language *string `json:"language,omitempty"`


	// DurationSeconds
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`


	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// TtsString
	TtsString *string `json:"ttsString,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// UploadUri
	UploadUri *string `json:"uploadUri,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// HasDefault
	HasDefault *bool `json:"hasDefault,omitempty"`


	// LanguageDefault
	LanguageDefault *bool `json:"languageDefault,omitempty"`


	// Tags
	Tags *map[string][]string `json:"tags,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Systempromptasset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Systempromptasset

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PromptId *string `json:"promptId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		DurationSeconds *float64 `json:"durationSeconds,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		TtsString *string `json:"ttsString,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UploadUri *string `json:"uploadUri,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		HasDefault *bool `json:"hasDefault,omitempty"`
		
		LanguageDefault *bool `json:"languageDefault,omitempty"`
		
		Tags *map[string][]string `json:"tags,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		PromptId: u.PromptId,
		
		Language: u.Language,
		
		DurationSeconds: u.DurationSeconds,
		
		MediaUri: u.MediaUri,
		
		TtsString: u.TtsString,
		
		Text: u.Text,
		
		UploadUri: u.UploadUri,
		
		UploadStatus: u.UploadStatus,
		
		HasDefault: u.HasDefault,
		
		LanguageDefault: u.LanguageDefault,
		
		Tags: u.Tags,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Systempromptasset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
