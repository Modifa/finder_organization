package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	models "github.com/Modifa/finder_organization.git/models"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func GetOrganizationProfile(Key string) (bool, *models.OrganizationResponse) {
	var cp *models.OrganizationResponse

	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0, // use default DB
	})

	val, err := rdb.Get(ctx, "ORGANIZATION:"+strings.Title(Key)).Result()

	defer rdb.Close()
	if err != nil {
		fmt.Println(err)
		return false, cp
	}

	byt := []byte(val)

	if err := json.Unmarshal(byt, &cp); err != nil {
		panic(err)
	}

	return true, cp
}

////Set Redis Developer Profile
func SaveOrganizationprofile(Organization models.OrganizationResponse) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0,
	})
	b, err := json.Marshal(Organization)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	/**/
	err = rdb.Set(ctx, "ORGANIZATION:"+strings.Title(Organization.EmailAddress), b, 0).Err()
	return err
}

//
func GetOrganizationProfileRedis(Key string) (bool, *models.OrganizationResponse) {
	var op *models.OrganizationResponse

	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0, // use default DB
	})

	//err := rdb.Set(ctx, "TAXIMONEY:TAXIPROFILE:"+taxino, taxino, 0).Err()
	//	Newkey := strings.Title(User.EmailAddress)

	val, err := rdb.Get(ctx, "ORGANIZATION:"+strings.Title(Key)).Result()

	defer rdb.Close()
	if err != nil {
		//panic(err)
		return false, op
	}

	byt := []byte(val)

	if err := json.Unmarshal(byt, &op); err != nil {
		panic(err)
	}

	//fmt.Println("key", val)

	return true, op

}
