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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    oss "github.com/jdcloud-api/jdcloud-sdk-go/services/oss/apis"
    "encoding/json"
    "errors"
)

type OssClient struct {
    core.JDCloudClient
}

func NewOssClient(credential *core.Credential) *OssClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("oss.jdcloud-api.com")

    return &OssClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "oss",
            Revision:    "0.2.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *OssClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *OssClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建bucket
 */
func (c *OssClient) PutBucket(request *oss.PutBucketRequest) (*oss.PutBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oss.PutBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除一个bucket
 */
func (c *OssClient) DeleteBucket(request *oss.DeleteBucketRequest) (*oss.DeleteBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oss.DeleteBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询bucket是否存在
 */
func (c *OssClient) HeadBucket(request *oss.HeadBucketRequest) (*oss.HeadBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oss.HeadBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 列出当前用户的所有bucket
 */
func (c *OssClient) ListBuckets(request *oss.ListBucketsRequest) (*oss.ListBucketsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oss.ListBucketsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

