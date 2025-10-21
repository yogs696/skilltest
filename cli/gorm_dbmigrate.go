package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/yogs696/skilltest/app"
	"github.com/yogs696/skilltest/cli/dbgorm"
)

// Main variable argument
var dbMigrate bool

// Option variable argument
var tableName string

var seedRefresh bool

var dbMigrateCommands = cli{
	argVar:   &dbMigrate,
	argName:  "db-migrate",
	argUsage: "-db-migrate To start migrations. If without sub-argument will migrate all table",
	run:      dbMigrateRun,
	stringOptions: []optionString{
		{
			optionVar:          &tableName,
			optionName:         "table",
			optionUsage:        "-table=<table name> Just migrate specific table instead migrate all",
			optiondefaultValue: "",
		},
	},
}

func dbMigrateRun() {
	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Doing DB migrations...")
	time.Sleep(time.Second)

	// Open DB connection
	spinnerLiveText.UpdateText("Opening DB connection...")
	app.Up(&app.AppArgs{}, 1)
	defer func() {
		// Closing DB connection
		spinnerLiveText.UpdateText("Closing DB connection...")
		app.Down(1)

		spinnerLiveText.Success("DB successfully migrated")
		fmt.Println()
	}()

	// Check DB connection
	if app.DBA == nil {
		spinnerLiveText.Fail("Failed to open DB connection")
		return
	}

	if tableName != "" && tableName != " " {
		// Start migration based on given table name if any
		spinnerLiveText.UpdateText(fmt.Sprintf("Just migrate %s table...", tableName))
		t := strings.Split(tableName, ",")
		if err := startMigrator(spinnerLiveText, t...); err != nil {
			spinnerLiveText.Fail(err.Error())
		}

		return
	}

	// Start migration all
	if err := startMigrator(spinnerLiveText); err != nil {
		spinnerLiveText.Fail(err.Error())
	}
}

// Helper function to execute migrator up
func startMigrator(st *pterm.SpinnerPrinter, t ...string) error {
	st.UpdateText("Start migration...")

	return dbgorm.RunDBMigrate(app.DBA.DB, seedRefresh, t...)
}
