package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerpreview
type Dialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId - The contact associated with this preview data pop
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId - The contactList associated with this preview data pop.
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId - The campaignId associated with this preview data pop.
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns - The phone number columns associated with this campaign
	PhoneNumberColumns *[]Phonenumbercolumn `json:"phoneNumberColumns,omitempty"`

}

func (o *Dialerpreview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerpreview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		CampaignId *string `json:"campaignId,omitempty"`
		
		PhoneNumberColumns *[]Phonenumbercolumn `json:"phoneNumberColumns,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ContactId: o.ContactId,
		
		ContactListId: o.ContactListId,
		
		CampaignId: o.CampaignId,
		
		PhoneNumberColumns: o.PhoneNumberColumns,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerpreview) UnmarshalJSON(b []byte) error {
	var DialerpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &DialerpreviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialerpreviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ContactId, ok := DialerpreviewMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
	
	if ContactListId, ok := DialerpreviewMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
	
	if CampaignId, ok := DialerpreviewMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
	
	if PhoneNumberColumns, ok := DialerpreviewMap["phoneNumberColumns"].([]interface{}); ok {
		PhoneNumberColumnsString, _ := json.Marshal(PhoneNumberColumns)
		json.Unmarshal(PhoneNumberColumnsString, &o.PhoneNumberColumns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
