package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Referrer
type Referrer struct { 
	// Url - Referrer URL.
	Url *string `json:"url,omitempty"`


	// Domain - Referrer URL domain.
	Domain *string `json:"domain,omitempty"`


	// Hostname - Referrer URL hostname.
	Hostname *string `json:"hostname,omitempty"`


	// Keywords - Referrer keywords.
	Keywords *string `json:"keywords,omitempty"`


	// Pathname - Referrer URL pathname.
	Pathname *string `json:"pathname,omitempty"`


	// QueryString - Referrer URL querystring.
	QueryString *string `json:"queryString,omitempty"`


	// Fragment - Referrer URL fragment.
	Fragment *string `json:"fragment,omitempty"`


	// Name - Name of referrer (e.g. Yahoo!, Google, InfoSpace).
	Name *string `json:"name,omitempty"`


	// Medium - Type of referrer (e.g. search, social).
	Medium *string `json:"medium,omitempty"`

}

func (o *Referrer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Referrer
	
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

func (o *Referrer) UnmarshalJSON(b []byte) error {
	var ReferrerMap map[string]interface{}
	err := json.Unmarshal(b, &ReferrerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ReferrerMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Domain, ok := ReferrerMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Hostname, ok := ReferrerMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := ReferrerMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Pathname, ok := ReferrerMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := ReferrerMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Fragment, ok := ReferrerMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Name, ok := ReferrerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Medium, ok := ReferrerMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Referrer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
