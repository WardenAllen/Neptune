package config

import "strings"

type StatEvtType int

const (
	StatEvtNone      	StatEvtType = 0
	StatEvtInt       	StatEvtType = 1
	StatEvtFloat     	StatEvtType = 2
	StatEvtString    	StatEvtType = 3
	StatEvtTimeStamp 	StatEvtType = 4
	StatEvtIP        	StatEvtType = 5
	StatEvtArray     	StatEvtType = 6
	StatEvtAmount    	StatEvtType = 7
)

const (
	StatEvtPriTimeStamp	= "_TS_"
	StatEvtPriIP		= "_IP_"
	StatEvtPriAry		= "_ARY_"
	StatEvtPriAmount	= "_AMT_"
)

type StatEventInfo struct {
	ID        			int
	Name      			string
	IntList   			string
	FloatList 			string
	StrList   			string
	Type      			StatEvtType
}

type StatEventItem struct {
	Name				string
	Type				StatEvtType
}

type StatEvent struct {
	ID					int
	Name				string
	IntList				[]StatEventItem
	FloatList			[]StatEventItem
	StrList				[]StatEventItem
}

type StatEventConf struct {
	List				[]StatEventInfo
}

type StatPtyType int

const (
	StatPtyNone      	StatPtyType = 0
	StatPtyInt       	StatPtyType = 1
	StatPtyFloat     	StatPtyType = 2
	StatPtyString    	StatPtyType = 3
	StatPtyTimeStamp 	StatPtyType = 4
	StatPtyIP        	StatPtyType = 5
	StatPtyArray     	StatPtyType = 6
	StatPtyAmount    	StatPtyType = 7
)

type StatProperty struct {
	ID   				int
	Name 				string
	Type 				StatPtyType
	Unique				int64
}

type StatPropertyConf struct {
	List				[]StatProperty
}

func GetRealName(name string) string {

	return strings.Split(name, "_")[2]

}

func CheckEvtItemType(name *string) StatEvtType {

	if strings.Contains(*name, StatEvtPriTimeStamp) {
		*name = GetRealName(*name)
		return StatEvtTimeStamp
	}

	if strings.Contains(*name, StatEvtPriIP) {
		*name = GetRealName(*name)
		return StatEvtIP
	}

	if strings.Contains(*name, StatEvtPriAry) {
		*name = GetRealName(*name)
		return StatEvtArray
	}

	if strings.Contains(*name, StatEvtPriAmount) {
		*name = GetRealName(*name)
		return StatEvtAmount
	}

	return StatEvtNone

}
