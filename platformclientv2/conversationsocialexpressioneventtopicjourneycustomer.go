package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicjourneycustomer
type Conversationsocialexpressioneventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicjourneycustomer
	
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

func (o *Conversationsocialexpressioneventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationsocialexpressioneventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationsocialexpressioneventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
