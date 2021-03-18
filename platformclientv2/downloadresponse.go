package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Downloadresponse
type Downloadresponse struct { 
	// ContentLocationUri
	ContentLocationUri *string `json:"contentLocationUri,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Downloadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
