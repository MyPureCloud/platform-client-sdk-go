package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulemetadata
type Wfmbuscheduletopicbuschedulemetadata struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *time.Time `json:"weekDate,omitempty"`


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

func (o *Wfmbuscheduletopicbuschedulemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulemetadata
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Wfmbuscheduletopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		ManagementUnits *[]Wfmbuscheduletopicbumanagementunitschedulesummary `json:"managementUnits,omitempty"`
		
		GenerationResults *Wfmbuscheduletopicbuschedulegenerationresultsummary `json:"generationResults,omitempty"`
		
		Metadata *Wfmbuscheduletopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: o.WeekCount,
		
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		ManagementUnits: o.ManagementUnits,
		
		GenerationResults: o.GenerationResults,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbuschedulemetadata) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbuschedulemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbuschedulemetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduletopicbuschedulemetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := WfmbuscheduletopicbuschedulemetadataMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if WeekCount, ok := WfmbuscheduletopicbuschedulemetadataMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Description, ok := WfmbuscheduletopicbuschedulemetadataMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := WfmbuscheduletopicbuschedulemetadataMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ShortTermForecast, ok := WfmbuscheduletopicbuschedulemetadataMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if ManagementUnits, ok := WfmbuscheduletopicbuschedulemetadataMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	
	if GenerationResults, ok := WfmbuscheduletopicbuschedulemetadataMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if Metadata, ok := WfmbuscheduletopicbuschedulemetadataMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
