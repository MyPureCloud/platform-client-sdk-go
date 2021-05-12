package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentlocation - Location object
type Contentlocation struct { 
	// Url - Location map url
	Url *string `json:"url,omitempty"`


	// Address - Location postal address
	Address *string `json:"address,omitempty"`


	// Text - Location name
	Text *string `json:"text,omitempty"`


	// Latitude - Latitude of the location
	Latitude *float64 `json:"latitude,omitempty"`


	// Longitude - Longitude of the location
	Longitude *float64 `json:"longitude,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
