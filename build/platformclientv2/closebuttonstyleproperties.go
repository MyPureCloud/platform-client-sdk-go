package platformclientv2
import (
	"encoding/json"
)

// Closebuttonstyleproperties
type Closebuttonstyleproperties struct { 
	// Color - Color of button. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// Opacity - Opacity of button.
	Opacity *float32 `json:"opacity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Closebuttonstyleproperties) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
