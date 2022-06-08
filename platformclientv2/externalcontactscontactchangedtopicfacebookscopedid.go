package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopicfacebookscopedid
type Externalcontactscontactchangedtopicfacebookscopedid struct { 
	// ScopedId
	ScopedId *string `json:"scopedId,omitempty"`

}

func (o *Externalcontactscontactchangedtopicfacebookscopedid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopicfacebookscopedid
	
	return json.Marshal(&struct { 
		ScopedId *string `json:"scopedId,omitempty"`
		*Alias
	}{ 
		ScopedId: o.ScopedId,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopicfacebookscopedid) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopicfacebookscopedidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopicfacebookscopedidMap)
	if err != nil {
		return err
	}
	
	if ScopedId, ok := ExternalcontactscontactchangedtopicfacebookscopedidMap["scopedId"].(string); ok {
		o.ScopedId = &ScopedId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopicfacebookscopedid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
