package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Postoutputcontract - The schemas defining all of the expected responses/outputs.
type Postoutputcontract struct { 
	// SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller.
	SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`

}

func (o *Postoutputcontract) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Postoutputcontract
	
	return json.Marshal(&struct { 
		SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`
		*Alias
	}{ 
		SuccessSchema: o.SuccessSchema,
		Alias:    (*Alias)(o),
	})
}

func (o *Postoutputcontract) UnmarshalJSON(b []byte) error {
	var PostoutputcontractMap map[string]interface{}
	err := json.Unmarshal(b, &PostoutputcontractMap)
	if err != nil {
		return err
	}
	
	if SuccessSchema, ok := PostoutputcontractMap["successSchema"].(map[string]interface{}); ok {
		SuccessSchemaString, _ := json.Marshal(SuccessSchema)
		json.Unmarshal(SuccessSchemaString, &o.SuccessSchema)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Postoutputcontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
