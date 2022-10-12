package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Initiatingaction
type Initiatingaction struct { 
	// TransactionId - Id of the audit initiating the transaction
	TransactionId *string `json:"transactionId,omitempty"`


	// ActionContext - Action of the audit initiating the transaction
	ActionContext *string `json:"actionContext,omitempty"`

}

func (o *Initiatingaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Initiatingaction
	
	return json.Marshal(&struct { 
		TransactionId *string `json:"transactionId,omitempty"`
		
		ActionContext *string `json:"actionContext,omitempty"`
		*Alias
	}{ 
		TransactionId: o.TransactionId,
		
		ActionContext: o.ActionContext,
		Alias:    (*Alias)(o),
	})
}

func (o *Initiatingaction) UnmarshalJSON(b []byte) error {
	var InitiatingactionMap map[string]interface{}
	err := json.Unmarshal(b, &InitiatingactionMap)
	if err != nil {
		return err
	}
	
	if TransactionId, ok := InitiatingactionMap["transactionId"].(string); ok {
		o.TransactionId = &TransactionId
	}
    
	if ActionContext, ok := InitiatingactionMap["actionContext"].(string); ok {
		o.ActionContext = &ActionContext
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Initiatingaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
