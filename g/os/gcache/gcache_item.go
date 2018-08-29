// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gcache

import "gitee.com/johng/gf/g/os/gtime"

// 判断缓存项是否已过期
func (item *CacheItem) IsExpired() bool {
    if item.e > gtime.Millisecond() {
        return false
    }
    return true
}