package services

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/hoannguyen02/self-go/models"
	"github.com/hoannguyen02/self-go/redis"
)

func usersKey(userId string) string {
	return "users#" + userId
}

func usernamesKey() string {
	return "usernames"
}

func usernamesUniqueKey() string {
	return "username#unique"
}

func genId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func CreateUser(user *models.User) error {
	// Check if username exists
	exists, err := redis.Client.SIsMember(usernamesUniqueKey(), user.Username).Result()
	if err != nil {
		return err
	}

	if exists {
		return errors.New("username is taken")
	}

	// Gen id
	id, err := genId()
	if err != nil {
		return err
	}
	user.Id = id
	// Hash password
	// Using interface such as Structs in order to write a common function to save data into redis
	// input struct, key
	// output error
	// Save user in hash map with key from user id
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	var userMap map[string]interface{}
	_ = json.Unmarshal(userJson, &userMap)
	err = redis.Client.HMSet(usersKey(user.Id), userMap).Err()
	if err != nil {
		return err
	}
	// Save username with user id to hash map with usernames key
	usernames := models.UserNames {
		Username: id,
	}
	var usernamesMap map[string]interface{}
	usernamesJson, err := json.Marshal(usernames)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(usernamesJson, &usernamesMap)
	err = redis.Client.HMSet(usernamesKey(), usernamesMap).Err()
	if err != nil {
		return err
	}

	err = redis.Client.SAdd(usernamesUniqueKey(), user.Username).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (models.User, error) {
	userId, err := redis.Client.HGet(usersKey(username), "username").Result()
	if err != nil {
		return models.User{}, err
	}

    return GetUserByUserId(userId)
}


func GetUserByUserId(userId string) (models.User, error) {
	userMap, err := redis.Client.HGetAll(usersKey(userId)).Result()
	if err != nil {
		return models.User{}, err
	}

    userJson, err := json.Marshal(userMap)
    if err != nil {
        return models.User{}, err
    }

    user := models.User{}
    if err := json.Unmarshal(userJson, &user); err != nil {
       return models.User{}, err
    }

	return user, err
}
