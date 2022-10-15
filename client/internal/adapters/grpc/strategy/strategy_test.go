package strategy

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"client/internal/entities"
)

func TestStrategy_Execute(t *testing.T) {
	TestTable := []struct {
		payload      *entities.Request
		expectedResp *entities.Response
		expectedErr  error
		testName     string
	}{
		{
			payload: &entities.Request{},
			expectedResp: &entities.Response{
				Version:     "test Version",
				ExecuteDate: "test ExecuteDate",
			},
			expectedErr: nil,
			testName:    "test 1: Ok",
		},
		{
			payload: &entities.Request{
				Name: "err",
			},
			expectedErr: fmt.Errorf("rpc error: code = Unknown desc = execute call err"),
			testName:    "test 2: Err",
		},
	}

	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			mockSrv := NewMockStrategy()
			resp, err := mockSrv.Execute(context.Background(), testCase.payload)
			assert.Equal(t, testCase.expectedResp, resp)
			if testCase.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, testCase.expectedErr.Error(), err.Error())
			}
		})
	}
}
