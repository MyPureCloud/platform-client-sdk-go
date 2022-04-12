package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentcard
type Conversationcontentcard struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// DefaultAction - Default action to be taken.
	DefaultAction *Conversationcardaction `json:"defaultAction,omitempty"`


	// Actions - A List of action objects.
	Actions *[]Conversationcardaction `json:"actions,omitempty"`


	// Image
	Image *string `json:"image,omitempty"`


	// Video
	Video *string `json:"video,omitempty"`

}

func (o *Conversationcontentcard) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentcard
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultAction *Conversationcardaction `json:"defaultAction,omitempty"`
		
		Actions *[]Conversationcardaction `json:"actions,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		DefaultAction: o.DefaultAction,
		
		Actions: o.Actions,
		
		Image: o.Image,
		
		Video: o.Video,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentcard) UnmarshalJSON(b []byte) error {
	var ConversationcontentcardMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentcardMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ConversationcontentcardMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := ConversationcontentcardMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if DefaultAction, ok := ConversationcontentcardMap["defaultAction"].(map[string]interface{}); ok {
		DefaultActionString, _ := json.Marshal(DefaultAction)
		json.Unmarshal(DefaultActionString, &o.DefaultAction)
	}
	
	if Actions, ok := ConversationcontentcardMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Image, ok := ConversationcontentcardMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if Video, ok := ConversationcontentcardMap["video"].(string); ok {
		o.Video = &Video
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentcard) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
