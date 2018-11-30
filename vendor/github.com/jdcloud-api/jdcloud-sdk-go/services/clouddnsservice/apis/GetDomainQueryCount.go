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
)

type GetDomainQueryCountRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID  */
    DomainId string `json:"domainId"`

    /* 查询的域名  */
    DomainName string `json:"domainName"`

    /* 起始时间, UTC时间例如2017-11-10T23:00:00Z  */
    Start string `json:"start"`

    /* 终止时间, UTC时间例如2017-11-10T23:00:00Z  */
    End string `json:"end"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID (Required)
 * param domainName: 查询的域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainQueryCountRequest(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *GetDomainQueryCountRequest {

	return &GetDomainQueryCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID (Required)
 * param domainName: 查询的域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 */
func NewGetDomainQueryCountRequestWithAllParams(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *GetDomainQueryCountRequest {

    return &GetDomainQueryCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainQueryCountRequestWithoutParam() *GetDomainQueryCountRequest {

    return &GetDomainQueryCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetDomainQueryCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID(Required) */
func (r *GetDomainQueryCountRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param domainName: 查询的域名(Required) */
func (r *GetDomainQueryCountRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainQueryCountRequest) SetStart(start string) {
    r.Start = start
}

/* param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainQueryCountRequest) SetEnd(end string) {
    r.End = end
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainQueryCountRequest) GetRegionId() string {
    return r.RegionId
}

type GetDomainQueryCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainQueryCountResult `json:"result"`
}

type GetDomainQueryCountResult struct {
    Time []int `json:"time"`
    Traffic []int `json:"traffic"`
}