package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopicentity
type Externalcontactsunresolvedcontactchangedtopicentity struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopicentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopicentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopicentity) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopicentityMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopicentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsunresolvedcontactchangedtopicentityMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopicentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
