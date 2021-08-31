package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicphone
type Phonechangetopicphone struct { 
	// UserAgentInfo
	UserAgentInfo *Phonechangetopicuseragentinfo `json:"userAgentInfo,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Status
	Status *Phonechangetopicphonestatus `json:"status,omitempty"`


	// SecondaryStatus
	SecondaryStatus *Phonechangetopicphonestatus `json:"secondaryStatus,omitempty"`

}

func (o *Phonechangetopicphone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicphone
	
	return json.Marshal(&struct { 
		UserAgentInfo *Phonechangetopicuseragentinfo `json:"userAgentInfo,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Status *Phonechangetopicphonestatus `json:"status,omitempty"`
		
		SecondaryStatus *Phonechangetopicphonestatus `json:"secondaryStatus,omitempty"`
		*Alias
	}{ 
		UserAgentInfo: o.UserAgentInfo,
		
		Id: o.Id,
		
		Status: o.Status,
		
		SecondaryStatus: o.SecondaryStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonechangetopicphone) UnmarshalJSON(b []byte) error {
	var PhonechangetopicphoneMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopicphoneMap)
	if err != nil {
		return err
	}
	
	if UserAgentInfo, ok := PhonechangetopicphoneMap["userAgentInfo"].(map[string]interface{}); ok {
		UserAgentInfoString, _ := json.Marshal(UserAgentInfo)
		json.Unmarshal(UserAgentInfoString, &o.UserAgentInfo)
	}
	
	if Id, ok := PhonechangetopicphoneMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Status, ok := PhonechangetopicphoneMap["status"].(map[string]interface{}); ok {
		StatusString, _ := json.Marshal(Status)
		json.Unmarshal(StatusString, &o.Status)
	}
	
	if SecondaryStatus, ok := PhonechangetopicphoneMap["secondaryStatus"].(map[string]interface{}); ok {
		SecondaryStatusString, _ := json.Marshal(SecondaryStatus)
		json.Unmarshal(SecondaryStatusString, &o.SecondaryStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopicphone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
