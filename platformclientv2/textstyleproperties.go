package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textstyleproperties
type Textstyleproperties struct { 
	// Color - Color of the text. (eg. #FFFFFF)
	Color *string `json:"color,omitempty"`


	// Font - Font of the text. (eg. Helvetica)
	Font *string `json:"font,omitempty"`


	// FontSize - Font size of the text. (eg. '12')
	FontSize *string `json:"fontSize,omitempty"`


	// TextAlign - Text alignment.
	TextAlign *string `json:"textAlign,omitempty"`

}

func (o *Textstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textstyleproperties
	
	return json.Marshal(&struct { 
		Color *string `json:"color,omitempty"`
		
		Font *string `json:"font,omitempty"`
		
		FontSize *string `json:"fontSize,omitempty"`
		
		TextAlign *string `json:"textAlign,omitempty"`
		*Alias
	}{ 
		Color: o.Color,
		
		Font: o.Font,
		
		FontSize: o.FontSize,
		
		TextAlign: o.TextAlign,
		Alias:    (*Alias)(o),
	})
}

func (o *Textstyleproperties) UnmarshalJSON(b []byte) error {
	var TextstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &TextstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Color, ok := TextstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
	
	if Font, ok := TextstylepropertiesMap["font"].(string); ok {
		o.Font = &Font
	}
	
	if FontSize, ok := TextstylepropertiesMap["fontSize"].(string); ok {
		o.FontSize = &FontSize
	}
	
	if TextAlign, ok := TextstylepropertiesMap["textAlign"].(string); ok {
		o.TextAlign = &TextAlign
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
