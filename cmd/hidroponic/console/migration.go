package console

import (
	"fmt"
	"hidroponic/internal/platform/configuration"
	"hidroponic/internal/platform/database"
	"hidroponic/internal/platform/database/migration"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func (cs Console) initCreateMigrationCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()
	cmd.Use = "migration:create"
	cmd.Short = "create migration files"
	cmd.Args = cobra.ExactArgs(1)
	cmd.Run = func(_ *cobra.Command, args []string) {
		name := strings.Replace(args[0], " ", "_", -1)
		migrator := initMigrator()
		if migrator != nil {
			if err := migrator.CreateMirationFiles(name); err != nil {
				fmt.Println("error creating migration files with message: ", err.Error())
			}
		}
	}

	cs.ConsoleService.RegisterCommand(cmd)
}

func (cs Console) initRunMigrateCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()
	cmd.Use = "migration:migrate"
	cmd.Short = "run migration"
	cmd.Run = func(_ *cobra.Command, args []string) {
		migrator := initMigrator()
		if migrator != nil {
			if err := migrator.Migrate(); err != nil {
				fmt.Println("error running migration with message: ", err.Error())
			}
		}
	}

	cs.ConsoleService.RegisterCommand(cmd)
}

func (cs Console) initRunForceMigrateCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()
	cmd.Use = "migration:force-migrate"
	cmd.Short = "run force migration"
	cmd.Args = cobra.ExactArgs(1)
	cmd.Run = func(_ *cobra.Command, args []string) {
		version, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid version value, skip process")
			return
		}

		migrator := initMigrator()
		if migrator != nil {
			if err := migrator.ForceMigrate(version); err != nil {
				fmt.Println("error running migration with message: ", err.Error())
			}
		}
	}

	cs.ConsoleService.RegisterCommand(cmd)
}

func (cs Console) initRunRollbackCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()
	cmd.Use = "migration:rollback"
	cmd.Short = "run rollback"
	cmd.Run = func(_ *cobra.Command, args []string) {
		migrator := initMigrator()
		if migrator != nil {
			if err := migrator.Rollback(); err != nil {
				fmt.Println("error running rollback with message: ", err.Error())
			}
		}
	}

	cs.ConsoleService.RegisterCommand(cmd)
}

func initMigrator() *migration.MigrationService {
	config := configuration.New(".", "config", "yaml")
	err := config.Setup()
	if err != nil {
		fmt.Println("cannot init configuration with message: ", err.Error())
		return nil
	}

	db := database.New(config)
	err = db.Setup()
	if err != nil {
		fmt.Println("cannot init database with message: ", err.Error())
		return nil
	}

	migrator := migration.New(db.DB.DB)
	err = migrator.Setup()
	if err != nil {
		fmt.Println("cannot init migration service with message: ", err.Error())
		return nil
	}

	return migrator
}
