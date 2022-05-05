package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentcarousel - Carousel content object.
type Contentcarousel struct { 
	// Cards - An array of card objects.
	Cards *[]Contentcard `json:"cards,omitempty"`

}

func (o *Contentcarousel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentcarousel
	
	return json.Marshal(&struct { 
		Cards *[]Contentcard `json:"cards,omitempty"`
		*Alias
	}{ 
		Cards: o.Cards,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentcarousel) UnmarshalJSON(b []byte) error {
	var ContentcarouselMap map[string]interface{}
	err := json.Unmarshal(b, &ContentcarouselMap)
	if err != nil {
		return err
	}
	
	if Cards, ok := ContentcarouselMap["cards"].([]interface{}); ok {
		CardsString, _ := json.Marshal(Cards)
		json.Unmarshal(CardsString, &o.Cards)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentcarousel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
