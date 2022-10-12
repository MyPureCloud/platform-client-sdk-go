package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outbounddomain
type Outbounddomain struct { 
	// Id - Unique Id of the domain such as: example.com
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CnameVerificationResult - CNAME registration Status
	CnameVerificationResult *Verificationresult `json:"cnameVerificationResult,omitempty"`


	// DkimVerificationResult - DKIM registration Status
	DkimVerificationResult *Verificationresult `json:"dkimVerificationResult,omitempty"`


	// SenderType - Sender Type
	SenderType *string `json:"senderType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Outbounddomain) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outbounddomain
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CnameVerificationResult *Verificationresult `json:"cnameVerificationResult,omitempty"`
		
		DkimVerificationResult *Verificationresult `json:"dkimVerificationResult,omitempty"`
		
		SenderType *string `json:"senderType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CnameVerificationResult: o.CnameVerificationResult,
		
		DkimVerificationResult: o.DkimVerificationResult,
		
		SenderType: o.SenderType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Outbounddomain) UnmarshalJSON(b []byte) error {
	var OutbounddomainMap map[string]interface{}
	err := json.Unmarshal(b, &OutbounddomainMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutbounddomainMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OutbounddomainMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CnameVerificationResult, ok := OutbounddomainMap["cnameVerificationResult"].(map[string]interface{}); ok {
		CnameVerificationResultString, _ := json.Marshal(CnameVerificationResult)
		json.Unmarshal(CnameVerificationResultString, &o.CnameVerificationResult)
	}
	
	if DkimVerificationResult, ok := OutbounddomainMap["dkimVerificationResult"].(map[string]interface{}); ok {
		DkimVerificationResultString, _ := json.Marshal(DkimVerificationResult)
		json.Unmarshal(DkimVerificationResultString, &o.DkimVerificationResult)
	}
	
	if SenderType, ok := OutbounddomainMap["senderType"].(string); ok {
		o.SenderType = &SenderType
	}
    
	if SelfUri, ok := OutbounddomainMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outbounddomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
