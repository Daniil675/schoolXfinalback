package datastore

import (
	"github.com/upper/db/v4"
	"reflect"
)

const ImageTable = "images"

func ImageAdd(item Image) (int, bool) {
	return Datastore.Add(ImageTable, item)
}

func ImageGetBy(key, value interface{}) (Image, bool) {
	var item Image
	ok := Datastore.GetBy(ImageTable, &item, key, value)
	return item, ok
}
func ImageGetByColumnOne(key, value, column, i interface{}) bool {
	return Datastore.GetByColumnOne(ImageTable, key, value, column, i)
}
func ImagesGetsBy(key, value interface{}) ([]Image, bool) {
	var items []Image
	ok := Datastore.GetsBy(ImageTable, &items, key, value)
	return items, ok
}
func ImageEdit(item Image) bool {
	_, ok := ImageGetBy("id", item.ID)
	if ok {
		return Datastore.EditBy(ImageTable, item, "id", item.ID)
	}
	return false
}
func ImageEditOne(key, value, field, i interface{}) bool {
	return Datastore.EditOne(ImageTable, key, value, field, i)
}
func ImagesGetAll() ([]Image, bool) {
	var items []Image
	ok := Datastore.GetAll(ImageTable, &items)
	return items, ok
}
func ImageDeleteBy(key, value interface{}) bool {
	ok := Datastore.DeleteBy(UserTable, db.Cond{key: value})
	return ok
}
func ImagesCount() (int, bool) {
	count, ok := Datastore.Count(ImageTable)
	return count, ok
}
func ImagesActiveCount() (int, bool) {
	return Datastore.CountByCond(ImageTable, db.Cond{"is_active": true})
}

//func ImagesGetsAllID(paginate int,tags []string) []int {
//	var (
//		users []Image
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
//func ImagesGetsAllByTagsID(tags []string) []int {
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
func (item *Image) IsEmpty() bool { return reflect.DeepEqual(Image{}, item) }
