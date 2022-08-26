package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Labelresponse
type Labelresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the label.
	Name *string `json:"name,omitempty"`


	// Color - The color of the label.
	Color *string `json:"color,omitempty"`


	// DateCreated - The creation date and time of the label. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The last modification date and time of the label. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DocumentCount - Number of documents assigned to this label.
	DocumentCount *int `json:"documentCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Labelresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Labelresponse
	
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
		
		Color *string `json:"color,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DocumentCount *int `json:"documentCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Color: o.Color,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DocumentCount: o.DocumentCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Labelresponse) UnmarshalJSON(b []byte) error {
	var LabelresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LabelresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LabelresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LabelresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Color, ok := LabelresponseMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if dateCreatedString, ok := LabelresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := LabelresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if DocumentCount, ok := LabelresponseMap["documentCount"].(float64); ok {
		DocumentCountInt := int(DocumentCount)
		o.DocumentCount = &DocumentCountInt
	}
	
	if SelfUri, ok := LabelresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Labelresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
