package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Page
type Page struct { 
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


	// RootContainer
	RootContainer *map[string]interface{} `json:"rootContainer,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Page) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Page
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		RootContainer *map[string]interface{} `json:"rootContainer,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VersionId: o.VersionId,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		RootContainer: o.RootContainer,
		
		Properties: o.Properties,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Page) UnmarshalJSON(b []byte) error {
	var PageMap map[string]interface{}
	err := json.Unmarshal(b, &PageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PageMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VersionId, ok := PageMap["versionId"].(string); ok {
		o.VersionId = &VersionId
	}
	
	if createdDateString, ok := PageMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := PageMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if RootContainer, ok := PageMap["rootContainer"].(map[string]interface{}); ok {
		RootContainerString, _ := json.Marshal(RootContainer)
		json.Unmarshal(RootContainerString, &o.RootContainer)
	}
	
	if Properties, ok := PageMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if SelfUri, ok := PageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Page) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
