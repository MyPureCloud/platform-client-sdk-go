package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageactivitycategoryresponse
type Historicalshrinkageactivitycategoryresponse struct { 
	// ActivityCategory - Activity category for which shrinkage data is provided
	ActivityCategory *string `json:"activityCategory,omitempty"`


	// ShrinkageForActivityCategory - Aggregated shrinkage data for the activity category
	ShrinkageForActivityCategory *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory,omitempty"`


	// ShrinkageForActivityCodes - Shrinkage for the activity codes under this activity category
	ShrinkageForActivityCodes *[]Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes,omitempty"`

}

func (o *Historicalshrinkageactivitycategoryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalshrinkageactivitycategoryresponse
	
	return json.Marshal(&struct { 
		ActivityCategory *string `json:"activityCategory,omitempty"`
		
		ShrinkageForActivityCategory *Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory,omitempty"`
		
		ShrinkageForActivityCodes *[]Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes,omitempty"`
		*Alias
	}{ 
		ActivityCategory: o.ActivityCategory,
		
		ShrinkageForActivityCategory: o.ShrinkageForActivityCategory,
		
		ShrinkageForActivityCodes: o.ShrinkageForActivityCodes,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalshrinkageactivitycategoryresponse) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageactivitycategoryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageactivitycategoryresponseMap)
	if err != nil {
		return err
	}
	
	if ActivityCategory, ok := HistoricalshrinkageactivitycategoryresponseMap["activityCategory"].(string); ok {
		o.ActivityCategory = &ActivityCategory
	}
    
	if ShrinkageForActivityCategory, ok := HistoricalshrinkageactivitycategoryresponseMap["shrinkageForActivityCategory"].(map[string]interface{}); ok {
		ShrinkageForActivityCategoryString, _ := json.Marshal(ShrinkageForActivityCategory)
		json.Unmarshal(ShrinkageForActivityCategoryString, &o.ShrinkageForActivityCategory)
	}
	
	if ShrinkageForActivityCodes, ok := HistoricalshrinkageactivitycategoryresponseMap["shrinkageForActivityCodes"].([]interface{}); ok {
		ShrinkageForActivityCodesString, _ := json.Marshal(ShrinkageForActivityCodes)
		json.Unmarshal(ShrinkageForActivityCodesString, &o.ShrinkageForActivityCodes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageactivitycategoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
