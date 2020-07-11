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

}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
