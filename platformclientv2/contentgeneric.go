package platformclientv2
import (
	"encoding/json"
)

// Contentgeneric - Generic content object
type Contentgeneric struct { 
	// Id - An ID assigned to this rich message content. Each instance inside the content array has a unique ID.
	Id *string `json:"id,omitempty"`


	// Title - Text to show in the title row
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description row. This is immediately below the title
	Description *string `json:"description,omitempty"`


	// Image - Path or URI to an image file
	Image *string `json:"image,omitempty"`


	// Video - Path or URI to a video file
	Video *string `json:"video,omitempty"`


	// Actions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
	Actions *Contentactions `json:"actions,omitempty"`


	// Components - An array of component objects
	Components *[]Buttoncomponent `json:"components,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentgeneric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
