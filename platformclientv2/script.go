package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Script
type Script struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VersionId
	VersionId *string `json:"versionId,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`


	// VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	VersionDate *time.Time `json:"versionDate,omitempty"`


	// StartPageId
	StartPageId *string `json:"startPageId,omitempty"`


	// StartPageName
	StartPageName *string `json:"startPageName,omitempty"`


	// Features
	Features *interface{} `json:"features,omitempty"`


	// Variables
	Variables *interface{} `json:"variables,omitempty"`


	// CustomActions
	CustomActions *interface{} `json:"customActions,omitempty"`


	// Pages
	Pages *[]Page `json:"pages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Script) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Script
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	PublishedDate := new(string)
	if o.PublishedDate != nil {
		
		*PublishedDate = timeutil.Strftime(o.PublishedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishedDate = nil
	}
	
	VersionDate := new(string)
	if o.VersionDate != nil {
		
		*VersionDate = timeutil.Strftime(o.VersionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VersionDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		VersionDate *string `json:"versionDate,omitempty"`
		
		StartPageId *string `json:"startPageId,omitempty"`
		
		StartPageName *string `json:"startPageName,omitempty"`
		
		Features *interface{} `json:"features,omitempty"`
		
		Variables *interface{} `json:"variables,omitempty"`
		
		CustomActions *interface{} `json:"customActions,omitempty"`
		
		Pages *[]Page `json:"pages,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VersionId: o.VersionId,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		PublishedDate: PublishedDate,
		
		VersionDate: VersionDate,
		
		StartPageId: o.StartPageId,
		
		StartPageName: o.StartPageName,
		
		Features: o.Features,
		
		Variables: o.Variables,
		
		CustomActions: o.CustomActions,
		
		Pages: o.Pages,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Script) UnmarshalJSON(b []byte) error {
	var ScriptMap map[string]interface{}
	err := json.Unmarshal(b, &ScriptMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScriptMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ScriptMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VersionId, ok := ScriptMap["versionId"].(string); ok {
		o.VersionId = &VersionId
	}
	
	if createdDateString, ok := ScriptMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := ScriptMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if publishedDateString, ok := ScriptMap["publishedDate"].(string); ok {
		PublishedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishedDateString)
		o.PublishedDate = &PublishedDate
	}
	
	if versionDateString, ok := ScriptMap["versionDate"].(string); ok {
		VersionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", versionDateString)
		o.VersionDate = &VersionDate
	}
	
	if StartPageId, ok := ScriptMap["startPageId"].(string); ok {
		o.StartPageId = &StartPageId
	}
	
	if StartPageName, ok := ScriptMap["startPageName"].(string); ok {
		o.StartPageName = &StartPageName
	}
	
	if Features, ok := ScriptMap["features"].(map[string]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	
	if Variables, ok := ScriptMap["variables"].(map[string]interface{}); ok {
		VariablesString, _ := json.Marshal(Variables)
		json.Unmarshal(VariablesString, &o.Variables)
	}
	
	if CustomActions, ok := ScriptMap["customActions"].(map[string]interface{}); ok {
		CustomActionsString, _ := json.Marshal(CustomActions)
		json.Unmarshal(CustomActionsString, &o.CustomActions)
	}
	
	if Pages, ok := ScriptMap["pages"].([]interface{}); ok {
		PagesString, _ := json.Marshal(Pages)
		json.Unmarshal(PagesString, &o.Pages)
	}
	
	if SelfUri, ok := ScriptMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Script) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
