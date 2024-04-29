package main

import (
	"bytes"
	"candies2/restapi/operations"
	"crypto/x509"
	"flag"
	"fmt"
	transport "github.com/go-openapi/runtime/client"
	"io"
	"net/http"
	"os"
)

func must(err error) {
	if err != nil {
		fmt.Printf("client error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	count := flag.Int64("c", 0, "candy count. Number of candies to buy")
	name := flag.String("k", "", "candy kind. Short name of the candy to buy. One of: CE AA NT DE YR")
	money := flag.Int64("m", 0, "money. Amount of money you want to put into the machine")

	flag.Parse()
	if count == nil || (*count) == 0 ||
		name == nil || (*name) == "" ||
		money == nil || (*money) == 0 {
		fmt.Println("Error parsing the flags")
		flag.Usage()
		return
	}
	params := operations.NewBuyCandyParams()
	params.Order.CandyCount = count
	params.Order.Money = money
	params.Order.CandyType = name

	tlsConfig, err := transport.TLSClientAuth(transport.TLSClientOptions{
		Certificate: "./certs/candies/cert.pem",
		Key:         "./certs/candies/key.pem",
		CA:          "./certs/minica.pem",
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			return nil
		},
	})
	must(err)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	body, err := params.Order.MarshalBinary()
	must(err)
	resp, err := client.Post("https://candies:8081/buy_candy", "application/json", bytes.NewReader(body))
	must(err)

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)

	switch resp.StatusCode {
	case operations.BuyCandyCreatedCode:
		{
			success := operations.BuyCandyCreatedBody{}
			must(success.UnmarshalBinary(respBody))
			fmt.Println(success.Thanks)
			if success.Change > 0 {
				fmt.Printf("Your change is %d\n", success.Change)
			}
		}
	case operations.BuyCandyBadRequestCode:
	case operations.BuyCandyPaymentRequiredCode:
		{
			fail := operations.BuyCandyBadRequestBody{}
			must(fail.UnmarshalBinary(respBody))
			fmt.Println(fail.Error)
		}
	default:
		fmt.Println(string(respBody))

	}

}

//func main() {
//	tlsConfig, err := transport.TLSClientAuth(transport.TLSClientOptions{
//		Certificate: "./certs/candies/cert.pem",
//		Key:         "./certs/candies/key.pem",
//		CA:          "./certs/minica.pem",
//		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
//			return nil
//		},
//	})
//
//	if err != nil {
//		fmt.Println("GET tls err: ", err.Error())
//	}
//
//	httpsClient := http.Client{
//		Transport: &http.Transport{
//			TLSClientConfig: tlsConfig,
//		},
//	}
//
//	cName := "AA"
//	cCount := int64(3)
//	cMoney := int64(60)
//
//	candyTransport := transport.NewWithClient("candies:8081", "", nil, &httpsClient)
//	candyClient := client.New(candyTransport, strfmt.Default)
//
//	requestBody := operations.NewBuyCandyParams()
//	requestBody.WithOrder(operations.BuyCandyBody{
//		CandyCount: &cCount,
//		CandyType:  &cName,
//		Money:      &cMoney,
//	})
//
//	resp, err := candyClient.Operations.BuyCandy(requestBody)
//
//	if err != nil {
//		fmt.Printf("Client err: %v", err.Error())
//	}
//
//	if resp != nil {
//		fmt.Println(resp.String())
//	}
//}
