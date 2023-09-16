package example

import (
	"context"
	"fmt"

	"github.com/dhimas-sirclo/apiclient/tokopedia"
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
	client := tokopedia.NewAPIClient(tokopedia.NewConfiguration())
	authCtx := context.WithValue(context.Background(), tokopedia.ContextBasicAuth, tokopedia.BasicAuth{
		UserName: CLIENT_ID,
		Password: CLIENT_SECRET,
	})
	// authCtx = context.WithValue(authCtx, tokopedia.ContextServerIndex, 1)
	tokenResp, httpResp, err := client.AuthenticationAPI.
		Authentication(authCtx).
		GrantType("client_credentials").
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(tokenResp)

	maxRetry := 10
	maxDelayMs := 8000

	ctx := context.WithValue(context.Background(), tokopedia.ContextAccessToken, tokenResp.AccessToken)
	// Get product by product id
	resp, httpResp, err := client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		Product(PRODUCT_ID).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by product id with retry
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		Product(PRODUCT_ID).
		ExecuteWithRetry(maxRetry, maxDelayMs)
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by product url
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		ProductUrl(PRODUCT_ID).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by product url with Retry
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		ProductUrl(PRODUCT_ID).
		ExecuteWithRetry(maxRetry, maxDelayMs)
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by product sku
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		Sku(PRODUCT_SKU).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by product sku with retry
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		Sku(PRODUCT_SKU).
		ExecuteWithRetry(maxRetry, maxDelayMs)
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by shop id
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		ShopId(SHOP_ID).
		Page(1).
		PerPage(50).
		LastSort("lastSort").
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)

	// Get product by shop id
	resp, httpResp, err = client.ProductAPI.
		GetProductInfo(ctx, FS_ID).
		ShopId(SHOP_ID).
		Page(1).
		PerPage(50).
		LastSort("lastSort").
		ExecuteWithRetry(maxRetry, maxDelayMs)
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode >= 400 {
		panic(httpResp.Status)
	}
	fmt.Println(resp)
}
