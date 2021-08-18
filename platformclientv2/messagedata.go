package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagedata
type Messagedata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ProviderMessageId - The unique identifier of the message from provider
	ProviderMessageId *string `json:"providerMessageId,omitempty"`


	// Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// FromAddress - The sender of the text message.
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The recipient of the text message.
	ToAddress *string `json:"toAddress,omitempty"`


	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`


	// MessengerType - Type of text messenger.
	MessengerType *string `json:"messengerType,omitempty"`


	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// Status - The status of the message.
	Status *string `json:"status,omitempty"`


	// Media - The media details associated to a message.
	Media *[]Messagemedia `json:"media,omitempty"`


	// Stickers - The sticker details associated to a message.
	Stickers *[]Messagesticker `json:"stickers,omitempty"`


	// CreatedBy - User who sent this message.
	CreatedBy *User `json:"createdBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Messagedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagedata

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ProviderMessageId *string `json:"providerMessageId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Media *[]Messagemedia `json:"media,omitempty"`
		
		Stickers *[]Messagesticker `json:"stickers,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ProviderMessageId: u.ProviderMessageId,
		
		Timestamp: Timestamp,
		
		FromAddress: u.FromAddress,
		
		ToAddress: u.ToAddress,
		
		Direction: u.Direction,
		
		MessengerType: u.MessengerType,
		
		TextBody: u.TextBody,
		
		Status: u.Status,
		
		Media: u.Media,
		
		Stickers: u.Stickers,
		
		CreatedBy: u.CreatedBy,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
