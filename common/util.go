/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package common

/*
 * 统计条件为 true 的数量
 * 例如 CountTrue(1 == 1, 2 < 3, false, true) = 3
 */
func CountTrue(conditions ...bool) int {
	var count = 0
	for _, condition := range conditions {
		if condition {
			count++
		}
	}
	return count
}
