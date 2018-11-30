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
import disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"

type Instance struct {

    /* 云主机ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 云主机名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例规格 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 主网卡所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 主网卡所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 主网卡主IP地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 主网卡主IP绑定弹性IP的ID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 主网卡主IP绑定弹性IP的地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 云主机状态，<a href="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a> (Optional) */
    Status string `json:"status"`

    /* 云主机描述 (Optional) */
    Description string `json:"description"`

    /* 镜像ID (Optional) */
    ImageId string `json:"imageId"`

    /* 系统盘配置 (Optional) */
    SystemDisk InstanceDiskAttachment `json:"systemDisk"`

    /* 数据盘配置 (Optional) */
    DataDisks []InstanceDiskAttachment `json:"dataDisks"`

    /* 主网卡配置 (Optional) */
    PrimaryNetworkInterface InstanceNetworkInterfaceAttachment `json:"primaryNetworkInterface"`

    /* 辅助网卡配置 (Optional) */
    SecondaryNetworkInterfaces []InstanceNetworkInterfaceAttachment `json:"secondaryNetworkInterfaces"`

    /* 创建时间 (Optional) */
    LaunchTime string `json:"launchTime"`

    /* 云主机所在可用区 (Optional) */
    Az string `json:"az"`

    /* 密钥对名称 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 高可用组，如果创建云主机使用了高可用组，此处可展示高可用组名称 (Optional) */
    Ag Ag `json:"ag"`

    /* 高可用组中的错误域 (Optional) */
    FaultDomain string `json:"faultDomain"`

    /* Tag信息 (Optional) */
    Tags []disk.Tag `json:"tags"`
}
