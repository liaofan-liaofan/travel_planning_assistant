package utils

//func GetWriteSyncer() (zapcore.WriteSyncer, error) {
//	fileWriter, err := zaprotatelogs.New(
//		path.Join(global.TPA_CONFIG.Zap.Director, "%Y-%m-%d.log"),
//		zaprotatelogs.WithMaxAge(7*24*time.Hour),
//		zaprotatelogs.WithRotationTime(24*time.Hour),
//	)
//	if global.TPA_CONFIG.Zap.LogInConsole {
//		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
//	}
//	return zapcore.AddSync(fileWriter), err
//}
