package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var body = `<SOAP>
	<SOAP-ENV:Header xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" soap:mustUnderstand="1">
			<wsse:UsernameToken>
				<wsse:Username>tm-siteminder-1378</wsse:Username>
				<wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">HjKmQpVrXsZyTdWf</wsse:Password>
			</wsse:UsernameToken>
		</wsse:Security>
	</SOAP-ENV:Header>
	<SOAP-ENV:Body xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<OTA_HotelAvailRQ xmlns="http://www.opentravel.org/OTA/2003/05" AvailRatesOnly="true" EchoToken="89dcd3e3-3dc3-4c96-a587-63e493c32c6a" TimeStamp="2025-02-11T16:20:58+11:00" Version="1.0">
			<AvailRequestSegments>
				<AvailRequestSegment AvailReqType="Room">
					<HotelSearchCriteria>
						<Criterion>
							<HotelRef HotelCode="1736908455"/>
						</Criterion>
					</HotelSearchCriteria>
				</AvailRequestSegment>
			</AvailRequestSegments>
		</OTA_HotelAvailRQ>
	</SOAP-ENV:Body>
</SOAP>`

func main() {
	endpoint := "https://cms-push-svc.tourmind.cn:7443/siteminder/api/v1/ota/retrieve/rooms"

	payload := strings.NewReader(body)
	req, _ := http.NewRequest("POST", endpoint, payload)

	req.Header.Set("Content-Type", "application/xml")

	cli := http.Client{
		Timeout: time.Second * 20,
	}
	resp, err := cli.Do(req)
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)
	log.Println(string(content))
}
