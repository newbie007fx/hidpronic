package console

import "hidroponic/internal/platform/console"

type Console struct {
	ConsoleService *console.ConsoleService
}

func InitApp(consoleService *console.ConsoleService) {
	cs := &Console{
		ConsoleService: consoleService,
	}

	cs.initServeCommand()
	cs.initCreateMigrationCommand()
	cs.initRunMigrateCommand()
	cs.initRunForceMigrateCommand()
	cs.initRunRollbackCommand()
}
