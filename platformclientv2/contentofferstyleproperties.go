package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentofferstyleproperties
type Contentofferstyleproperties struct { 
	// Padding - Padding of the offer. (eg. 10px)
	Padding *string `json:"padding,omitempty"`


	// Color - Text color of the offer. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// BackgroundColor - Background color of the offer. (eg. #000000)
	BackgroundColor *string `json:"backgroundColor,omitempty"`

}

func (o *Contentofferstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentofferstyleproperties
	
	return json.Marshal(&struct { 
		Padding *string `json:"padding,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		*Alias
	}{ 
		Padding: o.Padding,
		
		Color: o.Color,
		
		BackgroundColor: o.BackgroundColor,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentofferstyleproperties) UnmarshalJSON(b []byte) error {
	var ContentofferstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ContentofferstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Padding, ok := ContentofferstylepropertiesMap["padding"].(string); ok {
		o.Padding = &Padding
	}
    
	if Color, ok := ContentofferstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if BackgroundColor, ok := ContentofferstylepropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentofferstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
