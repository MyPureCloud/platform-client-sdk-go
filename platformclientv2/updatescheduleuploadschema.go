package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatescheduleuploadschema
type Updatescheduleuploadschema struct { 
	// Description - The description to set for the schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether to publish the schedule. Note: a schedule cannot be un-published unless another schedule is published over it
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast - The short term forecast to associate with the schedule
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// HeadcountForecast - The headcount forecast to associate with the schedule
	HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`


	// AgentSchedules - Individual agent schedules
	AgentSchedules *[]Buupdateagentscheduleuploadschema `json:"agentSchedules,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updatescheduleuploadschema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatescheduleuploadschema
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`
		
		AgentSchedules *[]Buupdateagentscheduleuploadschema `json:"agentSchedules,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		HeadcountForecast: o.HeadcountForecast,
		
		AgentSchedules: o.AgentSchedules,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatescheduleuploadschema) UnmarshalJSON(b []byte) error {
	var UpdatescheduleuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatescheduleuploadschemaMap)
	if err != nil {
		return err
	}
	
	if Description, ok := UpdatescheduleuploadschemaMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Published, ok := UpdatescheduleuploadschemaMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if ShortTermForecast, ok := UpdatescheduleuploadschemaMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if HeadcountForecast, ok := UpdatescheduleuploadschemaMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if AgentSchedules, ok := UpdatescheduleuploadschemaMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	
	if Metadata, ok := UpdatescheduleuploadschemaMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatescheduleuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
