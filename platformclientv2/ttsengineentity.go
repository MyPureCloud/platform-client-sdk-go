package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ttsengineentity
type Ttsengineentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Languages - The set of languages the TTS engine supports
	Languages *[]string `json:"languages,omitempty"`


	// OutputFormats - The set of output formats the TTS engine can produce
	OutputFormats *[]string `json:"outputFormats,omitempty"`


	// Voices - The set of voices the TTS engine supports
	Voices *[]Ttsvoiceentity `json:"voices,omitempty"`


	// IsDefault - The TTS engine is the global default engine
	IsDefault *bool `json:"isDefault,omitempty"`


	// IsSecure - The TTS engine can be used in a secure call flow
	IsSecure *bool `json:"isSecure,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Ttsengineentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ttsengineentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		OutputFormats *[]string `json:"outputFormats,omitempty"`
		
		Voices *[]Ttsvoiceentity `json:"voices,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		IsSecure *bool `json:"isSecure,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Languages: o.Languages,
		
		OutputFormats: o.OutputFormats,
		
		Voices: o.Voices,
		
		IsDefault: o.IsDefault,
		
		IsSecure: o.IsSecure,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Ttsengineentity) UnmarshalJSON(b []byte) error {
	var TtsengineentityMap map[string]interface{}
	err := json.Unmarshal(b, &TtsengineentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TtsengineentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TtsengineentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Languages, ok := TtsengineentityMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if OutputFormats, ok := TtsengineentityMap["outputFormats"].([]interface{}); ok {
		OutputFormatsString, _ := json.Marshal(OutputFormats)
		json.Unmarshal(OutputFormatsString, &o.OutputFormats)
	}
	
	if Voices, ok := TtsengineentityMap["voices"].([]interface{}); ok {
		VoicesString, _ := json.Marshal(Voices)
		json.Unmarshal(VoicesString, &o.Voices)
	}
	
	if IsDefault, ok := TtsengineentityMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if IsSecure, ok := TtsengineentityMap["isSecure"].(bool); ok {
		o.IsSecure = &IsSecure
	}
    
	if SelfUri, ok := TtsengineentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ttsengineentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
