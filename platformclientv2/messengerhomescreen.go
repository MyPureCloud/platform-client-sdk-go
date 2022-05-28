package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengerhomescreen
type Messengerhomescreen struct { 
	// Enabled - whether or not homescreen is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// LogoUrl - to capture uploaded company logoUrl
	LogoUrl *string `json:"logoUrl,omitempty"`

}

func (o *Messengerhomescreen) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengerhomescreen
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		LogoUrl *string `json:"logoUrl,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		LogoUrl: o.LogoUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Messengerhomescreen) UnmarshalJSON(b []byte) error {
	var MessengerhomescreenMap map[string]interface{}
	err := json.Unmarshal(b, &MessengerhomescreenMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := MessengerhomescreenMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if LogoUrl, ok := MessengerhomescreenMap["logoUrl"].(string); ok {
		o.LogoUrl = &LogoUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messengerhomescreen) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
