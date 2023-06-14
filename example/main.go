package example

import (
	"context"
	"fmt"

	"github.com/dhimas-sirclo/apiclient/tokopedia"
	"github.com/dhimas-sirclo/apiclient/tokopediaauth"
)

const (
	CLIENT_ID     = "clientID"
	CLIENT_SECRET = "clientSecret"
	FS_ID         = 12345
	PRODUCT_ID    = "12345"
)

func main() {
	authclient := tokopediaauth.NewAPIClient(tokopediaauth.NewConfiguration())
	ctx := context.WithValue(context.Background(), tokopediaauth.ContextBasicAuth, tokopediaauth.BasicAuth{
		UserName: CLIENT_ID,
		Password: CLIENT_SECRET,
	})
	tokenRes, httpRes, err := authclient.DefaultApi.TokenPost(ctx).Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(tokenRes)

	client := tokopedia.NewAPIClient(tokopedia.NewConfiguration())
	ctx = context.WithValue(context.Background(), tokopedia.ContextAccessToken, tokenRes.AccessToken)
	res, httpRes, err := client.DefaultApi.InventoryV1FsFsIdProductInfoGet(ctx, FS_ID).Product(PRODUCT_ID).Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(res)
}
