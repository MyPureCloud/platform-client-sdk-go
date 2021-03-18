package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicaddress
type Queueconversationsocialexpressioneventtopicaddress struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// NameRaw
	NameRaw *string `json:"nameRaw,omitempty"`


	// AddressNormalized
	AddressNormalized *string `json:"addressNormalized,omitempty"`


	// AddressRaw
	AddressRaw *string `json:"addressRaw,omitempty"`


	// AddressDisplayable
	AddressDisplayable *string `json:"addressDisplayable,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
