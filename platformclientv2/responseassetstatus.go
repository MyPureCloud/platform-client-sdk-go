package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseassetstatus
type Responseassetstatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Status - Upload status of the asset
	Status *string `json:"status,omitempty"`


	// ErrorCode - Error code. Used for localization
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessage - Error message that explains upload failure status 
	ErrorMessage *string `json:"errorMessage,omitempty"`

}

func (o *Responseassetstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseassetstatus
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseassetstatus) UnmarshalJSON(b []byte) error {
	var ResponseassetstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetstatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ResponseassetstatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := ResponseassetstatusMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorCode, ok := ResponseassetstatusMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := ResponseassetstatusMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responseassetstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
