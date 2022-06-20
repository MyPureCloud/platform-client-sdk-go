package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainnetworkcommandresponse
type Domainnetworkcommandresponse struct { 
	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// CommandName
	CommandName *string `json:"commandName,omitempty"`


	// Acknowledged
	Acknowledged *bool `json:"acknowledged,omitempty"`


	// ErrorInfo
	ErrorInfo *Errordetails `json:"errorInfo,omitempty"`

}

func (o *Domainnetworkcommandresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainnetworkcommandresponse
	
	return json.Marshal(&struct { 
		CorrelationId *string `json:"correlationId,omitempty"`
		
		CommandName *string `json:"commandName,omitempty"`
		
		Acknowledged *bool `json:"acknowledged,omitempty"`
		
		ErrorInfo *Errordetails `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		CorrelationId: o.CorrelationId,
		
		CommandName: o.CommandName,
		
		Acknowledged: o.Acknowledged,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainnetworkcommandresponse) UnmarshalJSON(b []byte) error {
	var DomainnetworkcommandresponseMap map[string]interface{}
	err := json.Unmarshal(b, &DomainnetworkcommandresponseMap)
	if err != nil {
		return err
	}
	
	if CorrelationId, ok := DomainnetworkcommandresponseMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if CommandName, ok := DomainnetworkcommandresponseMap["commandName"].(string); ok {
		o.CommandName = &CommandName
	}
    
	if Acknowledged, ok := DomainnetworkcommandresponseMap["acknowledged"].(bool); ok {
		o.Acknowledged = &Acknowledged
	}
    
	if ErrorInfo, ok := DomainnetworkcommandresponseMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainnetworkcommandresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
