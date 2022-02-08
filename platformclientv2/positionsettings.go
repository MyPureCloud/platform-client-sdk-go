package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Positionsettings - Settings concerning position
type Positionsettings struct { 
	// Alignment - The alignment for position
	Alignment *string `json:"alignment,omitempty"`


	// SideSpace - The sidespace value for position
	SideSpace *int `json:"sideSpace,omitempty"`


	// BottomSpace - The bottomspace value for position
	BottomSpace *int `json:"bottomSpace,omitempty"`

}

func (o *Positionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Positionsettings
	
	return json.Marshal(&struct { 
		Alignment *string `json:"alignment,omitempty"`
		
		SideSpace *int `json:"sideSpace,omitempty"`
		
		BottomSpace *int `json:"bottomSpace,omitempty"`
		*Alias
	}{ 
		Alignment: o.Alignment,
		
		SideSpace: o.SideSpace,
		
		BottomSpace: o.BottomSpace,
		Alias:    (*Alias)(o),
	})
}

func (o *Positionsettings) UnmarshalJSON(b []byte) error {
	var PositionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &PositionsettingsMap)
	if err != nil {
		return err
	}
	
	if Alignment, ok := PositionsettingsMap["alignment"].(string); ok {
		o.Alignment = &Alignment
	}
	
	if SideSpace, ok := PositionsettingsMap["sideSpace"].(float64); ok {
		SideSpaceInt := int(SideSpace)
		o.SideSpace = &SideSpaceInt
	}
	
	if BottomSpace, ok := PositionsettingsMap["bottomSpace"].(float64); ok {
		BottomSpaceInt := int(BottomSpace)
		o.BottomSpace = &BottomSpaceInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Positionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
