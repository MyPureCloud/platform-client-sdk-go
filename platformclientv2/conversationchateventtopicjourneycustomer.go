package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchateventtopicjourneycustomer - A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
type Conversationchateventtopicjourneycustomer struct { 
	// Id - An ID of a customer within the Journey System at a point-in-time.  Note that a customer entity can have multiple customerIds based on the stitching process.  Depending on the context within the PureCloud conversation, this may or may not be mutable.
	Id *string `json:"id,omitempty"`


	// IdType - The type of the customerId within the Journey System (e.g. cookie).
	IdType *string `json:"idType,omitempty"`

}

func (o *Conversationchateventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchateventtopicjourneycustomer
	
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

func (o *Conversationchateventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var ConversationchateventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchateventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationchateventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationchateventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
