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

func (u *Ttsengineentity) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Languages: u.Languages,
		
		OutputFormats: u.OutputFormats,
		
		Voices: u.Voices,
		
		IsDefault: u.IsDefault,
		
		IsSecure: u.IsSecure,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Ttsengineentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
