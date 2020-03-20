package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduletopicbushorttermforecastreference
type Wfmbuscheduletopicbushorttermforecastreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbushorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
