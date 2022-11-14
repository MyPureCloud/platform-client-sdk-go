package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authorizationsettings
type Authorizationsettings struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// AnalysisEnabled - Boolean showing if organization is opted in or not to unused role/perm analysis
	AnalysisEnabled *bool `json:"analysisEnabled,omitempty"`


	// AnalysisDays - Integer number of days to analyze user usage
	AnalysisDays *int `json:"analysisDays,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Authorizationsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authorizationsettings
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AnalysisEnabled *bool `json:"analysisEnabled,omitempty"`
		
		AnalysisDays *int `json:"analysisDays,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		AnalysisEnabled: o.AnalysisEnabled,
		
		AnalysisDays: o.AnalysisDays,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Authorizationsettings) UnmarshalJSON(b []byte) error {
	var AuthorizationsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AuthorizationsettingsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthorizationsettingsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if AnalysisEnabled, ok := AuthorizationsettingsMap["analysisEnabled"].(bool); ok {
		o.AnalysisEnabled = &AnalysisEnabled
	}
    
	if AnalysisDays, ok := AuthorizationsettingsMap["analysisDays"].(float64); ok {
		AnalysisDaysInt := int(AnalysisDays)
		o.AnalysisDays = &AnalysisDaysInt
	}
	
	if SelfUri, ok := AuthorizationsettingsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Authorizationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
