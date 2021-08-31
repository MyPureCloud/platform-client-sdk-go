package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicdialerpreview
type Conversationcallbackeventtopicdialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns
	PhoneNumberColumns *[]Conversationcallbackeventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationcallbackeventtopicdialerpreview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicdialerpreview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		CampaignId *string `json:"campaignId,omitempty"`
		
		PhoneNumberColumns *[]Conversationcallbackeventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
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

func (o *Conversationcallbackeventtopicdialerpreview) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicdialerpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicdialerpreviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcallbackeventtopicdialerpreviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ContactId, ok := ConversationcallbackeventtopicdialerpreviewMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
	
	if ContactListId, ok := ConversationcallbackeventtopicdialerpreviewMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
	
	if CampaignId, ok := ConversationcallbackeventtopicdialerpreviewMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
	
	if PhoneNumberColumns, ok := ConversationcallbackeventtopicdialerpreviewMap["phoneNumberColumns"].([]interface{}); ok {
		PhoneNumberColumnsString, _ := json.Marshal(PhoneNumberColumns)
		json.Unmarshal(PhoneNumberColumnsString, &o.PhoneNumberColumns)
	}
	
	if AdditionalProperties, ok := ConversationcallbackeventtopicdialerpreviewMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicdialerpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
