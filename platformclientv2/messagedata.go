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

func (o *Messagedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagedata
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		ProviderMessageId: o.ProviderMessageId,
		
		Timestamp: Timestamp,
		
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		Direction: o.Direction,
		
		MessengerType: o.MessengerType,
		
		TextBody: o.TextBody,
		
		Status: o.Status,
		
		Media: o.Media,
		
		Stickers: o.Stickers,
		
		CreatedBy: o.CreatedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagedata) UnmarshalJSON(b []byte) error {
	var MessagedataMap map[string]interface{}
	err := json.Unmarshal(b, &MessagedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagedataMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MessagedataMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ProviderMessageId, ok := MessagedataMap["providerMessageId"].(string); ok {
		o.ProviderMessageId = &ProviderMessageId
	}
	
	if timestampString, ok := MessagedataMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if FromAddress, ok := MessagedataMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
	
	if ToAddress, ok := MessagedataMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
	
	if Direction, ok := MessagedataMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if MessengerType, ok := MessagedataMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
	
	if TextBody, ok := MessagedataMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
	
	if Status, ok := MessagedataMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Media, ok := MessagedataMap["media"].([]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Stickers, ok := MessagedataMap["stickers"].([]interface{}); ok {
		StickersString, _ := json.Marshal(Stickers)
		json.Unmarshal(StickersString, &o.Stickers)
	}
	
	if CreatedBy, ok := MessagedataMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if SelfUri, ok := MessagedataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
