package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulelistitem
type Buschedulelistitem struct { 
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Buschedulelistitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulelistitem
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: o.WeekCount,
		
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		GenerationResults: o.GenerationResults,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulelistitem) UnmarshalJSON(b []byte) error {
	var BuschedulelistitemMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulelistitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BuschedulelistitemMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if weekDateString, ok := BuschedulelistitemMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if WeekCount, ok := BuschedulelistitemMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Description, ok := BuschedulelistitemMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Published, ok := BuschedulelistitemMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if ShortTermForecast, ok := BuschedulelistitemMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if GenerationResults, ok := BuschedulelistitemMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if Metadata, ok := BuschedulelistitemMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := BuschedulelistitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulelistitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
