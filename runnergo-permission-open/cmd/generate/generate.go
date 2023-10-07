package main

import (
	"fmt"

	"permission-open/internal/pkg/conf"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func MustInitConf() {
	viper.SetConfigFile("../../configs/dev.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&conf.Conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	fmt.Println("config initialized")
}

func main() {
	MustInitConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.Conf.MySQL.Username,
		conf.Conf.MySQL.Password, conf.Conf.MySQL.Host, conf.Conf.MySQL.Port, conf.Conf.MySQL.DBName, conf.Conf.MySQL.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/pkg/dal/query",
	})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("team"),
		g.GenerateModel("user_team"),
		g.GenerateModel("setting"),
		g.GenerateModel("public_function"),
		g.GenerateModel("company"),
		g.GenerateModel("user_company"),
		g.GenerateModel("role"),
		g.GenerateModel("user_role"),
		g.GenerateModel("permission"),
		g.GenerateModel("role_permission"),
		g.GenerateModel("role_type_permission"),
		g.GenerateModel("user_team_collection"),
		g.GenerateModel("third_notice"),
		g.GenerateModel("third_notice_channel"),
		g.GenerateModel("third_notice_group"),
		g.GenerateModel("third_notice_group_relate"),
	)

	g.Execute()
}
