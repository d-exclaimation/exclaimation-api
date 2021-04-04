//
//  reduce.go
//  slice
//
//  Created by d-exclaimation on 7:11 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package slice

func ReduceStr(some []string, reducer func(prev []string, curr string) []string) []string {
	res := make([]string, 0)
	for _, val := range some {
		res = reducer(res, val)
	}
	return res
}
