package datastore

import (
	"github.com/upper/db/v4"
	"reflect"
)

const UserTable = "users"

func UserAdd(item User) (int, bool) {
	return Datastore.Add(UserTable, item)
}

func UserGetBy(key, value interface{}) (User, bool) {
	var item User
	ok := Datastore.GetBy(UserTable, &item, key, value)
	return item, ok
}
func UserGetByColumnOne(key, value, column, i interface{}) bool {
	return Datastore.GetByColumnOne(UserTable, key, value, column, i)
}
func UsersGetsBy(key, value interface{}) ([]User, bool) {
	var items []User
	ok := Datastore.GetsBy(UserTable, &items, key, value)
	return items, ok
}
func UserEdit(item User) bool {
	_, ok := UserGetBy("id", item.ID)
	if ok {
		return Datastore.EditBy(UserTable, item, "id", item.ID)
	}
	return false
}
func UserEditOne(key, value, field, i interface{}) bool {
	return Datastore.EditOne(UserTable, key, value, field, i)
}
func UsersGetAll() ([]User, bool) {
	var items []User
	ok := Datastore.GetAll(UserTable, &items)
	return items, ok
}

func UsersCount() (int, bool) {
	count, ok := Datastore.Count(UserTable)
	return count, ok
}
func UsersActiveCount() (int, bool) {
	return Datastore.CountByCond(UserTable, db.Cond{"is_active": true})
}

//func UsersGetsAllID(paginate int,tags []string) []int {
//	var (
//		users []User
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
//func UsersGetsAllByTagsID(tags []string) []int {
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
func (item *User) IsEmpty() bool { return reflect.DeepEqual(User{}, item) }
