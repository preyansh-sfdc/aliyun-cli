package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ReservedInstance is a nested struct in ecs response
type ReservedInstance struct {
	Status               string                                    `json:"Status" xml:"Status"`
	CreationTime         string                                    `json:"CreationTime" xml:"CreationTime"`
	ReservedInstanceName string                                    `json:"ReservedInstanceName" xml:"ReservedInstanceName"`
	ReservedInstanceId   string                                    `json:"ReservedInstanceId" xml:"ReservedInstanceId"`
	InstanceType         string                                    `json:"InstanceType" xml:"InstanceType"`
	InstanceAmount       int                                       `json:"InstanceAmount" xml:"InstanceAmount"`
	RegionId             string                                    `json:"RegionId" xml:"RegionId"`
	OfferingType         string                                    `json:"OfferingType" xml:"OfferingType"`
	StartTime            string                                    `json:"StartTime" xml:"StartTime"`
	Description          string                                    `json:"Description" xml:"Description"`
	AllocationStatus     string                                    `json:"AllocationStatus" xml:"AllocationStatus"`
	ExpiredTime          string                                    `json:"ExpiredTime" xml:"ExpiredTime"`
	ResourceGroupId      string                                    `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ZoneId               string                                    `json:"ZoneId" xml:"ZoneId"`
	Platform             string                                    `json:"Platform" xml:"Platform"`
	Scope                string                                    `json:"Scope" xml:"Scope"`
	OperationLocks       OperationLocksInDescribeReservedInstances `json:"OperationLocks" xml:"OperationLocks"`
	Tags                 TagsInDescribeReservedInstances           `json:"Tags" xml:"Tags"`
}
