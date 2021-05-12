package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Gdprjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
