package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast
type Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast struct { 
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
	SourceDays *[]Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer `json:"sourceDays,omitempty"`


	// Modifications
	Modifications *[]Wfmbushorttermforecastgenerateprogresstopicbuforecastmodification `json:"modifications,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Metadata
	Metadata *Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata `json:"metadata,omitempty"`


	// CanUseForScheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

}

func (o *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		SourceDays *[]Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Wfmbushorttermforecastgenerateprogresstopicbuforecastmodification `json:"modifications,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Metadata *Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: o.WeekDate,
		
		CreationMethod: o.CreationMethod,
		
		Description: o.Description,
		
		Legacy: o.Legacy,
		
		ReferenceStartDate: ReferenceStartDate,
		
		SourceDays: o.SourceDays,
		
		Modifications: o.Modifications,
		
		TimeZone: o.TimeZone,
		
		PlanningGroupsVersion: o.PlanningGroupsVersion,
		
		WeekCount: o.WeekCount,
		
		Metadata: o.Metadata,
		
		CanUseForScheduling: o.CanUseForScheduling,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if WeekDate, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if CreationMethod, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["creationMethod"].(string); ok {
		o.CreationMethod = &CreationMethod
	}
	
	if Description, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Legacy, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["legacy"].(bool); ok {
		o.Legacy = &Legacy
	}
	
	if referenceStartDateString, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if SourceDays, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["sourceDays"].([]interface{}); ok {
		SourceDaysString, _ := json.Marshal(SourceDays)
		json.Unmarshal(SourceDaysString, &o.SourceDays)
	}
	
	if Modifications, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["modifications"].([]interface{}); ok {
		ModificationsString, _ := json.Marshal(Modifications)
		json.Unmarshal(ModificationsString, &o.Modifications)
	}
	
	if TimeZone, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if PlanningGroupsVersion, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	
	if WeekCount, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Metadata, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if CanUseForScheduling, ok := WfmbushorttermforecastgenerateprogresstopicbushorttermforecastMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
