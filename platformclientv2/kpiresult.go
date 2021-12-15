package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Kpiresult
type Kpiresult struct { 
	// KpiTotalOn - Absolute metric (in which the KPI is based) total for the interactions handled by predictive routing (GPR was on)
	KpiTotalOn *int `json:"kpiTotalOn,omitempty"`


	// KpiTotalOff - Absolute metric (in which the KPI is based) total for the interactions not routed by predictive routing (GPR was off)
	KpiTotalOff *int `json:"kpiTotalOff,omitempty"`


	// InteractionCountOn - Total interactions handled by predictive routing (GPR was on)
	InteractionCountOn *int `json:"interactionCountOn,omitempty"`


	// InteractionCountOff - Total interactions not routed by predictive routing (GPR was off)
	InteractionCountOff *int `json:"interactionCountOff,omitempty"`


	// MediaType - Media type used for the KPI
	MediaType *string `json:"mediaType,omitempty"`

}

func (o *Kpiresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Kpiresult
	
	return json.Marshal(&struct { 
		KpiTotalOn *int `json:"kpiTotalOn,omitempty"`
		
		KpiTotalOff *int `json:"kpiTotalOff,omitempty"`
		
		InteractionCountOn *int `json:"interactionCountOn,omitempty"`
		
		InteractionCountOff *int `json:"interactionCountOff,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		*Alias
	}{ 
		KpiTotalOn: o.KpiTotalOn,
		
		KpiTotalOff: o.KpiTotalOff,
		
		InteractionCountOn: o.InteractionCountOn,
		
		InteractionCountOff: o.InteractionCountOff,
		
		MediaType: o.MediaType,
		Alias:    (*Alias)(o),
	})
}

func (o *Kpiresult) UnmarshalJSON(b []byte) error {
	var KpiresultMap map[string]interface{}
	err := json.Unmarshal(b, &KpiresultMap)
	if err != nil {
		return err
	}
	
	if KpiTotalOn, ok := KpiresultMap["kpiTotalOn"].(float64); ok {
		KpiTotalOnInt := int(KpiTotalOn)
		o.KpiTotalOn = &KpiTotalOnInt
	}
	
	if KpiTotalOff, ok := KpiresultMap["kpiTotalOff"].(float64); ok {
		KpiTotalOffInt := int(KpiTotalOff)
		o.KpiTotalOff = &KpiTotalOffInt
	}
	
	if InteractionCountOn, ok := KpiresultMap["interactionCountOn"].(float64); ok {
		InteractionCountOnInt := int(InteractionCountOn)
		o.InteractionCountOn = &InteractionCountOnInt
	}
	
	if InteractionCountOff, ok := KpiresultMap["interactionCountOff"].(float64); ok {
		InteractionCountOffInt := int(InteractionCountOff)
		o.InteractionCountOff = &InteractionCountOffInt
	}
	
	if MediaType, ok := KpiresultMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Kpiresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
