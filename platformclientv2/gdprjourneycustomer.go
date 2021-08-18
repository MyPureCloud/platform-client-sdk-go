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

func (u *Gdprjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gdprjourneycustomer

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Gdprjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
