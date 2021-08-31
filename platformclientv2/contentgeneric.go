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

func (o *Contentgeneric) MarshalJSON() ([]byte, error) {
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
		Title: o.Title,
		
		Description: o.Description,
		
		Image: o.Image,
		
		Video: o.Video,
		
		Actions: o.Actions,
		
		Components: o.Components,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentgeneric) UnmarshalJSON(b []byte) error {
	var ContentgenericMap map[string]interface{}
	err := json.Unmarshal(b, &ContentgenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ContentgenericMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := ContentgenericMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Image, ok := ContentgenericMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if Video, ok := ContentgenericMap["video"].(string); ok {
		o.Video = &Video
	}
	
	if Actions, ok := ContentgenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := ContentgenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
