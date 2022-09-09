package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentcarousel - Carousel content object.
type Conversationcontentcarousel struct { 
	// Cards - An array of card objects.
	Cards *[]Conversationcontentcard `json:"cards,omitempty"`

}

func (o *Conversationcontentcarousel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentcarousel
	
	return json.Marshal(&struct { 
		Cards *[]Conversationcontentcard `json:"cards,omitempty"`
		*Alias
	}{ 
		Cards: o.Cards,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentcarousel) UnmarshalJSON(b []byte) error {
	var ConversationcontentcarouselMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentcarouselMap)
	if err != nil {
		return err
	}
	
	if Cards, ok := ConversationcontentcarouselMap["cards"].([]interface{}); ok {
		CardsString, _ := json.Marshal(Cards)
		json.Unmarshal(CardsString, &o.Cards)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentcarousel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
