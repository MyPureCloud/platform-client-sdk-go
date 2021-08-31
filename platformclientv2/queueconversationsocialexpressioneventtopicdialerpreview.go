package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicdialerpreview
type Queueconversationsocialexpressioneventtopicdialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns
	PhoneNumberColumns *[]Queueconversationsocialexpressioneventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicdialerpreview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicdialerpreview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		CampaignId *string `json:"campaignId,omitempty"`
		
		PhoneNumberColumns *[]Queueconversationsocialexpressioneventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`
		
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

func (o *Queueconversationsocialexpressioneventtopicdialerpreview) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicdialerpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicdialerpreviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ContactId, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
	
	if ContactListId, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
	
	if CampaignId, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
	
	if PhoneNumberColumns, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["phoneNumberColumns"].([]interface{}); ok {
		PhoneNumberColumnsString, _ := json.Marshal(PhoneNumberColumns)
		json.Unmarshal(PhoneNumberColumnsString, &o.PhoneNumberColumns)
	}
	
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicdialerpreviewMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicdialerpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
