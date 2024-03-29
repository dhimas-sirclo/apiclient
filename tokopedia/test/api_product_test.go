/*
Tokopedia API

Testing ProductAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tokopedia

import (
	"context"
	"testing"

	openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_tokopedia_ProductAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductAPIService GetProductInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ProductAPI.GetProductInfo(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService GetVariantsByCategoryId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ProductAPI.GetVariantsByCategoryId(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService GetVariantsByProductId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsId int64
		var productId int64

		resp, httpRes, err := apiClient.ProductAPI.GetVariantsByProductId(context.Background(), fsId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
