package setting

import "time"

type ServerSettings struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
}
type MysqlSettings struct {
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSections(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
