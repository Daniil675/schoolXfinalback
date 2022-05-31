package datastore

import (
	"github.com/upper/db/v4"
	"schoolXfinalback/utility"
)

func (d *DatastoreT) Add(tablename string, i interface{}) (int, bool) {
	id, err := d.Collection(tablename).Insert(i)
	if err != nil {
		utility.Error(err, "Add")
		return 0, false
	}
	return int(id.ID().(int64)), true
}
func (d *DatastoreT) AddWithoutID(tablename string, i interface{}) bool {
	_, err := d.Collection(tablename).Insert(i)
	if err != nil {
		utility.Error(err, "AddWithoutID")
		return false
	}
	return true
}

func (d *DatastoreT) GetBy(tablename string, i interface{}, key, value interface{}) bool {
	err := d.Collection(tablename).Find(db.Cond{key: value}).One(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetBy")
	}
	return true
}
func (d *DatastoreT) GetByCond(tablename string, i interface{}, cond db.Cond) bool {
	err := d.Collection(tablename).Find(cond).One(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetByCond")
	}
	return true
}

func (d *DatastoreT) GetLast(tablename string, i interface{}) bool {
	err := d.Collection(tablename).Find().OrderBy("-id").Limit(1).One(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetLast")
	}
	return true
}
func (d *DatastoreT) GetByColumn(tablename string, key, value interface{}, column, i []interface{}) bool {
	row, err := d.SQL().Select(column...).From(tablename).Where(key, value).QueryRow()
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetByColumn")
	}
	row.Scan(i...)
	return true
}
func (d *DatastoreT) GetByColumnOne(tablename string, key, value, column, i interface{}) bool {
	row, err := d.SQL().Select(column).From(tablename).Where(key, value).QueryRow()
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetByColumnOne")
	}
	row.Scan(i)
	return true
}
func (d *DatastoreT) GetsByCond(table string, i interface{}, cond db.Cond) bool {
	err := d.Collection(table).Find(cond).All(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetsByCond")
	}
	return true
}
func (d *DatastoreT) GetsBy(tablename string, i interface{}, key, value interface{}) bool {
	err := d.Collection(tablename).Find(db.Cond{key: value}).All(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetsBy")
	}
	return true
}
func (d *DatastoreT) GetByCondOrder(table string, i interface{}, cond db.Cond, order string) bool {
	err := d.Collection(table).Find(cond).OrderBy(order).All(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetByCondOrder")
	}
	return true
}
func (d *DatastoreT) GetByCondOrderLimit(table string, i interface{}, cond db.Cond, order string, limit int) bool {
	err := d.Collection(table).Find(cond).OrderBy(order).Limit(limit).All(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetByCondOrderLimit")
	}
	return true
}

func (d *DatastoreT) EditOne(tablename string, key, value, field, i interface{}) bool {
	_, err := d.SQL().Update(tablename).
		Set(field, i).
		Where(key, value).
		Exec()
	if err != nil {
		utility.Error(err, "EditOne")
		return false
	}
	return true
}

func (d *DatastoreT) Edit(tablename string, key, value, i interface{}) bool {
	res := d.Collection(tablename).Find(db.Cond{key: value})
	err := res.Update(i)
	if err != nil {
		utility.Error(err, "Edit")
		return false
	}
	if err := res.Close(); err != nil {
		utility.Error(err, "Edit")
	}
	return true
}
func (d *DatastoreT) EditByCond(tablename string, field, i interface{}, cond db.Cond) bool {
	_, err := d.SQL().Update(tablename).
		Set(field, i).
		Where(cond).
		Exec()
	if err != nil {
		utility.Error(err, "EditByCond")
		return false
	}
	return true
}
func (d *DatastoreT) GetAll(tablename string, i interface{}) bool {
	err := d.Collection(tablename).Find().All(i)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return false
		}
		utility.Error(err, "GetAll")
	}
	return true
}

func (d *DatastoreT) EditBy(tablename string, i interface{}, key string, value interface{}) bool {
	res := d.Collection(tablename).Find(db.Cond{key: value})
	err := res.Update(i)
	if err != nil {
		utility.Error(err, "EditBy")
		return false
	}
	if err := res.Close(); err != nil {
		utility.Error(err, "EditBy")
	}
	return true
}

func (d *DatastoreT) DeleteBy(table string, cond db.Cond) bool {
	res := d.Collection(table).Find(cond)
	err := res.Delete()
	if err != nil {
		return false
	}
	if err := res.Close(); err != nil {
		utility.Error(err, "DeleteBy")
	}
	return true
}

func (d *DatastoreT) Count(tablename string) (int, bool) {
	res := d.Collection(tablename).Find()
	count, err := res.Count()
	if err != nil {
		utility.Error(err, "Count")
		return 0, false
	}
	if err := res.Close(); err != nil {
		utility.Error(err, "Count")
	}
	return int(count), true
}
func (d *DatastoreT) CountByCond(table string, cond db.Cond) (int, bool) {
	res := d.Collection(table).Find(cond)
	count, err := res.Count()

	if err != nil {
		utility.Error(err, "CountByCond")
		return 0, false
	}
	if err := res.Close(); err != nil {
		utility.Error(err, "CountByCond")
	}
	return int(count), true
}

func (d *DatastoreT) GetCursor(table string, order string, paginate int, cond db.Cond) db.Result {
	return d.Collection(table).Find(cond).OrderBy(order).Paginate(uint(paginate)).Cursor(order)
}

func (d *DatastoreT) GetLarge(table string, order string, cond db.Cond) db.Result {
	return d.Collection(table).Find(cond).OrderBy(order)
}

//func (c *Picture) IsEmpty() bool { return reflect.DeepEqual(Picture{}, c) }
