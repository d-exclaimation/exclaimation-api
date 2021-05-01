//
//  isolate.go
//  pipes
//
//  Created by d-exclaimation on 6:58 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package pipes

// IsolateStringPipe isolates String from one and another
type IsolateStringPipe func(row string) bool
type RowStringReducer func(prev []string, curr string) []string

// IsolateReducer create reducer given all the pipes
func IsolateReducer(isolations []IsolateStringPipe) RowStringReducer {
	return func(prev []string, curr string) []string {
		if len(prev) == 0 {
			return []string{curr}
		}
		last := prev[len(prev) - 1]
		for _, iso := range isolations {
			if iso(curr) || iso(last) {
				return append(prev, curr)
			}
		}
		return append(prev[:len(prev) - 1], last + "\n" + curr)
	}
}
