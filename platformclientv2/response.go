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

func (o *Response) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Response
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		Libraries: o.Libraries,
		
		Texts: o.Texts,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		InteractionType: o.InteractionType,
		
		Substitutions: o.Substitutions,
		
		SubstitutionsSchema: o.SubstitutionsSchema,
		
		ResponseType: o.ResponseType,
		
		MessagingTemplate: o.MessagingTemplate,
		
		Assets: o.Assets,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Response) UnmarshalJSON(b []byte) error {
	var ResponseMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ResponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ResponseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Version, ok := ResponseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Libraries, ok := ResponseMap["libraries"].([]interface{}); ok {
		LibrariesString, _ := json.Marshal(Libraries)
		json.Unmarshal(LibrariesString, &o.Libraries)
	}
	
	if Texts, ok := ResponseMap["texts"].([]interface{}); ok {
		TextsString, _ := json.Marshal(Texts)
		json.Unmarshal(TextsString, &o.Texts)
	}
	
	if CreatedBy, ok := ResponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := ResponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if InteractionType, ok := ResponseMap["interactionType"].(string); ok {
		o.InteractionType = &InteractionType
	}
	
	if Substitutions, ok := ResponseMap["substitutions"].([]interface{}); ok {
		SubstitutionsString, _ := json.Marshal(Substitutions)
		json.Unmarshal(SubstitutionsString, &o.Substitutions)
	}
	
	if SubstitutionsSchema, ok := ResponseMap["substitutionsSchema"].(map[string]interface{}); ok {
		SubstitutionsSchemaString, _ := json.Marshal(SubstitutionsSchema)
		json.Unmarshal(SubstitutionsSchemaString, &o.SubstitutionsSchema)
	}
	
	if ResponseType, ok := ResponseMap["responseType"].(string); ok {
		o.ResponseType = &ResponseType
	}
	
	if MessagingTemplate, ok := ResponseMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if Assets, ok := ResponseMap["assets"].([]interface{}); ok {
		AssetsString, _ := json.Marshal(Assets)
		json.Unmarshal(AssetsString, &o.Assets)
	}
	
	if SelfUri, ok := ResponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Response) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
