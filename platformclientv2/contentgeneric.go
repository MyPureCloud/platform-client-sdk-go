package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentgeneric - Generic content object.
type Contentgeneric struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Image - URL of an image.
	Image *string `json:"image,omitempty"`


	// Video - URL of a video.
	Video *string `json:"video,omitempty"`


	// Actions - Actions to be taken.
	Actions *Contentactions `json:"actions,omitempty"`


	// Components - An array of component objects.
	Components *[]Buttoncomponent `json:"components,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
