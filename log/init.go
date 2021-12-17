package log

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Logger is an alias of logrus.Logger
type Logger = logrus.Logger

// Fields is an alias of logrus.Fields
type Fields = logrus.Fields

var (
	// flagset
	cli = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// viper
	v = viper.New()

	// DefaultEntry : default logger object
	qEntry = logrus.NewEntry(StandardLogger())

	qLoggerConfig = &loggerConfig{
		Path: make([]string, 0),
	}
)

const (
	keyConfigPath = "logger.config.path"
	keyConfigName = "logger.config.name"
	keyConfigType = "logger.config.type"
	keyConfigFile = "logger.config.file"

	keyReportCaller         = "logger.reportcaller"
	keyDefaultLevel         = "logger.level"
	keyDefaultFormatterName = "logger.formatter.name"
	keyDefaultFormatterOpts = "logger.formatter.opts"
)

func initFlags() error {
	cli.Bool(keyReportCaller, true, "logger.reportcaller")
	cli.String(keyDefaultLevel, "debug", "logger.level")
	cli.String(keyDefaultFormatterName, "classic", "logger.formatter.name")

	cli.StringSliceVar(&qLoggerConfig.Path, keyConfigPath, defaultConfigPath, "logger.config.path")
	cli.StringVar(&qLoggerConfig.Name, keyConfigName, defaultConfigName, "logger.config.name")
	cli.StringVar(&qLoggerConfig.Type, keyConfigType, defaultConfigType, "logger.config.type")

	cli.StringVar(&qLoggerConfig.File, keyConfigFile, "", "logger.config.file")

	return nil
}

func initViper() error {
	// read from flags
	_ = cli.Parse(filterLoggerFlags(os.Args[1:], true))
	_ = v.BindPFlags(cli)

	// read from env
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// set default

	// read from config file
	if len(qLoggerConfig.File) > 0 {
		v.SetConfigFile(qLoggerConfig.File)
	} else {
		for _, p := range qLoggerConfig.Path {
			v.AddConfigPath(p)
		}
		v.SetConfigName(qLoggerConfig.Name)
	}
	v.SetConfigType(qLoggerConfig.Type)

	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// no config file, enable stdout by default
			v.SetDefault("logger.stdout.enabled", true)
		default:
			return err
		}
	} else {
		// watch configs changes
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("[log] config changed: ", e.Name)
			resetLogger()
		})
	}

	//v.Debug()

	return nil
}

func resetLogger() {
	if err := configLogger(); err != nil {
		fmt.Printf("[log] reload config fail:%s, changes may not take effect!", err)
	}
}

func getDefaultFormatter() (logrus.Formatter, error) {
	return newFormatter(v.GetString(keyDefaultFormatterName), keyDefaultFormatterOpts)
}

func initActivateHooks() error {
	var err error
	var hook logrus.Hook
	var activateHooks = make(logrus.LevelHooks)

	for name := range gRegisteredHooks {
		n := strings.Join([]string{"logger", name, "enabled"}, ".")
		//fmt.Println("initHooks", n, v.GetBool(n))
		if v.GetBool(n) {
			if hook, err = newHook(name); err != nil {
				fmt.Printf("[log] init hook(%s) error:%s\n", name, err)
				continue
			}
			setupHooks[name] = hook
			activateHooks.Add(hook)
		}
	}

	if len(activateHooks) == 0 {
		return errors.New("no activate log hook")
	}

	return nil
}

func getActivateHooks() logrus.LevelHooks {
	var hooks = make(logrus.LevelHooks)

	for name := range setupHooks {
		hooks.Add(setupHooks[name])
	}

	if len(hooks) == 0 {
		return nil
	}
	return hooks
}

func configLogger() error {
	var err error

	qEntry.Logger.SetReportCaller(v.GetBool(keyReportCaller))
	level, err := logrus.ParseLevel(v.GetString(keyDefaultLevel))

	if err != nil {
		return fmt.Errorf("get default log level error: %s", err)
	}
	qEntry.Logger.SetLevel(level)

	formatter, err := getDefaultFormatter()
	if err != nil {
		return fmt.Errorf("get default formatters error: %s", err)
	}
	qEntry.Logger.SetFormatter(formatter)

	// SetLevel and SetFormatter must be called before getActivateHooks.
	err = initActivateHooks()

	if err != nil {
		fmt.Printf("[log] get hooks error: %s\n", err)
		qEntry.Logger.ReplaceHooks(nil)
		qEntry.Logger.SetOutput(os.Stderr)
		return nil
	}

	qEntry.Logger.SetOutput(ioutil.Discard)
	qEntry.Logger.ReplaceHooks(getActivateHooks())
	return nil
}

func init() {
	var err error

	if err = initFlags(); err != nil {
		panic(fmt.Sprint("[log] init flags error:", err))
	}

	if err = initViper(); err != nil {
		panic(fmt.Sprint("[log] init viper error:", err))
	}

	if err = configLogger(); err != nil {
		panic(fmt.Sprint("[log] configLogger fail:", err))
	}
}

func SetConfigFile(file string) (err error) {
	qLoggerConfig.File = file
	err = initViper()
	if err != nil {
		return
	}

	resetLogger()
	return
}
