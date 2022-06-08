package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopicfacebookscopedid
type Externalcontactsunresolvedcontactchangedtopicfacebookscopedid struct { 
	// ScopedId
	ScopedId *string `json:"scopedId,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopicfacebookscopedid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopicfacebookscopedid
	
	return json.Marshal(&struct { 
		ScopedId *string `json:"scopedId,omitempty"`
		*Alias
	}{ 
		ScopedId: o.ScopedId,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopicfacebookscopedid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopicfacebookscopedidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopicfacebookscopedidMap)
	if err != nil {
		return err
	}
	
	if ScopedId, ok := ExternalcontactsunresolvedcontactchangedtopicfacebookscopedidMap["scopedId"].(string); ok {
		o.ScopedId = &ScopedId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopicfacebookscopedid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
