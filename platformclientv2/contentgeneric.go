package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Contentgeneric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentgeneric

	

	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		
		Components *[]Buttoncomponent `json:"components,omitempty"`
		*Alias
	}{ 
		Title: u.Title,
		
		Description: u.Description,
		
		Image: u.Image,
		
		Video: u.Video,
		
		Actions: u.Actions,
		
		Components: u.Components,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
