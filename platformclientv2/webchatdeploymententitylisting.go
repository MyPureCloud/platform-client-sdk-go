package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatdeploymententitylisting
type Webchatdeploymententitylisting struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Webchatdeployment `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Webchatdeploymententitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatdeploymententitylisting
	
	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		Entities *[]Webchatdeployment `json:"entities,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Total: o.Total,
		
		Entities: o.Entities,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatdeploymententitylisting) UnmarshalJSON(b []byte) error {
	var WebchatdeploymententitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatdeploymententitylistingMap)
	if err != nil {
		return err
	}
	
	if Total, ok := WebchatdeploymententitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if Entities, ok := WebchatdeploymententitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if SelfUri, ok := WebchatdeploymententitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatdeploymententitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
