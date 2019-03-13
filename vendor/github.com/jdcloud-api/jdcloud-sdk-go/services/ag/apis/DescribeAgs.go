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
    ag "github.com/jdcloud-api/jdcloud-sdk-go/services/ag/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeAgsRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* agName - ag名字，支持模糊匹配
agId - agId，精确匹配
instanceTemplateId - 实例模板id，精确匹配
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAgsRequest(
    regionId string,
) *DescribeAgsRequest {

	return &DescribeAgsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availabilityGroups",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: agName - ag名字，支持模糊匹配
agId - agId，精确匹配
instanceTemplateId - 实例模板id，精确匹配
 (Optional)
 */
func NewDescribeAgsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeAgsRequest {

    return &DescribeAgsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAgsRequestWithoutParam() *DescribeAgsRequest {

    return &DescribeAgsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *DescribeAgsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeAgsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeAgsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: agName - ag名字，支持模糊匹配
agId - agId，精确匹配
instanceTemplateId - 实例模板id，精确匹配
(Optional) */
func (r *DescribeAgsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAgsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAgsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAgsResult `json:"result"`
}

type DescribeAgsResult struct {
    Ags []ag.AvailabilityGroup `json:"ags"`
    TotalCount int `json:"totalCount"`
}
