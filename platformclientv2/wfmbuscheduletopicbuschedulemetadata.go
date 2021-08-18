package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Wfmbuscheduletopicbuschedulemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulemetadata

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Wfmbuscheduletopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		ManagementUnits *[]Wfmbuscheduletopicbumanagementunitschedulesummary `json:"managementUnits,omitempty"`
		
		GenerationResults *Wfmbuscheduletopicbuschedulegenerationresultsummary `json:"generationResults,omitempty"`
		
		Metadata *Wfmbuscheduletopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		WeekCount: u.WeekCount,
		
		Description: u.Description,
		
		Published: u.Published,
		
		ShortTermForecast: u.ShortTermForecast,
		
		ManagementUnits: u.ManagementUnits,
		
		GenerationResults: u.GenerationResults,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
