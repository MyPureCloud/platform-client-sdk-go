package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterherostyle
type Supportcenterherostyle struct { 
	// BackgroundColor - Background color for hero section, in hexadecimal format, eg #ffffff
	BackgroundColor *string `json:"backgroundColor,omitempty"`


	// TextColor - Text color for hero section, in hexadecimal format, eg #ffffff
	TextColor *string `json:"textColor,omitempty"`

}

func (o *Supportcenterherostyle) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcenterherostyle
	
	return json.Marshal(&struct { 
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		TextColor *string `json:"textColor,omitempty"`
		*Alias
	}{ 
		BackgroundColor: o.BackgroundColor,
		
		TextColor: o.TextColor,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcenterherostyle) UnmarshalJSON(b []byte) error {
	var SupportcenterherostyleMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterherostyleMap)
	if err != nil {
		return err
	}
	
	if BackgroundColor, ok := SupportcenterherostyleMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
	
	if TextColor, ok := SupportcenterherostyleMap["textColor"].(string); ok {
		o.TextColor = &TextColor
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterherostyle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
