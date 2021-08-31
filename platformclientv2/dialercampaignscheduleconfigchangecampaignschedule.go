package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignscheduleconfigchangecampaignschedule
type Dialercampaignscheduleconfigchangecampaignschedule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// Intervals
	Intervals *[]Dialercampaignscheduleconfigchangescheduleinterval `json:"intervals,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// Campaign
	Campaign *Dialercampaignscheduleconfigchangeurireference `json:"campaign,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercampaignscheduleconfigchangecampaignschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignscheduleconfigchangecampaignschedule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Intervals *[]Dialercampaignscheduleconfigchangescheduleinterval `json:"intervals,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Campaign *Dialercampaignscheduleconfigchangeurireference `json:"campaign,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Intervals: o.Intervals,
		
		TimeZone: o.TimeZone,
		
		Campaign: o.Campaign,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignscheduleconfigchangecampaignschedule) UnmarshalJSON(b []byte) error {
	var DialercampaignscheduleconfigchangecampaignscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignscheduleconfigchangecampaignscheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Intervals, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["intervals"].([]interface{}); ok {
		IntervalsString, _ := json.Marshal(Intervals)
		json.Unmarshal(IntervalsString, &o.Intervals)
	}
	
	if TimeZone, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if Campaign, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if AdditionalProperties, ok := DialercampaignscheduleconfigchangecampaignscheduleMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangecampaignschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
