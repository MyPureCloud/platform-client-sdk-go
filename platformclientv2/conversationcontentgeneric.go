package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentgeneric - Generic content object (Deprecated).
type Conversationcontentgeneric struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Image - URL of an image.
	Image *string `json:"image,omitempty"`


	// Video - URL of a video.
	Video *string `json:"video,omitempty"`


	// Actions - Actions to be taken.
	Actions *Conversationcontentactions `json:"actions,omitempty"`


	// Components - An array of component objects.
	Components *[]Conversationbuttoncomponent `json:"components,omitempty"`

}

func (o *Conversationcontentgeneric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentgeneric
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		Actions *Conversationcontentactions `json:"actions,omitempty"`
		
		Components *[]Conversationbuttoncomponent `json:"components,omitempty"`
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

func (o *Conversationcontentgeneric) UnmarshalJSON(b []byte) error {
	var ConversationcontentgenericMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentgenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ConversationcontentgenericMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ConversationcontentgenericMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Image, ok := ConversationcontentgenericMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Video, ok := ConversationcontentgenericMap["video"].(string); ok {
		o.Video = &Video
	}
    
	if Actions, ok := ConversationcontentgenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := ConversationcontentgenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
