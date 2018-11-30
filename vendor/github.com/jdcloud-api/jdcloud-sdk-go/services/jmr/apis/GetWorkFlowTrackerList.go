// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type GetWorkFlowTrackerListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    SelectParams *jmr.SelectParams `json:"selectParams"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetWorkFlowTrackerListRequest(
    regionId string,
) *GetWorkFlowTrackerListRequest {

	return &GetWorkFlowTrackerListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/workFlowTracker:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param selectParams:  (Optional)
 */
func NewGetWorkFlowTrackerListRequestWithAllParams(
    regionId string,
    selectParams *jmr.SelectParams,
) *GetWorkFlowTrackerListRequest {

    return &GetWorkFlowTrackerListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/workFlowTracker:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SelectParams: selectParams,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetWorkFlowTrackerListRequestWithoutParam() *GetWorkFlowTrackerListRequest {

    return &GetWorkFlowTrackerListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/workFlowTracker:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetWorkFlowTrackerListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param selectParams: (Optional) */
func (r *GetWorkFlowTrackerListRequest) SetSelectParams(selectParams *jmr.SelectParams) {
    r.SelectParams = selectParams
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetWorkFlowTrackerListRequest) GetRegionId() string {
    return r.RegionId
}

type GetWorkFlowTrackerListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetWorkFlowTrackerListResult `json:"result"`
}

type GetWorkFlowTrackerListResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}