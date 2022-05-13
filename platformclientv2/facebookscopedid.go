package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookscopedid - Scoped ID for a Facebook user interacting with a page or app
type Facebookscopedid struct { 
	// ScopedId - The unique page/app-specific scopedId for the user
	ScopedId *string `json:"scopedId,omitempty"`

}

func (o *Facebookscopedid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookscopedid
	
	return json.Marshal(&struct { 
		ScopedId *string `json:"scopedId,omitempty"`
		*Alias
	}{ 
		ScopedId: o.ScopedId,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookscopedid) UnmarshalJSON(b []byte) error {
	var FacebookscopedidMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookscopedidMap)
	if err != nil {
		return err
	}
	
	if ScopedId, ok := FacebookscopedidMap["scopedId"].(string); ok {
		o.ScopedId = &ScopedId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookscopedid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
