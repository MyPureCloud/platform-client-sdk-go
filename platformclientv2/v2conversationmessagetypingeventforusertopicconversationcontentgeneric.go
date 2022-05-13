package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentgeneric
type V2conversationmessagetypingeventforusertopicconversationcontentgeneric struct { 
	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Image
	Image *string `json:"image,omitempty"`


	// Video
	Video *string `json:"video,omitempty"`


	// Actions
	Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`


	// Components
	Components *[]V2conversationmessagetypingeventforusertopicconversationbuttoncomponent `json:"components,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentgeneric
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`
		
		Components *[]V2conversationmessagetypingeventforusertopicconversationbuttoncomponent `json:"components,omitempty"`
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

func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentgenericMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentgenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Image, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Video, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["video"].(string); ok {
		o.Video = &Video
	}
    
	if Actions, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
