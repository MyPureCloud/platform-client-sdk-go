package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicjourneycustomer
type Queueconversationsocialexpressioneventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicjourneycustomer
	
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

func (o *Queueconversationsocialexpressioneventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := QueueconversationsocialexpressioneventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
