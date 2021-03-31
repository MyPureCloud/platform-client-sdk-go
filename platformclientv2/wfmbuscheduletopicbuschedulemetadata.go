package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulemetadata
type Wfmbuscheduletopicbuschedulemetadata struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast
	ShortTermForecast *Wfmbuscheduletopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// ManagementUnits
	ManagementUnits *[]Wfmbuscheduletopicbumanagementunitschedulesummary `json:"managementUnits,omitempty"`


	// GenerationResults
	GenerationResults *Wfmbuscheduletopicbuschedulegenerationresultsummary `json:"generationResults,omitempty"`


	// Metadata
	Metadata *Wfmbuscheduletopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
