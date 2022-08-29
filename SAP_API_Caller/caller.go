package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-chat-activity-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetChatActivityCollection(objectID, iD, text string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ChatActivityCollection":
			func() {
				c.ChatActivityCollection(objectID, iD)
				wg.Done()
			}()
		case "ChatActivityTextCollection":
			func() {
				c.ChatActivityTextCollection(objectID, text)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ChatActivityCollection(objectID, iD string) {
	chatActivityCollectionData, err := c.callChatActivitySrvAPIRequirementChatActivityCollection("ChatActivityCollection", objectID, iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(chatActivityCollectionData)

	chatActivityPartiesData, err := c.callChatActivityParties(chatActivityCollectionData[0].ToChatActivityParties)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(chatActivityPartiesData)

}

func (c *SAPAPICaller) callChatActivitySrvAPIRequirementChatActivityCollection(api, objectID, iD string) ([]sap_api_output_formatter.ChatActivityCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithChatActivityCollection(req, objectID, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToChatActivityCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callChatActivityParties(url string) ([]sap_api_output_formatter.ChatActivityParties, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToChatActivityParties(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ChatActivityTextCollection(objectID, text string) {
	data, err := c.callChatActivitySrvAPIRequirementChatActivityTextCollection("ChatActivityTextCollectionCollection", objectID, text)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callChatActivitySrvAPIRequirementChatActivityTextCollection(api, objectID, text string) ([]sap_api_output_formatter.ChatActivityTextCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithChatActivityCollection(req, objectID, text)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToChatActivityTextCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithChatActivityCollection(req *http.Request, objectID, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and ID eq '%s'", objectID, iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithChatActivityTextCollection(req *http.Request, objectID, text string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and substringof('%s', Text)", objectID, text))
	req.URL.RawQuery = params.Encode()
}
