package modelsbz

// Model 哦
type Model interface {
	GetInstance() interface{}
}

//GetFirst 获取第一条
func GetFirst(model Model) (result interface{}, err error) {
	i := model.GetInstance()
	err = DB.Debug().Where(i).First(i).Error
	return i, err
}