package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gdprjourneycustomer
type Gdprjourneycustomer struct { 
	// VarType - The type of the customerId within the Journey System (e.g. cookie). Required if `id` is defined.
	VarType *string `json:"type,omitempty"`


	// Id - An ID of a customer within the Journey System at a point-in-time. Required if `type` is defined.
	Id *string `json:"id,omitempty"`

}

func (o *Gdprjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gdprjourneycustomer
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Gdprjourneycustomer) UnmarshalJSON(b []byte) error {
	var GdprjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &GdprjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := GdprjourneycustomerMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Id, ok := GdprjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Gdprjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
