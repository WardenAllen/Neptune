package config

type NoticeInfo struct {
	Channel			int
	Content			string
}

type NoticeListConf struct {
	List			[]NoticeInfo
}