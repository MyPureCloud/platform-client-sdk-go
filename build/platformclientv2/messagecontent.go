package platformclientv2
import (
	"encoding/json"
)

// Messagecontent - Message content element
type Messagecontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Location - Location object
	Location *Contentlocation `json:"location,omitempty"`


	// Attachment - Attachment object
	Attachment *Contentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply object
	QuickReply *Contentquickreply `json:"quickReply,omitempty"`


	// Generic - Generic content object
	Generic *Contentgeneric `json:"generic,omitempty"`


	// List - List content object
	List *Contentlist `json:"list,omitempty"`


	// Template - Template notification object
	Template *Contentnotificationtemplate `json:"template,omitempty"`


	// Reactions - A list of reactions
	Reactions *[]Contentreaction `json:"reactions,omitempty"`


	// Mention - This is used to identify who the message is sent to, as well as who it was sent from. This information is channel specific - depends on capabilities to describe party by the platform
	Mention *Messagingrecipient `json:"mention,omitempty"`


	// Postback - The postback object result of a user clicking in a button
	Postback *Contentpostback `json:"postback,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
