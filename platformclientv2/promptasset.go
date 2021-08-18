package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Promptasset
type Promptasset struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PromptId - Associated prompt ID
	PromptId *string `json:"promptId,omitempty"`


	// Language - Prompt resource language
	Language *string `json:"language,omitempty"`


	// MediaUri - URI of the resource audio
	MediaUri *string `json:"mediaUri,omitempty"`


	// TtsString - Text to speech of the resource
	TtsString *string `json:"ttsString,omitempty"`


	// Text - Text of the resource
	Text *string `json:"text,omitempty"`


	// UploadStatus - Audio upload status
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// UploadUri - Upload URI for the resource audio
	UploadUri *string `json:"uploadUri,omitempty"`


	// LanguageDefault - Whether or not this resource locale is the default for the language
	LanguageDefault *bool `json:"languageDefault,omitempty"`


	// Tags
	Tags *map[string][]string `json:"tags,omitempty"`


	// DurationSeconds
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Promptasset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Promptasset

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PromptId *string `json:"promptId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		TtsString *string `json:"ttsString,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		UploadUri *string `json:"uploadUri,omitempty"`
		
		LanguageDefault *bool `json:"languageDefault,omitempty"`
		
		Tags *map[string][]string `json:"tags,omitempty"`
		
		DurationSeconds *float64 `json:"durationSeconds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		PromptId: u.PromptId,
		
		Language: u.Language,
		
		MediaUri: u.MediaUri,
		
		TtsString: u.TtsString,
		
		Text: u.Text,
		
		UploadStatus: u.UploadStatus,
		
		UploadUri: u.UploadUri,
		
		LanguageDefault: u.LanguageDefault,
		
		Tags: u.Tags,
		
		DurationSeconds: u.DurationSeconds,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Promptasset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
