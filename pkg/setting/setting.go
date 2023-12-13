package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func InitSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs/")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
