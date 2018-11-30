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

package models

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type Disk struct {

    /* 云硬盘ID (Optional) */
    DiskId string `json:"diskId"`

    /* 云硬盘所属AZ (Optional) */
    Az string `json:"az"`

    /* 云硬盘名称 (Optional) */
    Name string `json:"name"`

    /* 云硬盘描述 (Optional) */
    Description string `json:"description"`

    /* 磁盘类型，取值为 ssd 或 premium-hdd (Optional) */
    DiskType string `json:"diskType"`

    /* 磁盘大小，单位为 GiB (Optional) */
    DiskSizeGB int `json:"diskSizeGB"`

    /* 云硬盘状态，取值为 creating、available、in-use、extending、restoring、deleting、deleted、error_create、error_delete、error_restore、error_extend 之一 (Optional) */
    Status string `json:"status"`

    /* 挂载信息 (Optional) */
    Attachments []DiskAttachment `json:"attachments"`

    /* 创建该云硬盘的快照ID (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* 创建云硬盘时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 云硬盘计费配置信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`
}
