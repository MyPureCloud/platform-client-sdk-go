package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailoutbounddomainresult
type Emailoutbounddomainresult struct { 
	// DnsCnameBounceRecord
	DnsCnameBounceRecord *Dnsrecordentry `json:"dnsCnameBounceRecord,omitempty"`


	// DnsTxtSendingRecord
	DnsTxtSendingRecord *Dnsrecordentry `json:"dnsTxtSendingRecord,omitempty"`


	// DomainName
	DomainName *string `json:"domainName,omitempty"`


	// SenderStatus
	SenderStatus *string `json:"senderStatus,omitempty"`


	// SenderType
	SenderType *string `json:"senderType,omitempty"`

}

func (o *Emailoutbounddomainresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailoutbounddomainresult
	
	return json.Marshal(&struct { 
		DnsCnameBounceRecord *Dnsrecordentry `json:"dnsCnameBounceRecord,omitempty"`
		
		DnsTxtSendingRecord *Dnsrecordentry `json:"dnsTxtSendingRecord,omitempty"`
		
		DomainName *string `json:"domainName,omitempty"`
		
		SenderStatus *string `json:"senderStatus,omitempty"`
		
		SenderType *string `json:"senderType,omitempty"`
		*Alias
	}{ 
		DnsCnameBounceRecord: o.DnsCnameBounceRecord,
		
		DnsTxtSendingRecord: o.DnsTxtSendingRecord,
		
		DomainName: o.DomainName,
		
		SenderStatus: o.SenderStatus,
		
		SenderType: o.SenderType,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailoutbounddomainresult) UnmarshalJSON(b []byte) error {
	var EmailoutbounddomainresultMap map[string]interface{}
	err := json.Unmarshal(b, &EmailoutbounddomainresultMap)
	if err != nil {
		return err
	}
	
	if DnsCnameBounceRecord, ok := EmailoutbounddomainresultMap["dnsCnameBounceRecord"].(map[string]interface{}); ok {
		DnsCnameBounceRecordString, _ := json.Marshal(DnsCnameBounceRecord)
		json.Unmarshal(DnsCnameBounceRecordString, &o.DnsCnameBounceRecord)
	}
	
	if DnsTxtSendingRecord, ok := EmailoutbounddomainresultMap["dnsTxtSendingRecord"].(map[string]interface{}); ok {
		DnsTxtSendingRecordString, _ := json.Marshal(DnsTxtSendingRecord)
		json.Unmarshal(DnsTxtSendingRecordString, &o.DnsTxtSendingRecord)
	}
	
	if DomainName, ok := EmailoutbounddomainresultMap["domainName"].(string); ok {
		o.DomainName = &DomainName
	}
    
	if SenderStatus, ok := EmailoutbounddomainresultMap["senderStatus"].(string); ok {
		o.SenderStatus = &SenderStatus
	}
    
	if SenderType, ok := EmailoutbounddomainresultMap["senderType"].(string); ok {
		o.SenderType = &SenderType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailoutbounddomainresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
