package binance

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/util"
)

// requestWithBinanceHeader creates binance specific header [with secret, nonce, signature and etc.]
func requestWithBinanceHeader(body any, config config.Config) (*resty.Request, error) {
	client := resty.New()

	headers, err := createHeader(body, config)
	if err != nil {
		return nil, fmt.Errorf("error while creating binance specific header")
	}

	return client.R().SetHeaders(map[string]string{
		"content-type":              headers.ContentType,
		"binancepay-timestamp":      strconv.Itoa(int(headers.BinancepayTimestamp)),
		"binancepay-nonce":          headers.BinancepayNonce,
		"binancepay-certificate-sn": headers.BinancepayCertificateSN,
		"binancepay-signature":      headers.BinancepaySignature,
	}), nil
}

func createHeader(body any, config config.Config) (RequestHeader, error) {
	ts := util.DateToTimestamp(time.Now())

	nonce := generateNonce()

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return RequestHeader{}, fmt.Errorf("error while marshal body")
	}
	bodyInStr := string(b)

	secretKey := config.BinanceSecretKey

	payload := strconv.Itoa(int(ts)) + "\n" + nonce + "\n" + bodyInStr + "\n"

	signature := signature(payload, secretKey)

	return RequestHeader{
		ContentType:             "application/json",
		BinancepayTimestamp:     ts,
		BinancepayNonce:         nonce,
		BinancepayCertificateSN: config.BinanceApiKey,
		BinancepaySignature:     signature,
	}, nil
}

func signature(message, secret string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return strings.ToUpper(signingKey)
}

func generateNonce() string {
	return core.CreatePlainID()
}
