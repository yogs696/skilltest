package app

import (
	"github.com/yogs696/skilltest/internal/repo"
	"github.com/yogs696/skilltest/pkg/kemu"

	appv1auth "github.com/yogs696/skilltest/usecase/v1/auth"
	appv1userhttp "github.com/yogs696/skilltest/usecase/v1/auth/http"

	appv1schedule "github.com/yogs696/skilltest/usecase/v1/schedule"
	appv1schedulehttp "github.com/yogs696/skilltest/usecase/v1/schedule/http"
)

// Helper table user function will return entity repository that using gorm
func getRepoAuthGorm() *repo.UserRepoDB {
	return repo.NewUserRepoDB(DBA.DB, DBA.SQL)
}

// Helper table Schedule function will return entity repository that using gorm
func getRepoScheduleGorm() *repo.ScheduleRepoDB {
	return repo.NewScheduleRepoDB(DBA.DB, DBA.SQL)
}

// DoEnV1Register register domain entity handler version 1 into the app
func doEntV1Register(args *AppArgs) {
	kemu := kemu.New()

	if HardMaintenance == "false" {
		printOutUp("Registering domain entity handler...")

		// user
		appv1usersvc := appv1auth.NewService(
			kemu,
			getRepoAuthGorm(),
			printOutUp,
		)
		appv1userhttp.RegisterRoute(API.RouteGroup["v1"], *appv1usersvc, kemu)

		// Schedule
		appv1ScheduleSvc := appv1schedule.NewService(
			kemu,
			getRepoScheduleGorm(),
			printOutUp,
		)
		appv1schedulehttp.RegisterRoute(API.RouteGroupWithMiddleware["v1"], *appv1ScheduleSvc, kemu)

	}
}
