package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Systempromptasset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
