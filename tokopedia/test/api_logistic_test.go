/*
Tokopedia API

Testing LogisticAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tokopedia

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func Test_tokopedia_LogisticAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogisticAPIService GetActiveCourier", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.LogisticAPI.GetActiveCourier(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogisticAPIService GetShipmentInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.LogisticAPI.GetShipmentInfo(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogisticAPIService UpdateShipmentInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.LogisticAPI.UpdateShipmentInfo(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
