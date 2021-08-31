package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicegoaltemplate - Service Goal Template
type Servicegoaltemplate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ServiceLevel - Service level targets for this service goal template
	ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`


	// AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
	AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`


	// AbandonRate - Abandon rate targets for this service goal template
	AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`


	// Metadata - Version metadata for the service goal template
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Servicegoaltemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicegoaltemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ServiceLevel: o.ServiceLevel,
		
		AverageSpeedOfAnswer: o.AverageSpeedOfAnswer,
		
		AbandonRate: o.AbandonRate,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Servicegoaltemplate) UnmarshalJSON(b []byte) error {
	var ServicegoaltemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ServicegoaltemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ServicegoaltemplateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ServicegoaltemplateMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ServiceLevel, ok := ServicegoaltemplateMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if AverageSpeedOfAnswer, ok := ServicegoaltemplateMap["averageSpeedOfAnswer"].(map[string]interface{}); ok {
		AverageSpeedOfAnswerString, _ := json.Marshal(AverageSpeedOfAnswer)
		json.Unmarshal(AverageSpeedOfAnswerString, &o.AverageSpeedOfAnswer)
	}
	
	if AbandonRate, ok := ServicegoaltemplateMap["abandonRate"].(map[string]interface{}); ok {
		AbandonRateString, _ := json.Marshal(AbandonRate)
		json.Unmarshal(AbandonRateString, &o.AbandonRate)
	}
	
	if Metadata, ok := ServicegoaltemplateMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := ServicegoaltemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Servicegoaltemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
