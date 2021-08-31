package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicjourneycustomer
type Queueconversationcallbackeventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

func (o *Queueconversationcallbackeventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicjourneycustomer
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		IdType: o.IdType,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcallbackeventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := QueueconversationcallbackeventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
