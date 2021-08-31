package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ctabuttonstyleproperties
type Ctabuttonstyleproperties struct { 
	// Color - Color of the text. (eg. #FFFFFF)
	Color *string `json:"color,omitempty"`


	// Font - Font of the text. (eg. Helvetica)
	Font *string `json:"font,omitempty"`


	// FontSize - Font size of the text. (eg. '12')
	FontSize *string `json:"fontSize,omitempty"`


	// TextAlign - Text alignment.
	TextAlign *string `json:"textAlign,omitempty"`


	// BackgroundColor - Background color of the CTA button. (eg. #FF0000)
	BackgroundColor *string `json:"backgroundColor,omitempty"`

}

func (o *Ctabuttonstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ctabuttonstyleproperties
	
	return json.Marshal(&struct { 
		Color *string `json:"color,omitempty"`
		
		Font *string `json:"font,omitempty"`
		
		FontSize *string `json:"fontSize,omitempty"`
		
		TextAlign *string `json:"textAlign,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		*Alias
	}{ 
		Color: o.Color,
		
		Font: o.Font,
		
		FontSize: o.FontSize,
		
		TextAlign: o.TextAlign,
		
		BackgroundColor: o.BackgroundColor,
		Alias:    (*Alias)(o),
	})
}

func (o *Ctabuttonstyleproperties) UnmarshalJSON(b []byte) error {
	var CtabuttonstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &CtabuttonstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Color, ok := CtabuttonstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
	
	if Font, ok := CtabuttonstylepropertiesMap["font"].(string); ok {
		o.Font = &Font
	}
	
	if FontSize, ok := CtabuttonstylepropertiesMap["fontSize"].(string); ok {
		o.FontSize = &FontSize
	}
	
	if TextAlign, ok := CtabuttonstylepropertiesMap["textAlign"].(string); ok {
		o.TextAlign = &TextAlign
	}
	
	if BackgroundColor, ok := CtabuttonstylepropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ctabuttonstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
