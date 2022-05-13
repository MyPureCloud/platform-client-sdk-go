package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ttsvoiceentity
type Ttsvoiceentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Gender - The gender of the TTS voice
	Gender *string `json:"gender,omitempty"`


	// Language - The language supported by the TTS voice
	Language *string `json:"language,omitempty"`


	// Engine - Ths TTS engine this voice belongs to
	Engine *Ttsengineentity `json:"engine,omitempty"`


	// IsDefault - The voice is the default voice for its language
	IsDefault *bool `json:"isDefault,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Ttsvoiceentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ttsvoiceentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Gender *string `json:"gender,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Engine *Ttsengineentity `json:"engine,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Gender: o.Gender,
		
		Language: o.Language,
		
		Engine: o.Engine,
		
		IsDefault: o.IsDefault,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Ttsvoiceentity) UnmarshalJSON(b []byte) error {
	var TtsvoiceentityMap map[string]interface{}
	err := json.Unmarshal(b, &TtsvoiceentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TtsvoiceentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TtsvoiceentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Gender, ok := TtsvoiceentityMap["gender"].(string); ok {
		o.Gender = &Gender
	}
    
	if Language, ok := TtsvoiceentityMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Engine, ok := TtsvoiceentityMap["engine"].(map[string]interface{}); ok {
		EngineString, _ := json.Marshal(Engine)
		json.Unmarshal(EngineString, &o.Engine)
	}
	
	if IsDefault, ok := TtsvoiceentityMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if SelfUri, ok := TtsvoiceentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ttsvoiceentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
