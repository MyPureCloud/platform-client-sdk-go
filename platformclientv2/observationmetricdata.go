package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Observationmetricdata
type Observationmetricdata struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`


	// Truncated - Flag for a truncated list of observations. If truncated, the first half of the list of observations will contain the oldest observations and the second half the newest observations.
	Truncated *bool `json:"truncated,omitempty"`


	// Observations - List of observations sorted by timestamp in ascending order. This list may be truncated.
	Observations *[]Observationvalue `json:"observations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Observationmetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
