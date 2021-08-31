package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateservicegoaltemplate
type Updateservicegoaltemplate struct { 
	// Name - The name of the service goal template.
	Name *string `json:"name,omitempty"`


	// ServiceLevel - Service level targets for this service goal template
	ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`


	// AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
	AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`


	// AbandonRate - Abandon rate targets for this service goal template
	AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`


	// Metadata - Version metadata for the service goal template
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updateservicegoaltemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateservicegoaltemplate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		ServiceLevel: o.ServiceLevel,
		
		AverageSpeedOfAnswer: o.AverageSpeedOfAnswer,
		
		AbandonRate: o.AbandonRate,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateservicegoaltemplate) UnmarshalJSON(b []byte) error {
	var UpdateservicegoaltemplateMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateservicegoaltemplateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateservicegoaltemplateMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ServiceLevel, ok := UpdateservicegoaltemplateMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if AverageSpeedOfAnswer, ok := UpdateservicegoaltemplateMap["averageSpeedOfAnswer"].(map[string]interface{}); ok {
		AverageSpeedOfAnswerString, _ := json.Marshal(AverageSpeedOfAnswer)
		json.Unmarshal(AverageSpeedOfAnswerString, &o.AverageSpeedOfAnswer)
	}
	
	if AbandonRate, ok := UpdateservicegoaltemplateMap["abandonRate"].(map[string]interface{}); ok {
		AbandonRateString, _ := json.Marshal(AbandonRate)
		json.Unmarshal(AbandonRateString, &o.AbandonRate)
	}
	
	if Metadata, ok := UpdateservicegoaltemplateMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateservicegoaltemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
