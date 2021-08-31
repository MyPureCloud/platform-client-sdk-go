package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopiclinestatus
type Phonechangetopiclinestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Reachable
	Reachable *bool `json:"reachable,omitempty"`


	// AddressOfRecord
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`


	// ContactAddresses
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`


	// ReachableStateTime
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`

}

func (o *Phonechangetopiclinestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopiclinestatus
	
	ReachableStateTime := new(string)
	if o.ReachableStateTime != nil {
		
		*ReachableStateTime = timeutil.Strftime(o.ReachableStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReachableStateTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Reachable *bool `json:"reachable,omitempty"`
		
		AddressOfRecord *string `json:"addressOfRecord,omitempty"`
		
		ContactAddresses *[]string `json:"contactAddresses,omitempty"`
		
		ReachableStateTime *string `json:"reachableStateTime,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Reachable: o.Reachable,
		
		AddressOfRecord: o.AddressOfRecord,
		
		ContactAddresses: o.ContactAddresses,
		
		ReachableStateTime: ReachableStateTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonechangetopiclinestatus) UnmarshalJSON(b []byte) error {
	var PhonechangetopiclinestatusMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopiclinestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonechangetopiclinestatusMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Reachable, ok := PhonechangetopiclinestatusMap["reachable"].(bool); ok {
		o.Reachable = &Reachable
	}
	
	if AddressOfRecord, ok := PhonechangetopiclinestatusMap["addressOfRecord"].(string); ok {
		o.AddressOfRecord = &AddressOfRecord
	}
	
	if ContactAddresses, ok := PhonechangetopiclinestatusMap["contactAddresses"].([]interface{}); ok {
		ContactAddressesString, _ := json.Marshal(ContactAddresses)
		json.Unmarshal(ContactAddressesString, &o.ContactAddresses)
	}
	
	if reachableStateTimeString, ok := PhonechangetopiclinestatusMap["reachableStateTime"].(string); ok {
		ReachableStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", reachableStateTimeString)
		o.ReachableStateTime = &ReachableStateTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopiclinestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
