/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resmanager

// TODO: add UT, assign @stormgbs

// var commonTestFile = "test_common_file"
//
// type reconcileInfo struct {
// 	desc      string
// 	resources []ResourceUpdater
// 	expect    []ResourceUpdater
// }
//
// func Test_UpdateBatch(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		resources []ResourceUpdater
// 	}{
// 		{
// 			name: "test_update_valid",
// 			resources: []ResourceUpdater{
// 				NewCommonCgroupResourceUpdater(GroupOwnerRef("root"), "/", system.CPUShares, "1024"),
// 				NewCommonResourceUpdater(commonTestFile, "19"),
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
//
// 		t.Run(tt.name, func(t *testing.T) {
// 			helper := system.NewFileTestUtil(t)
// 			defer helper.Cleanup()
//
// 			helper.CreateCgroupFile("/", system.CPUShares)
// 			helper.CreateFile(commonTestFile)
//
// 			t.Logf("Cur CgroupFile filepath %v", system.Conf.CgroupRootDir)
//
// 			rm := NewResourceUpdateExecutor("test", 1)
// 			stop := make(chan struct{})
// 			rm.Run(stop)
// 			defer func() { stop <- struct{}{} }()
// 			rm.UpdateBatch(tt.resources...)
// 			got := getActualResources(tt.resources)
// 			equalResourceMap(t, tt.resources, got, "checkCurrentResource")
// 		})
//
// 	}
// }
//
// func Test_UpdateBatchByCache(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		initCache      []ResourceUpdater
// 		initFiles      []ResourceUpdater
// 		reconcileInfos []reconcileInfo
// 	}{
// 		{
// 			name: "test_cache_equal_but_force_update",
// 			initCache: []ResourceUpdater{
// 				&CgroupResourceUpdater{ParentDir: "/", file: system.CPUShares, value: "1024", lastUpdateTimestamp: time.Now().Add(-5 * time.Second), updateFunc: CommonCgroupUpdateFunc},
// 				&CommonResourceUpdater{key: commonTestFile, file: commonTestFile, value: "19", lastUpdateTimestamp: time.Now().Add(-5 * time.Second), updateFunc: CommonUpdateFunc},
// 			},
// 			initFiles: []ResourceUpdater{
// 				&CgroupResourceUpdater{ParentDir: "/", file: system.CPUShares, value: "2048", lastUpdateTimestamp: time.Now().Add(-5 * time.Second), updateFunc: CommonCgroupUpdateFunc},
// 				&CommonResourceUpdater{key: commonTestFile, file: commonTestFile, value: "20", lastUpdateTimestamp: time.Now().Add(-5 * time.Second), updateFunc: CommonUpdateFunc},
// 			},
// 			reconcileInfos: []reconcileInfo{
// 				{
// 					desc: "test_update",
// 					resources: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 						NewCommonResourceUpdater(commonTestFile, "19"),
// 					},
// 					expect: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 						NewCommonResourceUpdater(commonTestFile, "19"),
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "test_cache_equal_and_not_forceUpdate",
// 			initCache: []ResourceUpdater{
// 				&CgroupResourceUpdater{ParentDir: "/", file: system.CPUShares, value: "1024", lastUpdateTimestamp: time.Now(), updateFunc: CommonCgroupUpdateFunc},
// 				&CommonResourceUpdater{key: commonTestFile, file: commonTestFile, value: "19", lastUpdateTimestamp: time.Now(), updateFunc: CommonUpdateFunc},
// 			},
// 			initFiles: []ResourceUpdater{
// 				&CgroupResourceUpdater{ParentDir: "/", file: system.CPUShares, value: "2048", lastUpdateTimestamp: time.Now(), updateFunc: CommonCgroupUpdateFunc},
// 				&CommonResourceUpdater{key: commonTestFile, file: commonTestFile, value: "20", lastUpdateTimestamp: time.Now().Add(-5 * time.Second), updateFunc: CommonUpdateFunc},
// 			},
// 			reconcileInfos: []reconcileInfo{
// 				{
// 					desc: "test_update",
// 					resources: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 						NewCommonResourceUpdater(commonTestFile, "19"),
// 					},
// 					expect: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "2048"),
// 						NewCommonResourceUpdater(commonTestFile, "20"),
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:      "test_reconcile",
// 			initCache: []ResourceUpdater{},
// 			initFiles: []ResourceUpdater{
// 				NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "2"),
// 			},
// 			reconcileInfos: []reconcileInfo{
// 				{
// 					desc: "test_start",
// 					resources: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 					},
// 					expect: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 					},
// 				},
// 				{
// 					desc: "test_running_2",
// 					resources: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "2"),
// 					},
// 					expect: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "2"),
// 					},
// 				},
// 				{
// 					desc: "test_running_3",
// 					resources: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 					},
// 					expect: []ResourceUpdater{
// 						NewCommonCgroupResourceUpdater(PodOwnerRef("", "pod1"), "/", system.CPUShares, "1024"),
// 					},
// 				},
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			helper := system.NewFileTestUtil(t)
// 			// defer helper.Cleanup()
//
// 			prepareResourceFiles(helper, tt.initFiles)
//
// 			resourceCache := cache.NewCache(time.Second, time.Second)
// 			for _, resource := range tt.initCache {
// 				resourceCache.Set(resource.Key(), resource, resource.GetLastUpdateTimestamp().Sub(time.Now())+time.Second)
// 			}
//
// 			rm := ResourceUpdateExecutor{name: tt.name, forceUpdateSeconds: 1, resourceCache: resourceCache, locker: &sync.Mutex{}}
// 			stop := make(chan struct{})
// 			rm.Run(stop)
// 			defer func() { stop <- struct{}{} }()
//
// 			for _, info := range tt.reconcileInfos {
// 				rm.UpdateBatchByCache(info.resources...)
// 				got := getActualResources(info.resources)
// 				equalResourceMap(t, info.resources, got, fmt.Sprintf("case:%s,checkCurrentResource", info.desc))
// 			}
// 		})
// 	}
// }
//
// func prepareResourceFiles(helper *system.FileTestUtil, initFiles []ResourceUpdater) {
// 	for _, resource := range initFiles {
// 		var err error
// 		switch resource.(type) {
// 		case *CommonResourceUpdater:
// 			helper.CreateFile(resource.Key())
// 			err = system.CommonFileWrite(resource.Key(), resource.Value())
// 		case *CgroupResourceUpdater:
// 			cgroupResource := resource.(*CgroupResourceUpdater)
// 			helper.CreateCgroupFile(cgroupResource.ParentDir, cgroupResource.file)
// 			err = system.CgroupFileWrite(cgroupResource.ParentDir, cgroupResource.file, resource.Value())
// 		default:
// 			err = fmt.Errorf("unknown resource type %T", resource)
// 		}
// 		if err != nil {
// 			klog.Errorf("prepareResourceFiles failed for resource %v, err: %s", resource, err)
// 		}
// 	}
// }
//
// func getActualResources(expect []ResourceUpdater) map[string]ResourceUpdater {
// 	got := make(map[string]ResourceUpdater)
//
// 	for _, resource := range expect {
// 		var value string
// 		var err error
// 		gotResource := resource.Clone()
// 		switch gotResource.(type) {
// 		case *CommonResourceUpdater:
// 			value, err = system.CommonFileRead(resource.Key())
// 			if err != nil { // abort set value when file read failed
// 				klog.Errorf("getActualResources failed for common resource %s, err: %s", resource.Key(), err)
// 				continue
// 			}
// 		case *CgroupResourceUpdater:
// 			cgroupResource := gotResource.(*CgroupResourceUpdater)
// 			value, err = system.CgroupFileRead(cgroupResource.ParentDir, cgroupResource.file)
// 			if err != nil { // abort set value when file read failed
// 				klog.Errorf("getActualResources failed for cgroup resource %s, err: %s", resource.Key(), err)
// 				continue
// 			}
// 		default:
// 			klog.Errorf("getActualResources failed for unknown resource %v, type %T", resource, resource)
// 			continue
// 		}
// 		gotResource.SetValue(value)
// 		got[gotResource.Key()] = gotResource
// 	}
// 	return got
// }
//
// func equalResourceMap(t *testing.T, expect []ResourceUpdater, got map[string]ResourceUpdater, msg string) {
// 	if len(expect) != len(got) {
// 		t.Errorf("msg:%s,checkResources fail! len not equal! expect: %+v,but got %+v", msg, expect, got)
// 		return
// 	}
// 	for _, resource := range expect {
// 		gotResource, exist := got[resource.Key()]
// 		if !exist {
// 			t.Errorf("msg:%s,checkResources fail! expect: %+v, but got nil", msg, resource)
// 			return
// 		}
// 		assert.Equal(t, resource.Value(), gotResource.Value(), msg)
// 	}
// }
