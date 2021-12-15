package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicjourneycustomer - A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
type Conversationeventtopicjourneycustomer struct { 
	// Id - An ID of a customer within the Journey System at a point-in-time.  Note that a customer entity can have multiple customerIds based on the stitching process.  Depending on the context within the PureCloud conversation, this may or may not be mutable.
	Id *string `json:"id,omitempty"`


	// IdType - The type of the customerId within the Journey System (e.g. cookie).
	IdType *string `json:"idType,omitempty"`

}

func (o *Conversationeventtopicjourneycustomer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicjourneycustomer
	
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

func (o *Conversationeventtopicjourneycustomer) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicjourneycustomerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicjourneycustomerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicjourneycustomerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationeventtopicjourneycustomerMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
