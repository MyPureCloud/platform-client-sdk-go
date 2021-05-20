package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingofferfields
type Webmessagingofferfields struct { 
	// OfferText - Text value to be used when inviting a visitor to engage with a web messaging offer.
	OfferText *string `json:"offerText,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webmessagingofferfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
