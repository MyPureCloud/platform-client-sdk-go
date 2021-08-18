package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bushorttermforecast
type Bushorttermforecast struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks this forecast covers
	WeekCount *int `json:"weekCount,omitempty"`


	// CreationMethod - The method by which this forecast was created
	CreationMethod *string `json:"creationMethod,omitempty"`


	// Description - The description of this forecast
	Description *string `json:"description,omitempty"`


	// Legacy - Whether this forecast contains modifications on legacy metrics
	Legacy *bool `json:"legacy,omitempty"`


	// Metadata - Metadata for this forecast
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// CanUseForScheduling - Whether this forecast can be used for scheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`


	// ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// SourceDays - The source day pointers for this forecast
	SourceDays *[]Forecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications - Any manual modifications applied to this forecast
	Modifications *[]Buforecastmodification `json:"modifications,omitempty"`


	// GenerationResults - Generation result metadata
	GenerationResults *Buforecastgenerationresult `json:"generationResults,omitempty"`


	// TimeZone - The time zone for this forecast
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion - The version of the planning groups that was used for this forecast
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// PlanningGroups - A snapshot of the planning groups used for this forecast as of the version number indicated
	PlanningGroups *Forecastplanninggroupsresponse `json:"planningGroups,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Bushorttermforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bushorttermforecast

	
	WeekDate := new(string)
	if u.WeekDate != nil {
		*WeekDate = timeutil.Strftime(u.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	ReferenceStartDate := new(string)
	if u.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(u.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		CreationMethod *string `json:"creationMethod,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Legacy *bool `json:"legacy,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		SourceDays *[]Forecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Buforecastmodification `json:"modifications,omitempty"`
		
		GenerationResults *Buforecastgenerationresult `json:"generationResults,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		PlanningGroups *Forecastplanninggroupsresponse `json:"planningGroups,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: u.WeekCount,
		
		CreationMethod: u.CreationMethod,
		
		Description: u.Description,
		
		Legacy: u.Legacy,
		
		Metadata: u.Metadata,
		
		CanUseForScheduling: u.CanUseForScheduling,
		
		ReferenceStartDate: ReferenceStartDate,
		
		SourceDays: u.SourceDays,
		
		Modifications: u.Modifications,
		
		GenerationResults: u.GenerationResults,
		
		TimeZone: u.TimeZone,
		
		PlanningGroupsVersion: u.PlanningGroupsVersion,
		
		PlanningGroups: u.PlanningGroups,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
