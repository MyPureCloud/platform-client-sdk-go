package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintegrationshristimeofftypesjobtopichristimeofftype
type Wfmintegrationshristimeofftypesjobtopichristimeofftype struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SecondaryId
	SecondaryId *string `json:"secondaryId,omitempty"`

}

func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintegrationshristimeofftypesjobtopichristimeofftype
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SecondaryId *string `json:"secondaryId,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SecondaryId: o.SecondaryId,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftype) UnmarshalJSON(b []byte) error {
	var WfmintegrationshristimeofftypesjobtopichristimeofftypeMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintegrationshristimeofftypesjobtopichristimeofftypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SecondaryId, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypeMap["secondaryId"].(string); ok {
		o.SecondaryId = &SecondaryId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
