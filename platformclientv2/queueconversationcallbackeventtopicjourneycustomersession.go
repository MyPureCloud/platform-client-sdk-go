package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicjourneycustomersession - A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
type Queueconversationcallbackeventtopicjourneycustomersession struct { 
	// Id - An ID of a Customer/User's session within the Journey System at a point-in-time
	Id *string `json:"id,omitempty"`


	// VarType - The type of the Customer/User's session within the Journey System (e.g. web, app)
	VarType *string `json:"type,omitempty"`

}

func (o *Queueconversationcallbackeventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicjourneycustomersession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcallbackeventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := QueueconversationcallbackeventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
