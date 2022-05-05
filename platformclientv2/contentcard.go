package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentcard - Card content object.
type Contentcard struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Image - URL of an image.
	Image *string `json:"image,omitempty"`


	// Video - URL of a video.
	Video *string `json:"video,omitempty"`


	// DefaultAction - The default button action.
	DefaultAction *Contentcardaction `json:"defaultAction,omitempty"`


	// Actions - An array of action objects.
	Actions *[]Contentcardaction `json:"actions,omitempty"`

}

func (o *Contentcard) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentcard
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		DefaultAction *Contentcardaction `json:"defaultAction,omitempty"`
		
		Actions *[]Contentcardaction `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Image: o.Image,
		
		Video: o.Video,
		
		DefaultAction: o.DefaultAction,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentcard) UnmarshalJSON(b []byte) error {
	var ContentcardMap map[string]interface{}
	err := json.Unmarshal(b, &ContentcardMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ContentcardMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := ContentcardMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Image, ok := ContentcardMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if Video, ok := ContentcardMap["video"].(string); ok {
		o.Video = &Video
	}
	
	if DefaultAction, ok := ContentcardMap["defaultAction"].(map[string]interface{}); ok {
		DefaultActionString, _ := json.Marshal(DefaultAction)
		json.Unmarshal(DefaultActionString, &o.DefaultAction)
	}
	
	if Actions, ok := ContentcardMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentcard) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
