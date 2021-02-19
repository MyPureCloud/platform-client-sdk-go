package platformclientv2
import (
	"encoding/json"
)

// Patchtextstyleproperties
type Patchtextstyleproperties struct { 
	// Color - Color of the text. (eg. #FFFFFF)
	Color *string `json:"color,omitempty"`


	// Font - Font of the text. (eg. Helvetica)
	Font *string `json:"font,omitempty"`


	// FontSize - Font size of the text. (eg. '12')
	FontSize *string `json:"fontSize,omitempty"`


	// TextAlign - Text alignment.
	TextAlign *string `json:"textAlign,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchtextstyleproperties) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
