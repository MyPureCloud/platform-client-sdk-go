package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Searchshifttraderesponse
type Searchshifttraderesponse struct { 
	// Trade - A trade which matches search criteria
	Trade *Shifttraderesponse `json:"trade,omitempty"`


	// MatchingReceivingShiftIds - IDs of shifts which match the search criteria
	MatchingReceivingShiftIds *[]string `json:"matchingReceivingShiftIds,omitempty"`


	// Preview - A preview of what the shift trade would look like if matched
	Preview *Shifttradepreviewresponse `json:"preview,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
