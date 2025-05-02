package trip

import (
	"errors"
)

var session *UserSession

type Service struct {
	tripDAO *Dao
}

func (service *Service) GetTripByUser(user *User) ([]Trip, error) {
	var trips []Trip
	friends := user.Friends()
	loggedUser, err := session.GetLoggedUser()
	if err != nil {
		return trips, err
	}
	var isFriend bool
	if loggedUser != nil {
		for _, friend := range friends {
			if loggedUser.id == friend.id {
				isFriend = true
				break
			}
		}
		if isFriend {
			return service.tripDAO.FindTripsByUser(user)
		}
		return trips, err
	} else {
		return trips, errors.New("user not logged in")
	}
}

type UserSession struct{}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, errors.New("UserSession.GetLoggedUser() should not be called in an unit test")
}

type User struct {
	id      string
	friends []User
	trips   []Trip
}

func (user *User) Friends() []User {
	if user == nil {
		return nil
	}
	return user.friends
}

func (user *User) AddFriend(friend *User) {
	if friend != nil {
		user.friends = append(user.friends, *friend)
	}
}

func (user *User) Trips() []Trip {
	if user == nil {
		return nil
	}
	return user.trips
}

func (user *User) AddTrip(trip *Trip) {
	if trip != nil {
		user.trips = append(user.trips, *trip)
	}
}

type Trip struct{}

type Dao struct{}

func (dao *Dao) FindTripsByUser(user *User) ([]Trip, error) {
	return nil, errors.New("TripDAO should not be invoked on an unit test.")
}
