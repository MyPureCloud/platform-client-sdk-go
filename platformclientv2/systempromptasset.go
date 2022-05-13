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

func (o *Systempromptasset) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		PromptId: o.PromptId,
		
		Language: o.Language,
		
		DurationSeconds: o.DurationSeconds,
		
		MediaUri: o.MediaUri,
		
		TtsString: o.TtsString,
		
		Text: o.Text,
		
		UploadUri: o.UploadUri,
		
		UploadStatus: o.UploadStatus,
		
		HasDefault: o.HasDefault,
		
		LanguageDefault: o.LanguageDefault,
		
		Tags: o.Tags,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Systempromptasset) UnmarshalJSON(b []byte) error {
	var SystempromptassetMap map[string]interface{}
	err := json.Unmarshal(b, &SystempromptassetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SystempromptassetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SystempromptassetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PromptId, ok := SystempromptassetMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
    
	if Language, ok := SystempromptassetMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if DurationSeconds, ok := SystempromptassetMap["durationSeconds"].(float64); ok {
		o.DurationSeconds = &DurationSeconds
	}
    
	if MediaUri, ok := SystempromptassetMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if TtsString, ok := SystempromptassetMap["ttsString"].(string); ok {
		o.TtsString = &TtsString
	}
    
	if Text, ok := SystempromptassetMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if UploadUri, ok := SystempromptassetMap["uploadUri"].(string); ok {
		o.UploadUri = &UploadUri
	}
    
	if UploadStatus, ok := SystempromptassetMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    
	if HasDefault, ok := SystempromptassetMap["hasDefault"].(bool); ok {
		o.HasDefault = &HasDefault
	}
    
	if LanguageDefault, ok := SystempromptassetMap["languageDefault"].(bool); ok {
		o.LanguageDefault = &LanguageDefault
	}
    
	if Tags, ok := SystempromptassetMap["tags"].(map[string]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if SelfUri, ok := SystempromptassetMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Systempromptasset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
