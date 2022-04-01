package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric
type V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric struct { 
	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Image
	Image *string `json:"image,omitempty"`


	// Video
	Video *string `json:"video,omitempty"`


	// Actions
	Actions *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions `json:"actions,omitempty"`


	// Components
	Components *[]V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent `json:"components,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		Actions *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions `json:"actions,omitempty"`
		
		Components *[]V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent `json:"components,omitempty"`
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

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Image, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if Video, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["video"].(string); ok {
		o.Video = &Video
	}
	
	if Actions, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentgenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
