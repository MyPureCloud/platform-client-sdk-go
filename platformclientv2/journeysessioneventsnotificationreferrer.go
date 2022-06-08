package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationreferrer
type Journeysessioneventsnotificationreferrer struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`


	// Keywords
	Keywords *string `json:"keywords,omitempty"`


	// Pathname
	Pathname *string `json:"pathname,omitempty"`


	// QueryString
	QueryString *string `json:"queryString,omitempty"`


	// Fragment
	Fragment *string `json:"fragment,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Medium
	Medium *string `json:"medium,omitempty"`

}

func (o *Journeysessioneventsnotificationreferrer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationreferrer
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		
		Keywords *string `json:"keywords,omitempty"`
		
		Pathname *string `json:"pathname,omitempty"`
		
		QueryString *string `json:"queryString,omitempty"`
		
		Fragment *string `json:"fragment,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Medium *string `json:"medium,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Domain: o.Domain,
		
		Hostname: o.Hostname,
		
		Keywords: o.Keywords,
		
		Pathname: o.Pathname,
		
		QueryString: o.QueryString,
		
		Fragment: o.Fragment,
		
		Name: o.Name,
		
		Medium: o.Medium,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationreferrer) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationreferrerMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationreferrerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := JourneysessioneventsnotificationreferrerMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Domain, ok := JourneysessioneventsnotificationreferrerMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Hostname, ok := JourneysessioneventsnotificationreferrerMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := JourneysessioneventsnotificationreferrerMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Pathname, ok := JourneysessioneventsnotificationreferrerMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := JourneysessioneventsnotificationreferrerMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Fragment, ok := JourneysessioneventsnotificationreferrerMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Name, ok := JourneysessioneventsnotificationreferrerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Medium, ok := JourneysessioneventsnotificationreferrerMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationreferrer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
