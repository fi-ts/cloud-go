package cloudgo

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fi-ts/cloud-go/api/client/cluster"
	"github.com/fi-ts/cloud-go/api/models"
	"github.com/stretchr/testify/assert"
)

func TestAttemptGenericError(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		priorErrorRegex string
		newErrorString  string
	}{
		{
			name: "swagger error gets printed nicely",
			err: &cluster.FindClusterDefault{
				Payload: &models.HttperrorsHTTPErrorResponse{
					Message:    strPtr("some error message"),
					Statuscode: int32Ptr(500),
				},
			},
			priorErrorRegex: regexp.QuoteMeta("[GET /v1/cluster/{id}][0] findCluster default  &{Message:") + ".* Statuscode:.*" + regexp.QuoteMeta("}"),
			newErrorString:  "some error message (500)",
		},
		{
			name:            "non swagger error gets passed through",
			err:             fmt.Errorf("non-swagger-error"),
			priorErrorRegex: "non-swagger-error",
			newErrorString:  "non-swagger-error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := regexp.MustCompile(tt.priorErrorRegex)
			assert.Regexp(t, r, tt.err.Error())

			err := AttemptGenericError(tt.err)
			assert.Equal(t, err.Error(), tt.newErrorString)
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
