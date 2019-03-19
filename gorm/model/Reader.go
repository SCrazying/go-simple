package model

type Reader struct {
	Id          int
	Name        string
	Password    string
	Sex         string
	BorrowedNum int
}

func (*Reader) TableName() string {
	return "reader"
}

//使用用户名和密码查询用户
func FindReader(id int, password string) *Reader {
	reader := &Reader{}
	flag := Db.Where("id=? AND password=?", id, password).Find(&reader).RecordNotFound()
	if flag {
		return nil
	} else {
		return reader
	}
}

//获取所有用户
func FindReaders() []*Reader {
	readers := make([]*Reader, 0)
	Db.Find(&readers)
	return readers
}

//获取单独用户信息
func FindReaderInfo(id int) *Reader {
	reader := &Reader{}
	flag := Db.Where("id=?", id).Find(&reader).RecordNotFound()
	if flag {
		return nil
	} else {
		return reader
	}
}

//更新用户名
func UpdateReaderName(id int, name string) *Reader {
	reader := &Reader{}
	flag := Db.Where("id=?", id).Find(&reader).RecordNotFound()
	if flag {
		return nil
	}

	Db.Model(reader).Update("name", name)

	return reader
}

//更新用户密码
func UpdateReaderPassword(id int, password string) *Reader {
	reader := &Reader{}
	flag := Db.Where("id=?", id).Find(&reader).RecordNotFound()
	if flag {
		return nil
	}

	Db.Model(reader).Update("password", password)

	return reader

}

//插入用户
func InsertReaderInfo(name, password, sex string) *Reader {
	reader := Reader{Name: name, Password: password, Sex: sex}

	flag := Db.NewRecord(reader)
	if flag {
		Db.Create(&reader)
		flag = Db.NewRecord(reader)
		if !flag {
			return &reader
		} else {
			return nil
		}
	} else {
		return nil
	}
}
