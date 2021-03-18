package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceresultwrapper
type Wfmhistoricaladherenceresultwrapper struct { 
	// EntityId - The operation ID of the historical adherence query
	EntityId *string `json:"entityId,omitempty"`


	// Data - The list of historical adherence query results
	Data *[]Historicaladherencequeryresult `json:"data,omitempty"`


	// LookupIdToSecondaryPresenceId - Map of secondary presence lookup ID to corresponding secondary presence ID
	LookupIdToSecondaryPresenceId *map[string]string `json:"lookupIdToSecondaryPresenceId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresultwrapper) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
