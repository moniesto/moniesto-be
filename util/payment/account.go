package payment

import (
	"encoding/json"
	"fmt"

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
			// Requested: stripe.Bool(true),
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
		fmt.Println("ERROR", err.Error())
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
