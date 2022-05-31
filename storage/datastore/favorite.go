package datastore

import (
	"github.com/upper/db/v4"
	"reflect"
)

const FavoriteTable = "favorites"

func FavoriteAdd(item Favorite) (int, bool) {
	return Datastore.Add(FavoriteTable, item)
}

func FavoriteGetBy(key, value interface{}) (Favorite, bool) {
	var item Favorite
	ok := Datastore.GetBy(FavoriteTable, &item, key, value)
	return item, ok
}
func FavoriteGetByColumnOne(key, value, column, i interface{}) bool {
	return Datastore.GetByColumnOne(FavoriteTable, key, value, column, i)
}
func FavoritesGetsBy(key, value interface{}) ([]Favorite, bool) {
	var items []Favorite
	ok := Datastore.GetsBy(FavoriteTable, &items, key, value)
	return items, ok
}

//func FavoriteEdit(item Favorite) bool {
//	_, ok := FavoriteGetBy("id", item.ID)
//	if ok {
//		return Datastore.EditBy(FavoriteTable, item, "id", item.ID)
//	}
//	return false
//}
func FavoriteEditOne(key, value, field, i interface{}) bool {
	return Datastore.EditOne(FavoriteTable, key, value, field, i)
}
func FavoritesGetAll() ([]Favorite, bool) {
	var items []Favorite
	ok := Datastore.GetAll(FavoriteTable, &items)
	return items, ok
}

func FavoritesCount() (int, bool) {
	count, ok := Datastore.Count(FavoriteTable)
	return count, ok
}
func FavoritesActiveCount() (int, bool) {
	return Datastore.CountByCond(FavoriteTable, db.Cond{"is_active": true})
}

//func FavoritesGetsAllID(paginate int,tags []string) []int {
//	var (
//		users []Favorite
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
//func FavoritesGetsAllByTagsID(tags []string) []int {
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
func (item *Favorite) IsEmpty() bool { return reflect.DeepEqual(Favorite{}, item) }
