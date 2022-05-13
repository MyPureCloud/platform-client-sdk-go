package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Closebuttonstyleproperties
type Closebuttonstyleproperties struct { 
	// Color - Color of button. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// Opacity - Opacity of button.
	Opacity *float32 `json:"opacity,omitempty"`

}

func (o *Closebuttonstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Closebuttonstyleproperties
	
	return json.Marshal(&struct { 
		Color *string `json:"color,omitempty"`
		
		Opacity *float32 `json:"opacity,omitempty"`
		*Alias
	}{ 
		Color: o.Color,
		
		Opacity: o.Opacity,
		Alias:    (*Alias)(o),
	})
}

func (o *Closebuttonstyleproperties) UnmarshalJSON(b []byte) error {
	var ClosebuttonstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ClosebuttonstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Color, ok := ClosebuttonstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if Opacity, ok := ClosebuttonstylepropertiesMap["opacity"].(float64); ok {
		OpacityFloat32 := float32(Opacity)
		o.Opacity = &OpacityFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Closebuttonstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
