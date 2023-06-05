package payment

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/account"
	"github.com/stripe/stripe-go/v74/accountlink"
)

func CreateConnectedAccount(secret_key, email, name, surname, card_token string) {

	// acct_1N0WYkPK9Zx8Cbgj

	stripe.Key = secret_key

	params := &stripe.AccountParams{
		Email: &email,
		Individual: &stripe.PersonParams{
			FirstName: &name,
			LastName:  &surname,
			DOB: &stripe.PersonDOBParams{
				Day:   stripe.Int64(27),
				Month: stripe.Int64(7),
				Year:  stripe.Int64(1999),
			},
		},
		BusinessType: stripe.String("individual"),
		ExternalAccount: &stripe.AccountExternalAccountParams{
			Currency:          stripe.String("USD"),
			AccountHolderName: stripe.String(fmt.Sprintf(name, surname)),
			Token:             &card_token,
		},
		Type:    stripe.String(string(stripe.AccountTypeCustom)),
		Country: stripe.String("TR"),
		Capabilities: &stripe.AccountCapabilitiesParams{
			// CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
			// 	Requested: stripe.Bool(true),
			// },
			Transfers: &stripe.AccountCapabilitiesTransfersParams{Requested: stripe.Bool(true)},
		},
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			ServiceAgreement: stripe.String("recipient"),
			Date:             stripe.Int64(1682975483),
			IP:               stripe.String("109.127.34.47"),
		},
	}

	result, err := account.New(params)

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println("response", result)
	}
}

func GetAccountLink(secret_key string) {

	account_id := createAccount(secret_key)

	account_link := createAccountLink(secret_key, account_id)

	fmt.Println("Account link", account_link)
}

func createAccount(secret_key string) string {
	stripe.Key = secret_key

	params := &stripe.AccountParams{
		// BusinessType: stripe.String("individual"),
		Type: stripe.String(string(stripe.AccountTypeCustom)),
		Capabilities: &stripe.AccountCapabilitiesParams{
			// CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
			// 	Requested: stripe.Bool(true),
			// },
			// BACSDebitPayments: &stripe.AccountCapabilitiesBACSDebitPaymentsParams{
			// 	Requested: stripe.Bool(true),
			// },
			Transfers: &stripe.AccountCapabilitiesTransfersParams{Requested: stripe.Bool(true)},
		},
		Country: stripe.String("TR"),
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			ServiceAgreement: stripe.String("recipient"),
			Date:             stripe.Int64(1682975483),
			IP:               stripe.String("109.127.34.47"),
		},
	}

	result, err := account.New(params)

	if err != nil {
		fmt.Printf("\n\n\n\n")
		fmt.Println("ERROR", err.Error())
		fmt.Printf("\n\n\n\n")
	} else {

		b, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return ""
		}

		_ = b
		// fmt.Println("Response", string(b))
	}

	return result.ID
}

func createAccountLink(secret_key, account_id string) string {
	stripe.Key = secret_key

	params := &stripe.AccountLinkParams{
		Account:    stripe.String(account_id),
		RefreshURL: stripe.String("https://moniesto.com/RefreshURL"),
		ReturnURL:  stripe.String("https://moniesto.com/ReturnURL"), // success
		Type:       stripe.String("account_onboarding"),
	}

	result, err := accountlink.New(params)

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {

		b, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return ""
		}

		fmt.Println("Response", string(b))
	}

	return result.URL
}

func DeleteConnectedAccount(secret_key, acc_id string) {
	stripe.Key = secret_key

	account.Del(acc_id, nil)
}

// ---------------- BINANCE
func CreateOrder(config config.Config) {
	base_url := "https://bpay.binanceapi.com"
	uri := "/binancepay/openapi/v2/order"

	full_url := base_url + uri

	// send request

	client := resty.New()

	body := binance.CreateOrderRequest{
		Env: binance.Env{
			TerminalType: "WEB",
		},
		MerchantTradeNo: core.CreatePlainID(),
		OrderAmount:     5,
		Currency:        "USDT",
		Goods: binance.Goods{
			GoodsType:        "02",
			GoodsCategory:    "0000",
			ReferenceGoodsId: "7876763A3B",
			GoodsName:        "Moniest 1 - A",
		},
	}

	headers := createHeader(body, config)

	fmt.Println("headers.BinancepaySignature", headers.BinancepaySignature)
	fmt.Println("headers.BinancepayNonce", headers.BinancepayNonce)

	resp, err := client.R().
		SetHeaders(map[string]string{
			"content-type":              headers.ContentType,
			"binancepay-timestamp":      strconv.Itoa(int(headers.BinancepayTimestamp)),
			"binancepay-nonce":          headers.BinancepayNonce,
			"binancepay-certificate-sn": headers.BinancepayCertificateSN,
			"binancepay-signature":      headers.BinancepaySignature, // BinancePay-Signature
		}).SetBody(body).Post(full_url)

	if err != nil {
		fmt.Println("Error Send Request: ", err.Error())
	}

	// b, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("Response:", resp)
}

func createHeader(body binance.CreateOrderRequest, config config.Config) binance.RequestHeader {
	ts := util.DateToTimestamp(time.Now())
	// nonce, _ :=  generateRandomStringURLSafe(20)

	nonce := "5RhaTrZPhknNv0kDSA2UQ67cPMVNS4sA"

	fmt.Println("nonce", nonce)

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return binance.RequestHeader{}
	}
	bodyInStr := string(b)

	fmt.Println("bodyInStr", bodyInStr)

	secretKey := config.BinanceSecretKey

	payload := strconv.Itoa(int(ts)) + "\n" + nonce + "\n" + bodyInStr + "\n"

	fmt.Println("payload", payload)

	// secretKey, _ = hex.DecodeString(secretKey)

	// hmac := cryptoHmac.New(sha256.New, []byte(secretKey))
	// hmac.Write([]byte(payload)) // TODO: maybe Write()

	// signature := hmac.Sum(nil)

	// signatureStr := strings.ToUpper(string(signature[:]))

	signature := signature(payload, secretKey)

	fmt.Println("signature", signature)

	return binance.RequestHeader{
		ContentType:             "application/json",
		BinancepayTimestamp:     ts,
		BinancepayNonce:         nonce,
		BinancepayCertificateSN: config.BinanceApiKey,
		BinancepaySignature:     signature,
	}
}

func signature(message, secret string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return strings.ToUpper(signingKey)
}

func generateRandomStringURLSafe(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
