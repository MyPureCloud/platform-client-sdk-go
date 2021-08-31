package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recipient
type Recipient struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Flow - An automate flow object which defines the set of actions to be taken, when a message is received by this provisioned phone number.
	Flow *Flow `json:"flow,omitempty"`


	// DateCreated - Date this recipient was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this recipient was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User that created this recipient
	CreatedBy *User `json:"createdBy,omitempty"`


	// ModifiedBy - User that modified this recipient
	ModifiedBy *User `json:"modifiedBy,omitempty"`


	// MessengerType - The messenger type for this recipient
	MessengerType *string `json:"messengerType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Recipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recipient
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Flow *Flow `json:"flow,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		ModifiedBy *User `json:"modifiedBy,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Flow: o.Flow,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		MessengerType: o.MessengerType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Recipient) UnmarshalJSON(b []byte) error {
	var RecipientMap map[string]interface{}
	err := json.Unmarshal(b, &RecipientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecipientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := RecipientMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Flow, ok := RecipientMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if dateCreatedString, ok := RecipientMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := RecipientMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := RecipientMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := RecipientMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if MessengerType, ok := RecipientMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
	
	if SelfUri, ok := RecipientMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
