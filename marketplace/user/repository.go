package user

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID string) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "",
		DB:       1,
	})

	fmt.Println(client.Ping().String())

	value, err := client.Get("user_email").Result()
	if err == redis.Nil {
		mars, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return user, err
		}
		err = client.Set("user_email", mars, time.Second*15).Err()
		if err != nil {
			fmt.Println(err)
			return user, err
		}

		fmt.Println("Save")
		return user, nil

	} else if err != nil {
		fmt.Printf("error calling redis: %v\n", err)
		return user, err
	} else {
		err = json.Unmarshal([]byte(value), &user)
		if err != nil {
			return user, err
		}

		fmt.Println("Done")
		return user, nil
	}

	// return user, nil
}

func (r *repository) FindByID(ID string) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "",
		DB:       1,
	})

	fmt.Println(client.Ping().String())

	value, err := client.Get("one").Result()
	if err == redis.Nil {
		mars, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return user, err
		}
		err = client.Set("one", mars, time.Second*15).Err()
		if err != nil {
			fmt.Println(err)
			return user, err
		}

		fmt.Println("Save")
		return user, nil

	} else if err != nil {
		fmt.Printf("error calling redis: %v\n", err)
		return user, err
	} else {
		err = json.Unmarshal([]byte(value), &user)
		if err != nil {
			return user, err
		}

		fmt.Println("Done")
		return user, nil
	}

	// return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
