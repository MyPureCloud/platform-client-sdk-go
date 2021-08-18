package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentlocation - Location object.
type Contentlocation struct { 
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

func (u *Contentlocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentlocation

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Latitude *float64 `json:"latitude,omitempty"`
		
		Longitude *float64 `json:"longitude,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		
		Address: u.Address,
		
		Text: u.Text,
		
		Latitude: u.Latitude,
		
		Longitude: u.Longitude,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
