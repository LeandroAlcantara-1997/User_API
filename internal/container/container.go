package container

import (
	"context"

	"github.com/facily-tech/go-core/env"
	"github.com/facily-tech/go-core/log"
	"github.com/facily-tech/go-core/telemetry"
	"github.com/facily-tech/go-core/types"
	"github.com/facily-tech/go-scaffold/internal/config"
	postgresConfig "github.com/facily-tech/go-scaffold/pkg/core/postgres"
	"github.com/facily-tech/go-scaffold/pkg/domains/quote"
	"github.com/facily-tech/go-scaffold/pkg/domains/user"
	userRepo "github.com/facily-tech/go-scaffold/pkg/domains/user/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Components are a like service, but it doesn't include business case
// Or domains, but likely used by multiple domains

type components struct {
	Log            log.Logger
	Tracer         telemetry.Tracer
	PostgresClient *gorm.DB
	// Include your new components bellow
}

type envs struct {
	postgres postgresConfig.Config
}

// Services hold the business case, and make the bridge between
// Controllers and Domains
type Services struct {
	Quote quote.ServiceI
	User  user.ServiceI
	// Include your new services bellow
}

type Dependency struct {
	Components components
	Services   Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	envs, err := loadEnvs(ctx)
	if err != nil {
		return nil, nil, err
	}
	cmp, err := setupComponents(ctx, envs)
	if err != nil {
		return nil, nil, err
	}

	quoteService, err := quote.NewService(
		quote.NewRepository(cmp.Log),
		cmp.Log,
	)

	userService, err := user.NewService(
		userRepo.NewRepository(
			cmp.PostgresClient,
		),
		cmp.Log,
	)

	if err != nil {
		return nil, nil, err
	}

	srv := Services{
		User:  userService,
		Quote: quoteService,
		// include services initialized above here
	}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return ctx, &dep, err
}

func loadEnvs(ctx context.Context) (envs, error) {
	postDB := postgresConfig.Config{}

	err := env.LoadEnv(ctx, &postDB, postgresConfig.ConfigPrefix)
	if err != nil {
		return envs{}, err
	}
	return envs{
		postgres: postDB,
	}, nil
}

func setupComponents(ctx context.Context, envs envs) (*components, error) {
	version, ok := ctx.Value(types.ContextKey(types.Version)).(*config.Version)
	if !ok {
		return nil, config.ErrVersionTypeAssertion
	}

	telemetryConfig := telemetry.DataDogConfig{
		Version: version.GitCommitHash,
	}

	err := env.LoadEnv(ctx, &telemetryConfig, telemetry.DataDogConfigPrefix)
	if err != nil {
		return nil, err
	}

	tracer, err := telemetry.NewDataDog(telemetryConfig)

	if err != nil {
		return nil, err
	}

	l, err := log.NewLoggerZap(log.ZapConfig{
		Version:           version.GitCommitHash,
		DisableStackTrace: true,
		Tracer:            tracer,
	})

	if err != nil {
		return nil, err
	}

	// init postgres db
	db, err := gorm.Open(postgres.Open(envs.postgres.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &components{
		Log:            l,
		Tracer:         tracer,
		PostgresClient: db,
		// include components initialized bellow here
	}, nil
}
