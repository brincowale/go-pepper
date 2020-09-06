package pepper

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Pepper struct {
	Domain       string
	ClientKey    string
	ClientSecret string
	headers      map[string]string
	httpClient   *http.Client
}

func New(domain string, clientKey string, clientSecret string, pkg string) *Pepper {
	headers := make(map[string]string)
	headers["User-Agent"] = pkg + " ANDROID [v5.26.11] [22 | SM-G930K] [@2.0x]"
	headers["Pepper-Include-Counters"] = "unread_alerts"
	headers["Pepper-Include-Prev-And-Next-Ids"] = "true"
	headers["Pepper-JSON-Format"] = "thread=list,group=ids,type=light,event=light,user=full,badge=user,formatted_text=html,message=with_code"
	headers["Pepper-Hardware-Id"] = "5bce296a65215d0bb3b9751bb77b0a1d"
	headers["Host"] = domain
	config := oauth1.NewConfig(clientKey, clientSecret)
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient = config.Client(oauth1.NoContext, &oauth1.Token{Token: "", TokenSecret: ""})
	retryClient.RetryMax = 2
	retryClient.HTTPClient.Timeout = time.Second * 30
	httpClient := retryClient.StandardClient()
	return &Pepper{domain, clientKey, clientSecret, headers, httpClient}
}

func (d *Pepper) GetHotDeals(paramsOverride map[string]string) *Deals {
	// Params
	path := fmt.Sprintf("https://%s/rest_api/v2/thread", d.Domain)
	params := make(map[string]string)
	params["order_by"] = "hot"
	params["limit"] = "50"
	q := url.Values{}
	for k, v := range paramsOverride {
		params[k] = v
	}
	for k, v := range params {
		q.Add(k, v)
	}
	req, err := http.NewRequest("GET", path, nil)
	// Http request
	req.URL.RawQuery = q.Encode()
	resp, _ := d.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// Parsing
	deals := Deals{}
	err = json.Unmarshal(body, &deals)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &deals
}

func (d *Pepper) GetNewDeals(paramsOverride map[string]string) *Deals {
	// Params
	path := "https://" + d.Domain + "/rest_api/v2/thread"
	params := make(map[string]string)
	params["order_by"] = "new"
	params["limit"] = "50"
	q := url.Values{}
	for k, v := range paramsOverride {
		params[k] = v
	}
	for k, v := range params {
		q.Add(k, v)
	}
	req, err := http.NewRequest("GET", path, nil)
	req.URL.RawQuery = q.Encode()
	// Http request

	resp, _ := d.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// Parsing
	deals := Deals{}
	err = json.Unmarshal(body, &deals)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &deals
}

func (deal *Data) PublishedAt() time.Time {
	if deal.IsHot {
		return time.Unix(deal.HotDate, 0)
	}
	return time.Unix(deal.Submitted, 0)
}
