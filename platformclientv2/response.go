package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Response - Contains information about a response.
type Response struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// Libraries - One or more libraries response is associated with.
	Libraries *[]Domainentityref `json:"libraries,omitempty"`


	// Texts - One or more texts associated with the response.
	Texts *[]Responsetext `json:"texts,omitempty"`


	// CreatedBy - User that created the response
	CreatedBy *User `json:"createdBy,omitempty"`


	// DateCreated - The date and time the response was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// InteractionType - The interaction type for this response.
	InteractionType *string `json:"interactionType,omitempty"`


	// Substitutions - Details about any text substitutions used in the texts for this response.
	Substitutions *[]Responsesubstitution `json:"substitutions,omitempty"`


	// SubstitutionsSchema - Metadata about the text substitutions in json schema format.
	SubstitutionsSchema *Jsonschemadocument `json:"substitutionsSchema,omitempty"`


	// ResponseType - The response type represented by the response.
	ResponseType *string `json:"responseType,omitempty"`


	// MessagingTemplate - An optional messaging template definition for responseType.MessagingTemplate.
	MessagingTemplate *Messagingtemplate `json:"messagingTemplate,omitempty"`


	// Assets - Assets used in the response
	Assets *[]Addressableentityref `json:"assets,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Response) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Response

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Libraries *[]Domainentityref `json:"libraries,omitempty"`
		
		Texts *[]Responsetext `json:"texts,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		InteractionType *string `json:"interactionType,omitempty"`
		
		Substitutions *[]Responsesubstitution `json:"substitutions,omitempty"`
		
		SubstitutionsSchema *Jsonschemadocument `json:"substitutionsSchema,omitempty"`
		
		ResponseType *string `json:"responseType,omitempty"`
		
		MessagingTemplate *Messagingtemplate `json:"messagingTemplate,omitempty"`
		
		Assets *[]Addressableentityref `json:"assets,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Version: u.Version,
		
		Libraries: u.Libraries,
		
		Texts: u.Texts,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		InteractionType: u.InteractionType,
		
		Substitutions: u.Substitutions,
		
		SubstitutionsSchema: u.SubstitutionsSchema,
		
		ResponseType: u.ResponseType,
		
		MessagingTemplate: u.MessagingTemplate,
		
		Assets: u.Assets,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Response) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
