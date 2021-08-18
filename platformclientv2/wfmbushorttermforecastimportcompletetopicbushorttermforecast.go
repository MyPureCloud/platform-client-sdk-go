package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbushorttermforecast
type Wfmbushorttermforecastimportcompletetopicbushorttermforecast struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// CreationMethod
	CreationMethod *string `json:"creationMethod,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Legacy
	Legacy *bool `json:"legacy,omitempty"`


	// ReferenceStartDate
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// SourceDays
	SourceDays *[]Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastimportcompletetopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`


	// CanUseForScheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

}

func (u *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastimportcompletetopicbushorttermforecast

	
	ReferenceStartDate := new(string)
	if u.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(u.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		CreationMethod *string `json:"creationMethod,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Legacy *bool `json:"legacy,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		SourceDays *[]Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Wfmbushorttermforecastimportcompletetopicbuforecastmodification `json:"modifications,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Metadata *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		WeekDate: u.WeekDate,
		
		CreationMethod: u.CreationMethod,
		
		Description: u.Description,
		
		Legacy: u.Legacy,
		
		ReferenceStartDate: ReferenceStartDate,
		
		SourceDays: u.SourceDays,
		
		Modifications: u.Modifications,
		
		TimeZone: u.TimeZone,
		
		PlanningGroupsVersion: u.PlanningGroupsVersion,
		
		WeekCount: u.WeekCount,
		
		Metadata: u.Metadata,
		
		CanUseForScheduling: u.CanUseForScheduling,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
