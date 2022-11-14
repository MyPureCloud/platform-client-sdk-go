package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageactivitycoderesponse
type Historicalshrinkageactivitycoderesponse struct { 
	// ActivityCodeId - The ID of the activity code for which shrinkage data is provided
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// ShrinkageForActivityCode - Aggregated shrinkage data for the activity code
	ShrinkageForActivityCode *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCode,omitempty"`

}

func (o *Historicalshrinkageactivitycoderesponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalshrinkageactivitycoderesponse
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		ShrinkageForActivityCode *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCode,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		ShrinkageForActivityCode: o.ShrinkageForActivityCode,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalshrinkageactivitycoderesponse) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageactivitycoderesponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageactivitycoderesponseMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := HistoricalshrinkageactivitycoderesponseMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if ShrinkageForActivityCode, ok := HistoricalshrinkageactivitycoderesponseMap["shrinkageForActivityCode"].(map[string]interface{}); ok {
		ShrinkageForActivityCodeString, _ := json.Marshal(ShrinkageForActivityCode)
		json.Unmarshal(ShrinkageForActivityCodeString, &o.ShrinkageForActivityCode)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageactivitycoderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
