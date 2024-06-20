package data

//func NewUserData() *UserData {
//	return &UserData{}
//}
//
//type UserData struct {
//}
//
//func (s UserData) List() ([]*po.User, error) {
//	list := make([]*po.User, 0)
//	err := GetMerchantDB().Find(&list).Error
//	if err != nil {
//		logrus.Errorf("UserData List Err: %s", err.Error())
//		return nil, err
//	}
//	return list, nil
//}
//
//func (s UserData) InfoByName(name string) (*po.User, error) {
//	info := &po.User{}
//	err := GetMerchantDB().Find(info, "username = ?", name).Error
//	if err != nil {
//		logrus.Errorf("UserData info Err: %s", err.Error())
//		return nil, err
//	}
//	return info, nil
//}
//
//func (s UserData) Info(id int64) (*po.User, error) {
//	info := &po.User{}
//	err := GetMerchantDB().Find(info, "id = ?", id).Error
//	if err != nil {
//		logrus.Errorf("UserData info Err: %s", err.Error())
//		return nil, err
//	}
//	return info, nil
//}
//
//func (s UserData) Edit(eUser *po.User) error {
//	err := GetMerchantDB().Model(&po.User{}).Where("id = ?", eUser.Id).Updates(eUser).Error
//	if err != nil {
//		logrus.Errorf("UserData Edit Err: %s", err.Error())
//		return err
//	}
//	return nil
//}
//
//func (s UserData) Delete(eUser *po.User) error {
//	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
//	if err != nil {
//		logrus.Errorf("UserData Delete Err: %s", err.Error())
//		return err
//	}
//	return nil
//}
