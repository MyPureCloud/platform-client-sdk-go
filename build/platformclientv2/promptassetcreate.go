package platformclientv2
import (
	"encoding/json"
)

// Promptassetcreate
type Promptassetcreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PromptId - Associated prompt ID
	PromptId *string `json:"promptId,omitempty"`


	// Language - The prompt language.
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

// String returns a JSON representation of the model
func (o *Promptassetcreate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
