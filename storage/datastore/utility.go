package datastore

import "schoolXfinalback/utility"

//func contains(s []string, e string) bool {
//	for _, a := range s {
//		if a == e {
//			return true
//		}
//	}
//	return false
//}

func tagsCompare(first, second []string) bool {
	for _, f := range first {
		if utility.Contains(second, f) {
			return true
		}
	}
	return false
}
