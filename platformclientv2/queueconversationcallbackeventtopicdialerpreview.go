package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicdialerpreview
type Queueconversationcallbackeventtopicdialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId - The contact associated with this preview data pop
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId - The contactList associated with this preview data pop.
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId - The campaignId associated with this preview data pop.
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns - The phone number columns associated with this campaign
	PhoneNumberColumns *[]Queueconversationcallbackeventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationcallbackeventtopicdialerpreview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicdialerpreview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		CampaignId *string `json:"campaignId,omitempty"`
		
		PhoneNumberColumns *[]Queueconversationcallbackeventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ContactId: o.ContactId,
		
		ContactListId: o.ContactListId,
		
		CampaignId: o.CampaignId,
		
		PhoneNumberColumns: o.PhoneNumberColumns,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicdialerpreview) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicdialerpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicdialerpreviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcallbackeventtopicdialerpreviewMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ContactId, ok := QueueconversationcallbackeventtopicdialerpreviewMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
    
	if ContactListId, ok := QueueconversationcallbackeventtopicdialerpreviewMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if CampaignId, ok := QueueconversationcallbackeventtopicdialerpreviewMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
    
	if PhoneNumberColumns, ok := QueueconversationcallbackeventtopicdialerpreviewMap["phoneNumberColumns"].([]interface{}); ok {
		PhoneNumberColumnsString, _ := json.Marshal(PhoneNumberColumns)
		json.Unmarshal(PhoneNumberColumnsString, &o.PhoneNumberColumns)
	}
	
	if AdditionalProperties, ok := QueueconversationcallbackeventtopicdialerpreviewMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicdialerpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
