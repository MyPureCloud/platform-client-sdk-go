package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importscheduleuploadschema
type Importscheduleuploadschema struct { 
	// Description - The description for the imported schedule
	Description *string `json:"description,omitempty"`


	// WeekCount - The number of weeks the imported schedule will cover
	WeekCount *int `json:"weekCount,omitempty"`


	// Published - Whether the imported schedule should be immediately published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecast - The short term forecast to associate with the imported schedule
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// HeadcountForecast - The headcount forecast to associate with the imported schedule
	HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`


	// AgentSchedules - Individual agent schedules
	AgentSchedules *[]Buimportagentscheduleuploadschema `json:"agentSchedules,omitempty"`

}

func (o *Importscheduleuploadschema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importscheduleuploadschema
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`
		
		AgentSchedules *[]Buimportagentscheduleuploadschema `json:"agentSchedules,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		WeekCount: o.WeekCount,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		HeadcountForecast: o.HeadcountForecast,
		
		AgentSchedules: o.AgentSchedules,
		Alias:    (*Alias)(o),
	})
}

func (o *Importscheduleuploadschema) UnmarshalJSON(b []byte) error {
	var ImportscheduleuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &ImportscheduleuploadschemaMap)
	if err != nil {
		return err
	}
	
	if Description, ok := ImportscheduleuploadschemaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if WeekCount, ok := ImportscheduleuploadschemaMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Published, ok := ImportscheduleuploadschemaMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ShortTermForecast, ok := ImportscheduleuploadschemaMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if HeadcountForecast, ok := ImportscheduleuploadschemaMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if AgentSchedules, ok := ImportscheduleuploadschemaMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Importscheduleuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
