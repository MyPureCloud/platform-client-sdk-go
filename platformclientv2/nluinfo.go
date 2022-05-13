package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluinfo
type Nluinfo struct { 
	// Domain
	Domain *Addressableentityref `json:"domain,omitempty"`


	// Version
	Version **Nludomainversion `json:"version,omitempty"`


	// Intents
	Intents *[]Intent `json:"intents,omitempty"`


	// EngineVersion
	EngineVersion *string `json:"engineVersion,omitempty"`

}

func (o *Nluinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluinfo
	
	return json.Marshal(&struct { 
		Domain *Addressableentityref `json:"domain,omitempty"`
		
		Version **Nludomainversion `json:"version,omitempty"`
		
		Intents *[]Intent `json:"intents,omitempty"`
		
		EngineVersion *string `json:"engineVersion,omitempty"`
		*Alias
	}{ 
		Domain: o.Domain,
		
		Version: o.Version,
		
		Intents: o.Intents,
		
		EngineVersion: o.EngineVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluinfo) UnmarshalJSON(b []byte) error {
	var NluinfoMap map[string]interface{}
	err := json.Unmarshal(b, &NluinfoMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := NluinfoMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	
	if Version, ok := NluinfoMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if Intents, ok := NluinfoMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if EngineVersion, ok := NluinfoMap["engineVersion"].(string); ok {
		o.EngineVersion = &EngineVersion
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
