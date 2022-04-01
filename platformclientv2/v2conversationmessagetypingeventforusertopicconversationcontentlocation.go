package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentlocation
type V2conversationmessagetypingeventforusertopicconversationcontentlocation struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Latitude
	Latitude *float32 `json:"latitude,omitempty"`


	// Longitude
	Longitude *float32 `json:"longitude,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentlocation
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Latitude *float32 `json:"latitude,omitempty"`
		
		Longitude *float32 `json:"longitude,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Address: o.Address,
		
		Text: o.Text,
		
		Latitude: o.Latitude,
		
		Longitude: o.Longitude,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentlocationMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentlocationMap)
	if err != nil {
		return err
	}
	
	if Url, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Address, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Latitude, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["latitude"].(float64); ok {
		LatitudeFloat32 := float32(Latitude)
		o.Latitude = &LatitudeFloat32
	}
	
	if Longitude, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["longitude"].(float64); ok {
		LongitudeFloat32 := float32(Longitude)
		o.Longitude = &LongitudeFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
