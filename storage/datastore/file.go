package datastore

import (
	"github.com/upper/db/v4"
	"reflect"
)

const FileTable = "files"

func FileAdd(item File) (int, bool) {
	return Datastore.Add(FileTable, item)
}

func FileGetBy(key, value interface{}) (File, bool) {
	var item File
	ok := Datastore.GetBy(FileTable, &item, key, value)
	return item, ok
}
func FileGetByColumnOne(key, value, column, i interface{}) bool {
	return Datastore.GetByColumnOne(FileTable, key, value, column, i)
}
func FilesGetsBy(key, value interface{}) ([]File, bool) {
	var items []File
	ok := Datastore.GetsBy(FileTable, &items, key, value)
	return items, ok
}
func FileEdit(item File) bool {
	_, ok := FileGetBy("id", item.ID)
	if ok {
		return Datastore.EditBy(FileTable, item, "id", item.ID)
	}
	return false
}
func FileEditOne(key, value, field, i interface{}) bool {
	return Datastore.EditOne(FileTable, key, value, field, i)
}
func FilesGetAll() ([]File, bool) {
	var items []File
	ok := Datastore.GetAll(FileTable, &items)
	return items, ok
}

func FilesCount() (int, bool) {
	count, ok := Datastore.Count(FileTable)
	return count, ok
}
func FilesActiveCount() (int, bool) {
	return Datastore.CountByCond(FileTable, db.Cond{"is_active": true})
}

//func FilesGetsAllID(paginate int,tags []string) []int {
//	var (
//		users []File
//		IDs   []int
//	)
//	res := Datastore.GetCursor("users", "created", paginate, db.Cond{"is_active": true})/
//	total, _ := res.TotalPages()
//	_ = res.All(&users)
//
//	for _, u := range users {
//		IDs = append(IDs, u.ID)
//	}
//	for i := 1; i < int(total); i++ {
//		res = res.NextPage(users[len(users)-1].Created)
//		_ = res.All(&users)
//		for _, u := range users {
//			IDs = append(IDs, u.ID)
//		}
//	}
//	return IDs
//}
//func FilesGetsAllByTagsID(tags []string) []int {
//	var (
//		job Job
//		IDs []int
//	)
//
//	res := Datastore.GetLarge("users", "created", db.Cond{"is_active": true})
//	defer res.Close()
//
//	for res.Next(&job) {
//		if tagsCompare(job.Tags, tags) {
//			IDs = append(IDs, job.ID)
//		}
//	}
//	if err := res.Err(); err != nil {
//		log.Printf("ERROR: %v", err)
//		log.Fatalf(`SUGGESTION: change OrderBy("-ID") into OrderBy("id") and try again.`)
//	}
//	return IDs
//}
func (item *File) IsEmpty() bool { return reflect.DeepEqual(File{}, item) }
