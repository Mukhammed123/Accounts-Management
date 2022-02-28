package requester

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/antchfx/xmlquery"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

type (
	hdhsSoapEnvelope struct {
		XMLName           xml.Name `xml:"soap:Envelope"`
		XMLSchema         string   `xml:"xmlns:xsd,attr"`
		XMLSchemaInstance string   `xml:"xmlns:xsi,attr"`
		XMLNamespace      string   `xml:"xmlns:soap,attr"`
		Body              hdhsSoapBody
	}

	hdhsSoapBody struct {
		XMLName xml.Name `xml:"soap:Body"`
		Content hdhsSoapContent
	}

	hdhsSoapContent struct {
		XMLName     xml.Name
		RequestBody []byte `xml:"strJson"`
	}
)

func generateRequestBodyForHDHS(methodName string, requestData interface{}) (io.Reader, error) {
	if marshaledData, err := json.Marshal(requestData); err != nil {
		return nil, err
	} else {
		soapRequest := hdhsSoapEnvelope{
			XMLSchema:         "http://www.w3.org/2001/XMLSchema",
			XMLSchemaInstance: "http://www.w3.org/2001/XMLSchema-instance",
			XMLNamespace:      "http://www.w3.org/2003/05/soap-envelope",
			Body: hdhsSoapBody{
				Content: hdhsSoapContent{
					XMLName: xml.Name{
						Local: methodName,
						Space: "http://tempuri.org/",
					},
					RequestBody: marshaledData,
				},
			},
		}
		if requestbody, err := xml.Marshal(soapRequest); err != nil {
			return nil, err
		} else {
			return bytes.NewBuffer(append([]byte(xml.Header), requestbody...)), nil
		}
	}
}

func getResponseDataFromHDHS(methodName string, responseBody io.Reader) (interface{}, error) {
	var responseData interface{}
	if document, err := xmlquery.Parse(responseBody); err != nil {
		return nil, err
	} else if node, err := xmlquery.Query(document, fmt.Sprintf(
		"//soap:Envelope/soap:Body/%sResponse/%sResult", methodName, methodName)); err != nil {
		return nil, err
	} else if err := json.Unmarshal([]byte(node.InnerText()), &responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}

func MakeRequestToHDHS(methodName string, requestData datatypes.JSONMap) (interface{}, error) {
	requestData["authKey"] = "1145010010"
	if requestBody, err := generateRequestBodyForHDHS(methodName, requestData); err != nil {
		return nil, err
	} else if request, err := http.NewRequest(http.MethodPost, serviceURLOfHDHS, requestBody); err != nil {
		return nil, err
	} else {
		request.Header.Set("Content-Type", "application/soap+xml")
		request.Header.Set("Access-Control-Allow-Origin", "*")
		// request.Header.Set("Content-Type", "application/soap+xml")
		if response, err := http.DefaultClient.Do(request); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else if responseData, err := getResponseDataFromHDHS(methodName, response.Body); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			defer response.Body.Close()
			if responseMap, ok := responseData.(map[string]interface{}); ok {
				if responseMap["isSuccess"] == "Y" {
					return responseData, nil
				} else if errorMessage, ok := responseMap["message"]; ok {
					return nil, echo.NewHTTPError(http.StatusNotFound, errorMessage)
				}
			}
			if marshaledResponseData, err := json.Marshal(responseData); err != nil {
				panic(err)
			} else {
				return nil, echo.NewHTTPError(response.StatusCode, string(marshaledResponseData))
			}
		}
	}
}
