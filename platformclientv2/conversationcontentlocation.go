package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentlocation - Location object.
type Conversationcontentlocation struct { 
	// Url - URL of the Location.
	Url *string `json:"url,omitempty"`


	// Address - Location postal address.
	Address *string `json:"address,omitempty"`


	// Text - Location name.
	Text *string `json:"text,omitempty"`


	// Latitude - Latitude of the location.
	Latitude *float64 `json:"latitude,omitempty"`


	// Longitude - Longitude of the location.
	Longitude *float64 `json:"longitude,omitempty"`

}

func (o *Conversationcontentlocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentlocation
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Latitude *float64 `json:"latitude,omitempty"`
		
		Longitude *float64 `json:"longitude,omitempty"`
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

func (o *Conversationcontentlocation) UnmarshalJSON(b []byte) error {
	var ConversationcontentlocationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentlocationMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationcontentlocationMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Address, ok := ConversationcontentlocationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Text, ok := ConversationcontentlocationMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Latitude, ok := ConversationcontentlocationMap["latitude"].(float64); ok {
		o.Latitude = &Latitude
	}
    
	if Longitude, ok := ConversationcontentlocationMap["longitude"].(float64); ok {
		o.Longitude = &Longitude
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
