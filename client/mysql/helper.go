package mysql

import (
	"caixin.app/caixos/tokit/constant"
	"errors"
)

func First(query interface{}, args []interface{}, bean interface{}) error {
	has, err := GetDB().SQL(query, args...).Get(bean)
	if err != nil {
		return err
	}
	if !has {
		return errors.New(constant.ErrNotExist)
	}
	return nil
}

func Fetch(query interface{}, args []interface{}, bean interface{}) error {
	return GetDB().SQL(query, args...).Find(bean)
}

func Insert(bean interface{}) bool {
	affected, _ := GetDB().Insert(bean)
	if affected == 1 {
		return true
	}
	return false
}

func Update(bean interface{}, cond interface{}) bool {
	affected, _ := GetDB().Update(bean, cond)
	if affected == 1 {
		return true
	}
	return false
}
