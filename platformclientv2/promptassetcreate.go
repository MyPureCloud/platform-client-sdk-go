package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Promptassetcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Promptassetcreate
	
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
		Id: o.Id,
		
		Name: o.Name,
		
		PromptId: o.PromptId,
		
		Language: o.Language,
		
		MediaUri: o.MediaUri,
		
		TtsString: o.TtsString,
		
		Text: o.Text,
		
		UploadStatus: o.UploadStatus,
		
		UploadUri: o.UploadUri,
		
		LanguageDefault: o.LanguageDefault,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Promptassetcreate) UnmarshalJSON(b []byte) error {
	var PromptassetcreateMap map[string]interface{}
	err := json.Unmarshal(b, &PromptassetcreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PromptassetcreateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PromptassetcreateMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if PromptId, ok := PromptassetcreateMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
	
	if Language, ok := PromptassetcreateMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if MediaUri, ok := PromptassetcreateMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
	
	if TtsString, ok := PromptassetcreateMap["ttsString"].(string); ok {
		o.TtsString = &TtsString
	}
	
	if Text, ok := PromptassetcreateMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if UploadStatus, ok := PromptassetcreateMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	
	if UploadUri, ok := PromptassetcreateMap["uploadUri"].(string); ok {
		o.UploadUri = &UploadUri
	}
	
	if LanguageDefault, ok := PromptassetcreateMap["languageDefault"].(bool); ok {
		o.LanguageDefault = &LanguageDefault
	}
	
	if Tags, ok := PromptassetcreateMap["tags"].(map[string]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := PromptassetcreateMap["durationSeconds"].(float64); ok {
		o.DurationSeconds = &DurationSeconds
	}
	
	if SelfUri, ok := PromptassetcreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Promptassetcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
