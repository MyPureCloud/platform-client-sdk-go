package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestaterequestitem
type Bulkupdateshifttradestaterequestitem struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// State - The new state to set on the shift trade
	State *string `json:"state,omitempty"`


	// Metadata - Version metadata for the shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestaterequestitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
