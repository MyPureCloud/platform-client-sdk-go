package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulemetadata
type Buschedulemetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks spanned by this schedule
	WeekCount *int `json:"weekCount,omitempty"`


	// Description - The description of this schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether this schedule is published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast - The forecast used for this schedule, if applicable
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// GenerationResults - Generation result summary for this schedule, if applicable
	GenerationResults *Schedulegenerationresultsummary `json:"generationResults,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// ManagementUnits - High level per-management unit schedule metadata
	ManagementUnits *[]Bumanagementunitschedulesummary `json:"managementUnits,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Buschedulemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulemetadata

	
	WeekDate := new(string)
	if u.WeekDate != nil {
		*WeekDate = timeutil.Strftime(u.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		GenerationResults *Schedulegenerationresultsummary `json:"generationResults,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		ManagementUnits *[]Bumanagementunitschedulesummary `json:"managementUnits,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: u.WeekCount,
		
		Description: u.Description,
		
		Published: u.Published,
		
		ShortTermForecast: u.ShortTermForecast,
		
		GenerationResults: u.GenerationResults,
		
		Metadata: u.Metadata,
		
		ManagementUnits: u.ManagementUnits,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
