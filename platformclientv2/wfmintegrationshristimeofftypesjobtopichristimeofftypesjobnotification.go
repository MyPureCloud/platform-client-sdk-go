package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification
type Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Entities
	Entities *[]Wfmintegrationshristimeofftypesjobtopichristimeofftype `json:"entities,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// VarError
	VarError *Wfmintegrationshristimeofftypesjobtopicerrorbody `json:"error,omitempty"`

}

func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Entities *[]Wfmintegrationshristimeofftypesjobtopichristimeofftype `json:"entities,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *Wfmintegrationshristimeofftypesjobtopicerrorbody `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Entities: o.Entities,
		
		Status: o.Status,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification) UnmarshalJSON(b []byte) error {
	var WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Entities, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Status, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarError, ok := WfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotificationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintegrationshristimeofftypesjobtopichristimeofftypesjobnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
