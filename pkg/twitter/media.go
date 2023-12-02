package twitter

import (
	"context"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/moniesto/moniesto-be/config"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	tweetTypes "github.com/michimani/gotwi/tweet/managetweet/types"
)

// set API key as env var (for gotwi package)
func initAPIKey(config config.Config) error {
	err := os.Setenv("GOTWI_API_KEY", config.TwitterApiKey)
	if err != nil {
		return err
	}

	err = os.Setenv("GOTWI_API_KEY_SECRET", config.TwitterApiKeySecret)
	if err != nil {
		return err
	}

	return nil
}

func UploadMedia(config config.Config, base64String string) (mediaID string, err error) {
	api := anaconda.NewTwitterApiWithCredentials(config.TwitterAccessToken, config.TwitterAccessTokenSecret, config.TwitterApiKey, config.TwitterApiKeySecret)

	media, err := api.UploadMedia(base64String)
	if err != nil {
		return "", err
	}

	return media.MediaIDString, nil
}

func ShareTweet(config config.Config, content string, mediaIDs []string) error {
	if err := initAPIKey(config); err != nil {
		return err
	}

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           config.TwitterAccessToken,
		OAuthTokenSecret:     config.TwitterAccessTokenSecret,
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		return err
	}

	p := &tweetTypes.CreateInput{
		Text: gotwi.String(content),
	}

	if len(mediaIDs) > 0 {
		p.Media = &tweetTypes.CreateInputMedia{
			MediaIDs: mediaIDs,
		}
	}

	_, err = managetweet.Create(context.Background(), c, p)
	if err != nil {
		return err
	}

	return nil
}
