package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller/requests"
	"sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller/responses"

	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}
func (c *SAPAPICaller) AsyncPostPurchasingInfoRecord(
	generalPurchasingOrganizationPlant *requests.GeneralPurchasingOrganizationPlant,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "GeneralPurchasingOrganizationPlant":
			func() {
				c.GeneralPurchasingOrganizationPlant(generalPurchasingOrganizationPlant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) GeneralPurchasingOrganizationPlant(generalPurchasingOrganizationPlant *requests.GeneralPurchasingOrganizationPlant) {
	err := c.callPurchasingInfoRecordSrvAPIRequirementGeneralPurchasingOrganizationPlant("A_PurchasingInfoRecord", generalPurchasingOrganizationPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalPurchasingOrganizationPlant)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementGeneralPurchasingOrganizationPlant(api string, generalPurchasingOrganizationPlant *requests.GeneralPurchasingOrganizationPlant) error {
	body, err := json.Marshal(generalPurchasingOrganizationPlant)
	if err != nil {
		return xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return xerrors.Errorf("bad response:%s", string(byteArray))
	}

	resBody := responses.General{}
	json.Unmarshal(byteArray, &resBody)
	if err != nil {
		return xerrors.Errorf("convert error: %w", err)
	}

	generalPurchasingOrganizationPlant.PurchasingInfoRecord = resBody.D.PurchasingInfoRecord
	return nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
