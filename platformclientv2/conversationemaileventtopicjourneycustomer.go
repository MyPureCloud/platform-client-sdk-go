package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneycustomer
type Conversationemaileventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

func (o *Conversationemaileventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicjourneycustomer
	
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

func (o *Conversationemaileventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var ConversationemaileventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationemaileventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationemaileventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationemaileventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
