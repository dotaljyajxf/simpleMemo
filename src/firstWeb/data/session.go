/**********************************************************************************************************************
 *
 * Copyright (c) 2010 babeltime.com, Inc. All Rights Reserved
 * $
 *
 **********************************************************************************************************************/

/**
 * @file $
 * @author $(liujianyong@babeltime.com)
 * @date $
 * @version $
 * @brief
 *
 **/
package data

import "errors"

var Session *SessionStruct

type SessionStruct struct {
	users map[string]*UserData
}

type UserData struct {
	data map[string]string
}

func newUserData() *UserData {
	return &UserData{
		make(map[string]string),
	}
}

func NewSession() *SessionStruct {
	return &SessionStruct{
		make(map[string]*UserData),
	}
}

func (s *SessionStruct) addUser(token string) {
	if _, ok := s.users[token]; ok {
		return
	}
	s.users[token] = newUserData()
}

func (s *SessionStruct) GetSession(token string) *UserData {
	s.addUser(token)
	return s.users[token]
}

func (u *UserData) Store(key string, val string) error {
	if _, ok := u.data[key]; ok {
		return errors.New("KeyExist")
	}
	s.data[key] = val
	return nil
}

func (u *UserData) Get(key string) (string, error) {
	if _, ok := u.data[key]; !ok {
		return "", errors.New("KeyNotExist")
	}
	return u.data[key], nil
}
