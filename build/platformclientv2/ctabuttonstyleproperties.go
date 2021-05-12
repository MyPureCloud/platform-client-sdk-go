package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Ctabuttonstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
