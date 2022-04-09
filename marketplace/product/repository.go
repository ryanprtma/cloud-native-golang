package product

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Repository interface {
	Save(product Products) (Products, error)
	Update(id string, quantity int) (Products, error)
	FindByIDRow(id string) (Products, error)
	FindByID(ID string) (Products, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Products) (Products, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(id string, quantity int) (Products, error) {
	db := r.db.Debug()
	product := Products{}

	tx := db.Begin()

	err := tx.Model(&product).Where("id = ?", id).Update("stock", quantity).Error
	if err != nil {
		tx.Rollback()
		return product, err
	}

	tx.Commit()
	return product, nil
}

func (r *repository) FindByIDRow(id string) (Products, error) {
	var product Products

	err := r.db.First(&product, id).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByID(ID string) (Products, error) {
	var product Products

	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	fmt.Println("tes redis")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "",
		DB:       1,
	})

	fmt.Println(client.Ping().String())

	value, err := client.Get("one").Result()
	if err == redis.Nil {
		mars, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err)
			return product, err
		}
		err = client.Set("one", mars, time.Second*15).Err()
		if err != nil {
			fmt.Println(err)
			return product, err
		}

		fmt.Println("Save")
		return product, nil

	} else if err != nil {
		fmt.Printf("error calling redis: %v\n", err)
		return product, err
	} else {
		err = json.Unmarshal([]byte(value), &product)
		if err != nil {
			return product, err
		}

		fmt.Println("Done")
		return product, nil
	}
}
