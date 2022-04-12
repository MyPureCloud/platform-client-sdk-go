package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterstylesetting
type Supportcenterstylesetting struct { 
	// HeroStyle - Support center hero customizations
	HeroStyle *Supportcenterherostyle `json:"heroStyle,omitempty"`


	// GlobalStyle - Support center global customizations
	GlobalStyle *Supportcenterglobalstyle `json:"globalStyle,omitempty"`

}

func (o *Supportcenterstylesetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcenterstylesetting
	
	return json.Marshal(&struct { 
		HeroStyle *Supportcenterherostyle `json:"heroStyle,omitempty"`
		
		GlobalStyle *Supportcenterglobalstyle `json:"globalStyle,omitempty"`
		*Alias
	}{ 
		HeroStyle: o.HeroStyle,
		
		GlobalStyle: o.GlobalStyle,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcenterstylesetting) UnmarshalJSON(b []byte) error {
	var SupportcenterstylesettingMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterstylesettingMap)
	if err != nil {
		return err
	}
	
	if HeroStyle, ok := SupportcenterstylesettingMap["heroStyle"].(map[string]interface{}); ok {
		HeroStyleString, _ := json.Marshal(HeroStyle)
		json.Unmarshal(HeroStyleString, &o.HeroStyle)
	}
	
	if GlobalStyle, ok := SupportcenterstylesettingMap["globalStyle"].(map[string]interface{}); ok {
		GlobalStyleString, _ := json.Marshal(GlobalStyle)
		json.Unmarshal(GlobalStyleString, &o.GlobalStyle)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterstylesetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
