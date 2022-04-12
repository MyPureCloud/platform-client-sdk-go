package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterglobalstyle
type Supportcenterglobalstyle struct { 
	// BackgroundColor - Global background color, in hexadecimal format, eg #ffffff
	BackgroundColor *string `json:"backgroundColor,omitempty"`


	// PrimaryColor - Global primary color, in hexadecimal format, eg #ffffff
	PrimaryColor *string `json:"primaryColor,omitempty"`


	// PrimaryColorDark - Global dark primary color, in hexadecimal format, eg #ffffff
	PrimaryColorDark *string `json:"primaryColorDark,omitempty"`


	// PrimaryColorLight - Global light primary color, in hexadecimal format, eg #ffffff
	PrimaryColorLight *string `json:"primaryColorLight,omitempty"`


	// TextColor - Global text color, in hexadecimal format, eg #ffffff
	TextColor *string `json:"textColor,omitempty"`


	// FontFamily - Global font family
	FontFamily *string `json:"fontFamily,omitempty"`

}

func (o *Supportcenterglobalstyle) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcenterglobalstyle
	
	return json.Marshal(&struct { 
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		PrimaryColor *string `json:"primaryColor,omitempty"`
		
		PrimaryColorDark *string `json:"primaryColorDark,omitempty"`
		
		PrimaryColorLight *string `json:"primaryColorLight,omitempty"`
		
		TextColor *string `json:"textColor,omitempty"`
		
		FontFamily *string `json:"fontFamily,omitempty"`
		*Alias
	}{ 
		BackgroundColor: o.BackgroundColor,
		
		PrimaryColor: o.PrimaryColor,
		
		PrimaryColorDark: o.PrimaryColorDark,
		
		PrimaryColorLight: o.PrimaryColorLight,
		
		TextColor: o.TextColor,
		
		FontFamily: o.FontFamily,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcenterglobalstyle) UnmarshalJSON(b []byte) error {
	var SupportcenterglobalstyleMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterglobalstyleMap)
	if err != nil {
		return err
	}
	
	if BackgroundColor, ok := SupportcenterglobalstyleMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
	
	if PrimaryColor, ok := SupportcenterglobalstyleMap["primaryColor"].(string); ok {
		o.PrimaryColor = &PrimaryColor
	}
	
	if PrimaryColorDark, ok := SupportcenterglobalstyleMap["primaryColorDark"].(string); ok {
		o.PrimaryColorDark = &PrimaryColorDark
	}
	
	if PrimaryColorLight, ok := SupportcenterglobalstyleMap["primaryColorLight"].(string); ok {
		o.PrimaryColorLight = &PrimaryColorLight
	}
	
	if TextColor, ok := SupportcenterglobalstyleMap["textColor"].(string); ok {
		o.TextColor = &TextColor
	}
	
	if FontFamily, ok := SupportcenterglobalstyleMap["fontFamily"].(string); ok {
		o.FontFamily = &FontFamily
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterglobalstyle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
