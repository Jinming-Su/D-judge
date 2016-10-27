package config

var GlobalConfig SystemConfig

type SystemConfig struct {
	HostName         string `toml:"host_name"`
	EndpointUser     string `toml:"endpoint_user"`
	EndpointName     string `toml:"endpoint_name"`
	EndpointURL      string `toml:"endpoint_url"`
	MaxCacheSize     int    `toml:"max_cache_size"`
	EndpointPassword string `toml:"endpoint_password"`
	JudgeRoot        string `toml:"judge_root"`
	DockerImage      string `toml:"docker_image"`
	DockerServer     string `toml:"docker_server"`
	CacheRoot        string `toml:"cache_root"`
}

type JudgeInfo struct {
	SubmitID      int64  `json:"submitid"`
	ContestID     int64  `json:"cid"`
	TeamID        int64  `json:"teamid"`
	JudgingID     int64  `json:"judgingid"`
	ProblemID     int64  `json:"probid"`
	Language      string `json:"langid"`
	TimeLimit     int64  `json:"maxruntime"`
	MemLimit      int64  `json:"memlimit"`
	OutputLimit   int64  `json:"output_limit"`
	BuildZip      string `json:"compile_script"`
	BuildZipMD5   string `json:"compile_script_md5sum"`
	RunZip        string `json:"run"`
	RunScriptMD5  string `json:"run_md5sum"`
	CompareZip    string `json:"compare"`
	CompareZipMD5 string `json:"compare_script_md5sum"`
	CompareArgs   string `json:"compare_args"`
}

type TestcaseInfo struct {
	TestcaseID   int64  `json:"testcaseid"`
	Rank         int64  `json:"rank"`
	ProblemID    int64  `json:"probid"`
	MD5SumInput  string `json:"md5sum_input"`
	MD5SumOutput string `json:"md5sum_input"`
}

type SubmissionInfo struct {
	info []SubmissionFileInfo `json:""`
}

type SubmissionFileInfo struct {
	FileName string `json:"filename"`
	Content  string `json:"contetn"`
}
