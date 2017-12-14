package service

import (
	"Agenda_hw2/entity"
	"Agenda_hw2/loghelper"
	"log"
)

var curuserinfoPath = "../data/curuser.txt"
var errLog *log.Logger

type User entity.User
type Meeting entity.Meeting

func init() {
	errLog = loghelper.Error
}
func UserLogout() bool {
	if err := entity.Logout(); err != nil {
		return false
	} else {
		return true
	}
}
func GetCurUser() (entity.User, bool) {
	if cu, err := entity.GetCurUser(); err != nil {
		return cu, false
	} else {
		return cu, true
	}
}
func UserLogin(username string, password string) bool {
	user := entity.QueryUser(func(u *entity.User) bool {
		if u.Name == username && u.Password == password {
			return true
		}
		return false
	})
	if len(user) == 0 {
		errLog.Println("Login: User not Exist")
		return false
	}
	entity.SetCurUser(&user[0])
	if err := entity.Sync(); err != nil {
		errLog.Println("Login: error occurred when set curuser")
		return false
	}
	return true
}

func UserRegister(username string, password string, email string, phone string) (bool, error) {
	user := entity.QueryUser(func(u *entity.User) bool {
		return u.Name == username
	})
	if len(user) == 1 {
		errLog.Println("User Register: Already exist username")
		return false, nil
	}
	entity.CreateUser(&entity.User{username, password, email, phone})
	if err := entity.Sync(); err != nil {
		return true, err
	}
	return true, nil
}

func DeleteUser(username string) bool {
	entity.DeleteUser(func(u *entity.User) bool {
		return u.Name == username
	})
	entity.UpdateMeeting(
		func(m *entity.Meeting) bool {
			return m.IsParticipator(username)
		},
		func(m *entity.Meeting) {
			m.DeleteParticipator(username)
		})
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == username || len(m.GetParticipator()) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return UserLogout()
}

func ListAllUser() []entity.User {
	return entity.QueryUser(func(u *entity.User) bool {
		return true
	})
}
