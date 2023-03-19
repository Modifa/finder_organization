package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/Modifa/finder_organization.git/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
}

//Get Developer Profile
func (db *DB) GetOrganizationProfile(functionnamewithschema string, m interface{}) ([]models.OrganizationResponse, error) {
	Org := []models.OrganizationResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &Org, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return Org, err
}

//Add
func (db *DB) SAVEONDB(functionnamewithschema string, m interface{}) (int64, error) {
	var returnInt int64
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, errCon := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	if errCon != nil {
		return 0, errCon
	}
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &returnInt, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return returnInt, err
}

//Add
func (db *DB) SAVEONDBNPRETURN(functionnamewithschema string, m interface{}) error {
	var returnInt []int64
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &returnInt, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return err
}

//DocumentReponse

//Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}
