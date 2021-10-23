package cmd

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	leftUsername  string
	rightUsername string
	leftPassword  string
	rightPassword string
	leftHost      string
	rightHost     string
	leftPort      string
	rightPort     string
	verbose       bool
)

func hydrate() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     leftHost + ":" + leftPort,
		Password: leftPassword,
		DB:       0,
	})
	if rdb != nil {
		defer rdb.Close()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	val, _ := rdb.Do(ctx, "PING").Result()
	fmt.Printf("%v\n", val)

	db, err := sql.Open("mysql", rightUsername+":"+rightPassword+"@tcp("+rightHost+":"+rightPort+")/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(`SHOW DATABASES;`)
	if err != nil {
		panic(err)
	}
	cols, _ := rows.Columns()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s:%s\n", cols[0], name)
		}
	}
}

var hydrateCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Long:  "run",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if !f.Changed && viper.IsSet(f.Name) {
				val := viper.Get(f.Name)
				cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			}
		})

	},
	Run: func(cmd *cobra.Command, args []string) {
		hydrate()
	},
}

func getHydrateCmd() *cobra.Command {
	hydrateCmd.Flags().StringVarP(&leftUsername, "lusername", "", "", "Left service username")
	hydrateCmd.Flags().StringVarP(&rightUsername, "rusername", "", "", "Right service username")
	hydrateCmd.Flags().StringVarP(&leftPassword, "lpassword", "", "", "Left service username")
	hydrateCmd.Flags().StringVarP(&rightPassword, "rpassword", "", "", "Right service username")
	hydrateCmd.Flags().StringVarP(&leftHost, "lhost", "", "127.0.0.1", "Left host IP")
	hydrateCmd.Flags().StringVarP(&rightHost, "rhost", "", "127.0.0.1", "Right host IP")
	hydrateCmd.Flags().StringVarP(&leftPort, "lport", "", "6379", "Left host port")
	hydrateCmd.Flags().StringVarP(&rightPort, "rport", "", "3306", "Right host port")
	hydrateCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	hydrateCmd.MarkFlagRequired("rusername")

	return hydrateCmd
}
