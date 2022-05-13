package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcallbackrequest
type Contactcallbackrequest struct { 
	// CampaignId - Campaign identifier
	CampaignId *string `json:"campaignId,omitempty"`


	// ContactListId - Contact list identifier
	ContactListId *string `json:"contactListId,omitempty"`


	// ContactId - Contact identifier
	ContactId *string `json:"contactId,omitempty"`


	// PhoneColumn - Name of the phone column containing the number to be called
	PhoneColumn *string `json:"phoneColumn,omitempty"`


	// Schedule - The scheduled time for the callback as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ\", example = \"2016-01-02T16:59:59\"
	Schedule *string `json:"schedule,omitempty"`

}

func (o *Contactcallbackrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcallbackrequest
	
	return json.Marshal(&struct { 
		CampaignId *string `json:"campaignId,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		ContactId *string `json:"contactId,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		
		Schedule *string `json:"schedule,omitempty"`
		*Alias
	}{ 
		CampaignId: o.CampaignId,
		
		ContactListId: o.ContactListId,
		
		ContactId: o.ContactId,
		
		PhoneColumn: o.PhoneColumn,
		
		Schedule: o.Schedule,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactcallbackrequest) UnmarshalJSON(b []byte) error {
	var ContactcallbackrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcallbackrequestMap)
	if err != nil {
		return err
	}
	
	if CampaignId, ok := ContactcallbackrequestMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
    
	if ContactListId, ok := ContactcallbackrequestMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if ContactId, ok := ContactcallbackrequestMap["contactId"].(string); ok {
		o.ContactId = &ContactId
	}
    
	if PhoneColumn, ok := ContactcallbackrequestMap["phoneColumn"].(string); ok {
		o.PhoneColumn = &PhoneColumn
	}
    
	if Schedule, ok := ContactcallbackrequestMap["schedule"].(string); ok {
		o.Schedule = &Schedule
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcallbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
