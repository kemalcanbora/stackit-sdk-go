package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/argus"
)

type apiClientMocked struct {
	getFails      bool
	resourceState *string
	jobs          []argus.Job
}

func (a *apiClientMocked) GetInstanceExecute(_ context.Context, _, _ string) (*argus.InstanceResponse, error) {
	if a.getFails {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &argus.InstanceResponse{
		Id:     utils.Ptr("iid"),
		Status: a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetScrapeConfigsExecute(_ context.Context, _, _ string) (*argus.ScrapeConfigsResponse, error) {
	if a.getFails {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &argus.ScrapeConfigsResponse{
		Data: &a.jobs,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(CreateSuccess),
			wantErr:       false,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: utils.Ptr(CreateFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "broken_response",
			getFails:      false,
			resourceState: nil,
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *argus.InstanceResponse
			if !tt.getFails {
				wantRes = &argus.InstanceResponse{
					Id:     utils.Ptr("iid"),
					Status: tt.resourceState,
				}
			} else {
				wantRes = nil
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(UpdateSuccess),
			wantErr:       false,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: utils.Ptr(UpdateFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *argus.InstanceResponse
			if !tt.getFails {
				wantRes = &argus.InstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			} else {
				wantRes = nil
			}

			handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(DeleteSuccess),
			wantErr:       false,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: utils.Ptr(DeleteFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *argus.InstanceResponse
			if !tt.getFails {
				wantRes = &argus.InstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			} else {
				wantRes = nil
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []argus.Job
		wantErr  bool
	}{
		{
			desc:     "create_succeeded",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("job")}, {JobName: utils.Ptr("other-job")}},
			wantErr:  false,
		},
		{
			desc:     "create_failed and timeout",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []argus.Job{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *argus.ScrapeConfigsResponse
			if !tt.getFails {
				wantRes = &argus.ScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			} else {
				wantRes = nil
			}

			handler := CreateScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []argus.Job
		wantErr  bool
	}{
		{
			desc:     "delete_succeeded",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  false,
		},
		{
			desc:     "timeout",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("job")}},
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []argus.Job{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *argus.ScrapeConfigsResponse
			if !tt.getFails {
				wantRes = &argus.ScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			} else {
				wantRes = nil
			}

			handler := DeleteScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}