package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

//UserList as a database
type UserList []User

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d Not find in UserList", id)
}
func (u *UserList) addUser(user User) {

	*u = append(*u, user)
}

type UserListProxy struct {
	//source data
	SomeDatabase UserList
	//cache data
	StackCache UserList
	//cache size
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returing user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}
	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, fmt.Errorf("Not find user %d in database", id)
	}
	u.addUserToStack(user)
	fmt.Println("Returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.SomeDatabase[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}
