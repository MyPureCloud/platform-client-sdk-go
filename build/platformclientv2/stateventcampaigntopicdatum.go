package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventcampaigntopicdatum
type Stateventcampaigntopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventcampaigntopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
