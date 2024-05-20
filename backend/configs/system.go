package configs

type System struct {
	Port           string `mapstructure:"port" yaml:"port"`
	Ipv6           string `mapstructure:"ipv6" yaml:"ipv6"`
	BindAddress    string `mapstructure:"bindAddress" yaml:"bindAddress"`
	SSL            string `mapstructure:"ssl" yaml:"ssl"`
	DbFile         string `mapstructure:"db_file" yaml:"db_file"`
	DbPath         string `mapstructure:"db_path" yaml:"db_path"`
	LogPath        string `mapstructure:"log_path" yaml:"log_path"`
	DataDir        string `mapstructure:"data_dir" yaml:"data_dir"`
	TmpDir         string `mapstructure:"tmp_dir" yaml:"tmp_dir"`
	Cache          string `mapstructure:"cache" yaml:"cache"`
	Backup         string `mapstructure:"backup" yaml:"backup"`
	EncryptKey     string `mapstructure:"encrypt_key" yaml:"encrypt_key"`
	BaseDir        string `mapstructure:"base_dir" yaml:"base_dir"`
	Mode           string `mapstructure:"mode" yaml:"mode"`
	RepoUrl        string `mapstructure:"repo_url" yaml:"repo_url"`
	Version        string `mapstructure:"version" yaml:"version"`
	Username       string `mapstructure:"username" yaml:"username"`
	Password       string `mapstructure:"password" yaml:"password"`
	Entrance       string `mapstructure:"entrance" yaml:"entrance"`
	IsDemo         bool   `mapstructure:"is_demo" yaml:"is_demo"`
	AppRepo        string `mapstructure:"app_repo" yaml:"app_repo"`
	ChangeUserInfo string `mapstructure:"change_user_info" yaml:"change_user_info"`
	OneDriveID     string `mapstructure:"one_drive_id" yaml:"one_drive_id"`
	OneDriveSc     string `mapstructure:"one_drive_sc" yaml:"one_drive_sc"`
}
