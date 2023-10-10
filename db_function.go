package orm

// Insert 插入一条数据
func Insert(data interface{}) (int64, error) {
	result := Orm.Create(data)
	return result.RowsAffected, result.Error
}

// Select 根据主键id查询数据
// data 数据结构
// var user = User{ID: 10}  db.First(user)
func Select(data interface{}) (int64, error) {
	result := Orm.First(data)
	return result.RowsAffected, result.Error
}

// SelectByID 根据主键id查询数据
// data 数据结构
// value 主键ID值（int or string)都行
// db.First(user, 10)
func SelectByID(data interface{}, value interface{}) (int64, error) {
	result := Orm.First(data, value)
	return result.RowsAffected, result.Error
}

// SelectWhere 根据条件查询
// db.Where("name = ?", "jinzhu").First(user)
func SelectWhere(data interface{}, key string, value string) (int64, error) {
	result := Orm.Where(key+" = ?", value).First(data)
	return result.RowsAffected, result.Error
}

// SelectAll 查询全部数据
// SELECT * FROM users;
func SelectAll(data interface{}) (int64, error) {
	result := Orm.Find(data)
	return result.RowsAffected, result.Error
}

func SelectAllWhere(data interface{}, key string, value interface{}) (int64, error) {
	result := Orm.Where(key+" = ?", value).Find(data)
	return result.RowsAffected, result.Error
}

func SelectAllLikeWhere(data interface{}, key string, value string) (int64, error) {
	result := Orm.Where(key+" LIKE ?", "%"+value+"%").Find(data)
	return result.RowsAffected, result.Error
}

// Update 根据 `struct` 更新属性，只会更新非零值的字段
// db.Model(user).Updates(User{Name: "hello", Age: 18, Active: false})
func Update(data interface{}) (int64, error) {
	result := Orm.Model(data).Updates(data)
	return result.RowsAffected, result.Error
}

// UpdateMap 根据 `map` 更新属性
// db.Model(user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
func UpdateMap(data interface{}, values map[string]interface{}) (int64, error) {
	result := Orm.Model(data).Updates(data)
	return result.RowsAffected, result.Error
}

// Delete 删除一条数据
func Delete(data interface{}) (int64, error) {
	result := Orm.Delete(data)
	return result.RowsAffected, result.Error
}
