package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengerpositionsettings - Position settings for messenger
type Messengerpositionsettings struct { 
	// Alignment - The alignment for messenger position
	Alignment *string `json:"alignment,omitempty"`


	// SideSpace - The sidespace value for messenger position
	SideSpace *int `json:"sideSpace,omitempty"`


	// BottomSpace - The bottomspace value for messenger position
	BottomSpace *int `json:"bottomSpace,omitempty"`

}

func (o *Messengerpositionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengerpositionsettings
	
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

func (o *Messengerpositionsettings) UnmarshalJSON(b []byte) error {
	var MessengerpositionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &MessengerpositionsettingsMap)
	if err != nil {
		return err
	}
	
	if Alignment, ok := MessengerpositionsettingsMap["alignment"].(string); ok {
		o.Alignment = &Alignment
	}
	
	if SideSpace, ok := MessengerpositionsettingsMap["sideSpace"].(float64); ok {
		SideSpaceInt := int(SideSpace)
		o.SideSpace = &SideSpaceInt
	}
	
	if BottomSpace, ok := MessengerpositionsettingsMap["bottomSpace"].(float64); ok {
		BottomSpaceInt := int(BottomSpace)
		o.BottomSpace = &BottomSpaceInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messengerpositionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
