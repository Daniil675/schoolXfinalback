package utility

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"math/rand"
	"time"
)

var (
	//Logger      *zap.Logger
	//Sugar       *zap.SugaredLogger
	letterRunes = []rune("01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	//cfg := zap.NewProductionConfig()
	//cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//cfg.OutputPaths = append(cfg.OutputPaths, "log.txt")
	//
	//
	//var err error
	//Logger, err = cfg.Build()
	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:   "./adbot.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28, // days
	//})
	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	//	w,
	//	zap.InfoLevel,
	//)
	//Logger = zap.New(core)
	//Logger.Info("successfully initialize zap logger")
	//Sugar = Logger.Sugar()
	//Sugar.Info("successfully initialize zap sugared logger")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/adbot.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})
	logrus.SetOutput(lumberjackLogger)
	rand.Seed(time.Now().UnixNano())
}

func Check(msg string, e error) {
	if e != nil {
		logrus.Error(msg, e)
	}
}
func CheckErr(e error) {
	if e != nil {
		logrus.Error(e)
	}
}
func Error(e error, data string) {
	logrus.WithField("data", data).Error(e)
}

func GenerateApiKey() string {
	return RandStringRunes(64)
}
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//func Contains(s []string, e string) bool {
//	for _, a := range s {
//		if a == e {
//			return true
//		}
//	}
//	return false
//}
func Contains(s []string, e string) bool {
	_, ok := ContainsWithId(s, e)
	return ok
}
func ContainsWithId(s []string, e string) (int, bool) {
	for i, a := range s {
		if a == e {
			return i, true
		}
	}
	return -1, false
}
func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
