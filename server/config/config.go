package config

import (
	"jobs/internal/job"
	"jobs/internal/platform/idgenerator"
	"jobs/internal/platform/validate"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Cfg struct {
	Server Server
}
type Server struct {
	Port            string `validate:"required"`
	WriteTimeoutSec int64  `validate:"required"`
	ReadTimeoutSec  int64  `validate:"required"`
	ShutdownTime    int64  `validate:"required"`
	QueueCapacity   int64  `validate:"required"`
}

// Builder defines the parametric information of a server instance.
type Builder struct {
	Cfg           Cfg
	Environment   string
	Validate      *validate.Validator
	Identificator *idgenerator.Counter
	Jobs          *job.Config
}

// InitCFG initializes the server builder with properties retrieved from Viper.
func Init(v *viper.Viper) *Builder {
	var b Builder
	var err error
	if err = v.Unmarshal(&b.Cfg); err != nil {
		log.Panic(errors.Wrap(err, "Config.Init() could not unmarshal config "))
	}
	b.Validate = validate.Init()
	if err = b.Validate.Struct(b.Cfg); err != nil {
		log.Panic("Config.Init() could not validate config: ", err)
	}
	b.Environment = v.GetString("ENV")
	b.Identificator = idgenerator.Init()
	b.Jobs = job.Init()

	return &b
}
