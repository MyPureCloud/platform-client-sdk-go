package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulelistitemresponse
type Weekschedulelistitemresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - Description of the week schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether the week schedule is published
	Published *bool `json:"published,omitempty"`


	// GenerationResults - Summary of the results from the schedule run
	GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`


	// ShortTermForecast - Short term forecast associated with this schedule
	ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Weekschedulelistitemresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedulelistitemresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`
		
		ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		WeekDate: o.WeekDate,
		
		Description: o.Description,
		
		Published: o.Published,
		
		GenerationResults: o.GenerationResults,
		
		ShortTermForecast: o.ShortTermForecast,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekschedulelistitemresponse) UnmarshalJSON(b []byte) error {
	var WeekschedulelistitemresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulelistitemresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WeekschedulelistitemresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := WeekschedulelistitemresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if WeekDate, ok := WeekschedulelistitemresponseMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if Description, ok := WeekschedulelistitemresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := WeekschedulelistitemresponseMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if GenerationResults, ok := WeekschedulelistitemresponseMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if ShortTermForecast, ok := WeekschedulelistitemresponseMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if Metadata, ok := WeekschedulelistitemresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulelistitemresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
