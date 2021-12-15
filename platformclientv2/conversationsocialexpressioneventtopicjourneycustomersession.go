package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicjourneycustomersession - A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
type Conversationsocialexpressioneventtopicjourneycustomersession struct { 
	// Id - An ID of a Customer/User's session within the Journey System at a point-in-time
	Id *string `json:"id,omitempty"`


	// VarType - The type of the Customer/User's session within the Journey System (e.g. web, app)
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicjourneycustomersession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicjourneycustomersession
	
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

func (o *Conversationsocialexpressioneventtopicjourneycustomersession) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicjourneycustomersessionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicjourneycustomersessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationsocialexpressioneventtopicjourneycustomersessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if VarType, ok := ConversationsocialexpressioneventtopicjourneycustomersessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
