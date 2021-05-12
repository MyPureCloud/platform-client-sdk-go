package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Contactcallbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
