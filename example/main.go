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
	PRODUCT_SKU   = "SKU"
	SHOP_ID       = 12345
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
	// Get product by product id
	res, httpRes, err := client.DefaultApi.InventoryV1FsFsIdProductInfoGet(ctx, FS_ID).
		Product(PRODUCT_ID).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(res)

	// Get product by product url
	res, httpRes, err = client.DefaultApi.InventoryV1FsFsIdProductInfoGet(ctx, FS_ID).
		ProductUrl(PRODUCT_ID).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(res)

	// Get product by product sku
	res, httpRes, err = client.DefaultApi.InventoryV1FsFsIdProductInfoGet(ctx, FS_ID).
		Sku(PRODUCT_SKU).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(res)

	// Get product by shop id
	res, httpRes, err = client.DefaultApi.InventoryV1FsFsIdProductInfoGet(ctx, FS_ID).
		ShopId(SHOP_ID).
		Page(1).
		PerPage(50).
		LastSort("lastSort").
		Execute()
	if err != nil {
		panic(err)
	}
	if httpRes.StatusCode >= 400 {
		panic(httpRes.Status)
	}
	fmt.Println(res)
}
