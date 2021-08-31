package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicjourneycustomer
type Conversationvideoeventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

func (o *Conversationvideoeventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicjourneycustomer
	
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

func (o *Conversationvideoeventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var ConversationvideoeventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationvideoeventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationvideoeventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationvideoeventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
