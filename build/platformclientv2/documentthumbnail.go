package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Documentthumbnail
type Documentthumbnail struct { 
	// Resolution
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Height
	Height *int `json:"height,omitempty"`


	// Width
	Width *int `json:"width,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentthumbnail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
