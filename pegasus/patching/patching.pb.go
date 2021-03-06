// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/pegasus/patching/patching.proto
// DO NOT EDIT!

/*
Package patching is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/pegasus/patching/patching.proto

It has these top-level messages:
	DataOnlyPatching
	Telemetry
*/
package patching

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DataOnlyPatching_Status int32

const (
	DataOnlyPatching_SUCCEED                DataOnlyPatching_Status = 0
	DataOnlyPatching_SUCCEED_WITH_CACHE     DataOnlyPatching_Status = 1
	DataOnlyPatching_SUCCEED_WITH_TIMEOVER  DataOnlyPatching_Status = 2
	DataOnlyPatching_FAILED_GENERIC         DataOnlyPatching_Status = 3
	DataOnlyPatching_FAILED_DOWNLOADING     DataOnlyPatching_Status = 4
	DataOnlyPatching_FAILED_BAD_DATA        DataOnlyPatching_Status = 5
	DataOnlyPatching_FAILED_MD5_MISMATCH    DataOnlyPatching_Status = 6
	DataOnlyPatching_FAILED_BAD_ASSETBUNDLE DataOnlyPatching_Status = 7
	DataOnlyPatching_STARTED                DataOnlyPatching_Status = 8
)

var DataOnlyPatching_Status_name = map[int32]string{
	0: "SUCCEED",
	1: "SUCCEED_WITH_CACHE",
	2: "SUCCEED_WITH_TIMEOVER",
	3: "FAILED_GENERIC",
	4: "FAILED_DOWNLOADING",
	5: "FAILED_BAD_DATA",
	6: "FAILED_MD5_MISMATCH",
	7: "FAILED_BAD_ASSETBUNDLE",
	8: "STARTED",
}
var DataOnlyPatching_Status_value = map[string]int32{
	"SUCCEED":                0,
	"SUCCEED_WITH_CACHE":     1,
	"SUCCEED_WITH_TIMEOVER":  2,
	"FAILED_GENERIC":         3,
	"FAILED_DOWNLOADING":     4,
	"FAILED_BAD_DATA":        5,
	"FAILED_MD5_MISMATCH":    6,
	"FAILED_BAD_ASSETBUNDLE": 7,
	"STARTED":                8,
}

func (x DataOnlyPatching_Status) Enum() *DataOnlyPatching_Status {
	p := new(DataOnlyPatching_Status)
	*p = x
	return p
}
func (x DataOnlyPatching_Status) String() string {
	return proto.EnumName(DataOnlyPatching_Status_name, int32(x))
}
func (x *DataOnlyPatching_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataOnlyPatching_Status_value, data, "DataOnlyPatching_Status")
	if err != nil {
		return err
	}
	*x = DataOnlyPatching_Status(value)
	return nil
}

type DataOnlyPatching_Locale int32

const (
	DataOnlyPatching_UnknownLocale DataOnlyPatching_Locale = 0
	DataOnlyPatching_enUS          DataOnlyPatching_Locale = 1
	DataOnlyPatching_enGB          DataOnlyPatching_Locale = 2
	DataOnlyPatching_frFR          DataOnlyPatching_Locale = 3
	DataOnlyPatching_deDE          DataOnlyPatching_Locale = 4
	DataOnlyPatching_koKR          DataOnlyPatching_Locale = 5
	DataOnlyPatching_esES          DataOnlyPatching_Locale = 6
	DataOnlyPatching_esMX          DataOnlyPatching_Locale = 7
	DataOnlyPatching_ruRU          DataOnlyPatching_Locale = 8
	DataOnlyPatching_zhTW          DataOnlyPatching_Locale = 9
	DataOnlyPatching_zhCN          DataOnlyPatching_Locale = 10
	DataOnlyPatching_itIT          DataOnlyPatching_Locale = 11
	DataOnlyPatching_ptBR          DataOnlyPatching_Locale = 12
	DataOnlyPatching_plPL          DataOnlyPatching_Locale = 13
	DataOnlyPatching_Locale15      DataOnlyPatching_Locale = 15
	DataOnlyPatching_Locale16      DataOnlyPatching_Locale = 16
)

var DataOnlyPatching_Locale_name = map[int32]string{
	0:  "UnknownLocale",
	1:  "enUS",
	2:  "enGB",
	3:  "frFR",
	4:  "deDE",
	5:  "koKR",
	6:  "esES",
	7:  "esMX",
	8:  "ruRU",
	9:  "zhTW",
	10: "zhCN",
	11: "itIT",
	12: "ptBR",
	13: "plPL",
	15: "Locale15",
	16: "Locale16",
}
var DataOnlyPatching_Locale_value = map[string]int32{
	"UnknownLocale": 0,
	"enUS":          1,
	"enGB":          2,
	"frFR":          3,
	"deDE":          4,
	"koKR":          5,
	"esES":          6,
	"esMX":          7,
	"ruRU":          8,
	"zhTW":          9,
	"zhCN":          10,
	"itIT":          11,
	"ptBR":          12,
	"plPL":          13,
	"Locale15":      15,
	"Locale16":      16,
}

func (x DataOnlyPatching_Locale) Enum() *DataOnlyPatching_Locale {
	p := new(DataOnlyPatching_Locale)
	*p = x
	return p
}
func (x DataOnlyPatching_Locale) String() string {
	return proto.EnumName(DataOnlyPatching_Locale_name, int32(x))
}
func (x *DataOnlyPatching_Locale) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataOnlyPatching_Locale_value, data, "DataOnlyPatching_Locale")
	if err != nil {
		return err
	}
	*x = DataOnlyPatching_Locale(value)
	return nil
}

type DataOnlyPatching_Platform int32

const (
	DataOnlyPatching_UnknownPlatform DataOnlyPatching_Platform = 0
	DataOnlyPatching_Windows         DataOnlyPatching_Platform = 1
	DataOnlyPatching_Mac             DataOnlyPatching_Platform = 2
	DataOnlyPatching_iPad            DataOnlyPatching_Platform = 3
	DataOnlyPatching_iPhone          DataOnlyPatching_Platform = 4
	DataOnlyPatching_Android_Tablet  DataOnlyPatching_Platform = 5
	DataOnlyPatching_Android_Phone   DataOnlyPatching_Platform = 6
)

var DataOnlyPatching_Platform_name = map[int32]string{
	0: "UnknownPlatform",
	1: "Windows",
	2: "Mac",
	3: "iPad",
	4: "iPhone",
	5: "Android_Tablet",
	6: "Android_Phone",
}
var DataOnlyPatching_Platform_value = map[string]int32{
	"UnknownPlatform": 0,
	"Windows":         1,
	"Mac":             2,
	"iPad":            3,
	"iPhone":          4,
	"Android_Tablet":  5,
	"Android_Phone":   6,
}

func (x DataOnlyPatching_Platform) Enum() *DataOnlyPatching_Platform {
	p := new(DataOnlyPatching_Platform)
	*p = x
	return p
}
func (x DataOnlyPatching_Platform) String() string {
	return proto.EnumName(DataOnlyPatching_Platform_name, int32(x))
}
func (x *DataOnlyPatching_Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataOnlyPatching_Platform_value, data, "DataOnlyPatching_Platform")
	if err != nil {
		return err
	}
	*x = DataOnlyPatching_Platform(value)
	return nil
}

type DataOnlyPatching_BnetRegion int32

const (
	DataOnlyPatching_REGION_UNINITIALIZED     DataOnlyPatching_BnetRegion = -1
	DataOnlyPatching_REGION_UNKNOWN           DataOnlyPatching_BnetRegion = 0
	DataOnlyPatching_REGION_US                DataOnlyPatching_BnetRegion = 1
	DataOnlyPatching_REGION_EU                DataOnlyPatching_BnetRegion = 2
	DataOnlyPatching_REGION_KR                DataOnlyPatching_BnetRegion = 3
	DataOnlyPatching_REGION_TW                DataOnlyPatching_BnetRegion = 4
	DataOnlyPatching_REGION_CN                DataOnlyPatching_BnetRegion = 5
	DataOnlyPatching_REGION_LIVE_VERIFICATION DataOnlyPatching_BnetRegion = 40
	DataOnlyPatching_REGION_PTR_LOC           DataOnlyPatching_BnetRegion = 41
	DataOnlyPatching_REGION_DEV               DataOnlyPatching_BnetRegion = 60
)

var DataOnlyPatching_BnetRegion_name = map[int32]string{
	-1: "REGION_UNINITIALIZED",
	0:  "REGION_UNKNOWN",
	1:  "REGION_US",
	2:  "REGION_EU",
	3:  "REGION_KR",
	4:  "REGION_TW",
	5:  "REGION_CN",
	40: "REGION_LIVE_VERIFICATION",
	41: "REGION_PTR_LOC",
	60: "REGION_DEV",
}
var DataOnlyPatching_BnetRegion_value = map[string]int32{
	"REGION_UNINITIALIZED":     -1,
	"REGION_UNKNOWN":           0,
	"REGION_US":                1,
	"REGION_EU":                2,
	"REGION_KR":                3,
	"REGION_TW":                4,
	"REGION_CN":                5,
	"REGION_LIVE_VERIFICATION": 40,
	"REGION_PTR_LOC":           41,
	"REGION_DEV":               60,
}

func (x DataOnlyPatching_BnetRegion) Enum() *DataOnlyPatching_BnetRegion {
	p := new(DataOnlyPatching_BnetRegion)
	*p = x
	return p
}
func (x DataOnlyPatching_BnetRegion) String() string {
	return proto.EnumName(DataOnlyPatching_BnetRegion_name, int32(x))
}
func (x *DataOnlyPatching_BnetRegion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataOnlyPatching_BnetRegion_value, data, "DataOnlyPatching_BnetRegion")
	if err != nil {
		return err
	}
	*x = DataOnlyPatching_BnetRegion(value)
	return nil
}

type Telemetry_Level int32

const (
	Telemetry_LEVEL_NONE  Telemetry_Level = 0
	Telemetry_LEVEL_INFO  Telemetry_Level = 1
	Telemetry_LEVEL_WARN  Telemetry_Level = 2
	Telemetry_LEVEL_ERROR Telemetry_Level = 3
)

var Telemetry_Level_name = map[int32]string{
	0: "LEVEL_NONE",
	1: "LEVEL_INFO",
	2: "LEVEL_WARN",
	3: "LEVEL_ERROR",
}
var Telemetry_Level_value = map[string]int32{
	"LEVEL_NONE":  0,
	"LEVEL_INFO":  1,
	"LEVEL_WARN":  2,
	"LEVEL_ERROR": 3,
}

func (x Telemetry_Level) Enum() *Telemetry_Level {
	p := new(Telemetry_Level)
	*p = x
	return p
}
func (x Telemetry_Level) String() string {
	return proto.EnumName(Telemetry_Level_name, int32(x))
}
func (x *Telemetry_Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_Level_value, data, "Telemetry_Level")
	if err != nil {
		return err
	}
	*x = Telemetry_Level(value)
	return nil
}

type Telemetry_Locale int32

const (
	Telemetry_LOCALE_UNKNOWN Telemetry_Locale = 0
	Telemetry_LOCALE_ENUS    Telemetry_Locale = 1
	Telemetry_LOCALE_ENGB    Telemetry_Locale = 2
	Telemetry_LOCALE_FRFR    Telemetry_Locale = 3
	Telemetry_LOCALE_DEDE    Telemetry_Locale = 4
	Telemetry_LOCALE_KOKR    Telemetry_Locale = 5
	Telemetry_LOCALE_ESES    Telemetry_Locale = 6
	Telemetry_LOCALE_ESMX    Telemetry_Locale = 7
	Telemetry_LOCALE_RURU    Telemetry_Locale = 8
	Telemetry_LOCALE_ZHTW    Telemetry_Locale = 9
	Telemetry_LOCALE_ZHCN    Telemetry_Locale = 10
	Telemetry_LOCALE_ITIT    Telemetry_Locale = 11
	Telemetry_LOCALE_PTBR    Telemetry_Locale = 12
	Telemetry_LOCALE_PLPL    Telemetry_Locale = 13
	Telemetry_LOCALE_15      Telemetry_Locale = 15
	Telemetry_LOCALE_16      Telemetry_Locale = 16
)

var Telemetry_Locale_name = map[int32]string{
	0:  "LOCALE_UNKNOWN",
	1:  "LOCALE_ENUS",
	2:  "LOCALE_ENGB",
	3:  "LOCALE_FRFR",
	4:  "LOCALE_DEDE",
	5:  "LOCALE_KOKR",
	6:  "LOCALE_ESES",
	7:  "LOCALE_ESMX",
	8:  "LOCALE_RURU",
	9:  "LOCALE_ZHTW",
	10: "LOCALE_ZHCN",
	11: "LOCALE_ITIT",
	12: "LOCALE_PTBR",
	13: "LOCALE_PLPL",
	15: "LOCALE_15",
	16: "LOCALE_16",
}
var Telemetry_Locale_value = map[string]int32{
	"LOCALE_UNKNOWN": 0,
	"LOCALE_ENUS":    1,
	"LOCALE_ENGB":    2,
	"LOCALE_FRFR":    3,
	"LOCALE_DEDE":    4,
	"LOCALE_KOKR":    5,
	"LOCALE_ESES":    6,
	"LOCALE_ESMX":    7,
	"LOCALE_RURU":    8,
	"LOCALE_ZHTW":    9,
	"LOCALE_ZHCN":    10,
	"LOCALE_ITIT":    11,
	"LOCALE_PTBR":    12,
	"LOCALE_PLPL":    13,
	"LOCALE_15":      15,
	"LOCALE_16":      16,
}

func (x Telemetry_Locale) Enum() *Telemetry_Locale {
	p := new(Telemetry_Locale)
	*p = x
	return p
}
func (x Telemetry_Locale) String() string {
	return proto.EnumName(Telemetry_Locale_name, int32(x))
}
func (x *Telemetry_Locale) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_Locale_value, data, "Telemetry_Locale")
	if err != nil {
		return err
	}
	*x = Telemetry_Locale(value)
	return nil
}

type Telemetry_Platform int32

const (
	Telemetry_PLATFORM_UNKNOWN Telemetry_Platform = 0
	Telemetry_PLATFORM_PC      Telemetry_Platform = 1
	Telemetry_PLATFORM_MAC     Telemetry_Platform = 2
	Telemetry_PLATFORM_IOS     Telemetry_Platform = 3
	Telemetry_PLATFORM_ANDROID Telemetry_Platform = 4
)

var Telemetry_Platform_name = map[int32]string{
	0: "PLATFORM_UNKNOWN",
	1: "PLATFORM_PC",
	2: "PLATFORM_MAC",
	3: "PLATFORM_IOS",
	4: "PLATFORM_ANDROID",
}
var Telemetry_Platform_value = map[string]int32{
	"PLATFORM_UNKNOWN": 0,
	"PLATFORM_PC":      1,
	"PLATFORM_MAC":     2,
	"PLATFORM_IOS":     3,
	"PLATFORM_ANDROID": 4,
}

func (x Telemetry_Platform) Enum() *Telemetry_Platform {
	p := new(Telemetry_Platform)
	*p = x
	return p
}
func (x Telemetry_Platform) String() string {
	return proto.EnumName(Telemetry_Platform_name, int32(x))
}
func (x *Telemetry_Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_Platform_value, data, "Telemetry_Platform")
	if err != nil {
		return err
	}
	*x = Telemetry_Platform(value)
	return nil
}

type Telemetry_ScreenUI int32

const (
	Telemetry_SCREENUI_UNKNOWN Telemetry_ScreenUI = 0
	Telemetry_SCREENUI_DESKTOP Telemetry_ScreenUI = 1
	Telemetry_SCREENUI_TABLET  Telemetry_ScreenUI = 2
	Telemetry_SCREENUI_PHONE   Telemetry_ScreenUI = 3
)

var Telemetry_ScreenUI_name = map[int32]string{
	0: "SCREENUI_UNKNOWN",
	1: "SCREENUI_DESKTOP",
	2: "SCREENUI_TABLET",
	3: "SCREENUI_PHONE",
}
var Telemetry_ScreenUI_value = map[string]int32{
	"SCREENUI_UNKNOWN": 0,
	"SCREENUI_DESKTOP": 1,
	"SCREENUI_TABLET":  2,
	"SCREENUI_PHONE":   3,
}

func (x Telemetry_ScreenUI) Enum() *Telemetry_ScreenUI {
	p := new(Telemetry_ScreenUI)
	*p = x
	return p
}
func (x Telemetry_ScreenUI) String() string {
	return proto.EnumName(Telemetry_ScreenUI_name, int32(x))
}
func (x *Telemetry_ScreenUI) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_ScreenUI_value, data, "Telemetry_ScreenUI")
	if err != nil {
		return err
	}
	*x = Telemetry_ScreenUI(value)
	return nil
}

type Telemetry_Store int32

const (
	Telemetry_STORE_BLIZZARD   Telemetry_Store = 0
	Telemetry_STORE_IOS        Telemetry_Store = 1
	Telemetry_STORE_GOOGLEPLAY Telemetry_Store = 2
	Telemetry_STORE_AMAZON     Telemetry_Store = 3
)

var Telemetry_Store_name = map[int32]string{
	0: "STORE_BLIZZARD",
	1: "STORE_IOS",
	2: "STORE_GOOGLEPLAY",
	3: "STORE_AMAZON",
}
var Telemetry_Store_value = map[string]int32{
	"STORE_BLIZZARD":   0,
	"STORE_IOS":        1,
	"STORE_GOOGLEPLAY": 2,
	"STORE_AMAZON":     3,
}

func (x Telemetry_Store) Enum() *Telemetry_Store {
	p := new(Telemetry_Store)
	*p = x
	return p
}
func (x Telemetry_Store) String() string {
	return proto.EnumName(Telemetry_Store_name, int32(x))
}
func (x *Telemetry_Store) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_Store_value, data, "Telemetry_Store")
	if err != nil {
		return err
	}
	*x = Telemetry_Store(value)
	return nil
}

type Telemetry_BnetRegion int32

const (
	Telemetry_REGION_UNINITIALIZED     Telemetry_BnetRegion = -1
	Telemetry_REGION_UNKNOWN           Telemetry_BnetRegion = 0
	Telemetry_REGION_US                Telemetry_BnetRegion = 1
	Telemetry_REGION_EU                Telemetry_BnetRegion = 2
	Telemetry_REGION_KR                Telemetry_BnetRegion = 3
	Telemetry_REGION_TW                Telemetry_BnetRegion = 4
	Telemetry_REGION_CN                Telemetry_BnetRegion = 5
	Telemetry_REGION_LIVE_VERIFICATION Telemetry_BnetRegion = 40
	Telemetry_REGION_PTR_LOC           Telemetry_BnetRegion = 41
	Telemetry_REGION_DEV               Telemetry_BnetRegion = 60
	Telemetry_REGION_PTR               Telemetry_BnetRegion = 98
)

var Telemetry_BnetRegion_name = map[int32]string{
	-1: "REGION_UNINITIALIZED",
	0:  "REGION_UNKNOWN",
	1:  "REGION_US",
	2:  "REGION_EU",
	3:  "REGION_KR",
	4:  "REGION_TW",
	5:  "REGION_CN",
	40: "REGION_LIVE_VERIFICATION",
	41: "REGION_PTR_LOC",
	60: "REGION_DEV",
	98: "REGION_PTR",
}
var Telemetry_BnetRegion_value = map[string]int32{
	"REGION_UNINITIALIZED":     -1,
	"REGION_UNKNOWN":           0,
	"REGION_US":                1,
	"REGION_EU":                2,
	"REGION_KR":                3,
	"REGION_TW":                4,
	"REGION_CN":                5,
	"REGION_LIVE_VERIFICATION": 40,
	"REGION_PTR_LOC":           41,
	"REGION_DEV":               60,
	"REGION_PTR":               98,
}

func (x Telemetry_BnetRegion) Enum() *Telemetry_BnetRegion {
	p := new(Telemetry_BnetRegion)
	*p = x
	return p
}
func (x Telemetry_BnetRegion) String() string {
	return proto.EnumName(Telemetry_BnetRegion_name, int32(x))
}
func (x *Telemetry_BnetRegion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Telemetry_BnetRegion_value, data, "Telemetry_BnetRegion")
	if err != nil {
		return err
	}
	*x = Telemetry_BnetRegion(value)
	return nil
}

type DataOnlyPatching struct {
	Status                 *DataOnlyPatching_Status     `protobuf:"varint,1,req,name=status,enum=patching.DataOnlyPatching_Status" json:"status,omitempty"`
	Locale                 *DataOnlyPatching_Locale     `protobuf:"varint,2,req,name=locale,enum=patching.DataOnlyPatching_Locale" json:"locale,omitempty"`
	Platform               *DataOnlyPatching_Platform   `protobuf:"varint,3,req,name=platform,enum=patching.DataOnlyPatching_Platform" json:"platform,omitempty"`
	BnetRegion             *DataOnlyPatching_BnetRegion `protobuf:"varint,4,req,name=bnet_region,enum=patching.DataOnlyPatching_BnetRegion" json:"bnet_region,omitempty"`
	GameAccountId          *uint64                      `protobuf:"varint,5,req,name=game_account_id" json:"game_account_id,omitempty"`
	CurrentBuild           *int32                       `protobuf:"varint,6,opt,name=current_build" json:"current_build,omitempty"`
	NewBuild               *int32                       `protobuf:"varint,7,opt,name=new_build" json:"new_build,omitempty"`
	SessionId              *string                      `protobuf:"bytes,8,req,name=session_id" json:"session_id,omitempty"`
	DeviceUniqueIdentifier *string                      `protobuf:"bytes,9,req,name=device_unique_identifier" json:"device_unique_identifier,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *DataOnlyPatching) Reset()         { *m = DataOnlyPatching{} }
func (m *DataOnlyPatching) String() string { return proto.CompactTextString(m) }
func (*DataOnlyPatching) ProtoMessage()    {}

func (m *DataOnlyPatching) GetStatus() DataOnlyPatching_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return DataOnlyPatching_SUCCEED
}

func (m *DataOnlyPatching) GetLocale() DataOnlyPatching_Locale {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return DataOnlyPatching_UnknownLocale
}

func (m *DataOnlyPatching) GetPlatform() DataOnlyPatching_Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return DataOnlyPatching_UnknownPlatform
}

func (m *DataOnlyPatching) GetBnetRegion() DataOnlyPatching_BnetRegion {
	if m != nil && m.BnetRegion != nil {
		return *m.BnetRegion
	}
	return DataOnlyPatching_REGION_UNINITIALIZED
}

func (m *DataOnlyPatching) GetGameAccountId() uint64 {
	if m != nil && m.GameAccountId != nil {
		return *m.GameAccountId
	}
	return 0
}

func (m *DataOnlyPatching) GetCurrentBuild() int32 {
	if m != nil && m.CurrentBuild != nil {
		return *m.CurrentBuild
	}
	return 0
}

func (m *DataOnlyPatching) GetNewBuild() int32 {
	if m != nil && m.NewBuild != nil {
		return *m.NewBuild
	}
	return 0
}

func (m *DataOnlyPatching) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *DataOnlyPatching) GetDeviceUniqueIdentifier() string {
	if m != nil && m.DeviceUniqueIdentifier != nil {
		return *m.DeviceUniqueIdentifier
	}
	return ""
}

type Telemetry struct {
	Time                   *int64                `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Level                  *Telemetry_Level      `protobuf:"varint,2,req,name=level,enum=patching.Telemetry_Level" json:"level,omitempty"`
	Locale                 *Telemetry_Locale     `protobuf:"varint,3,req,name=locale,enum=patching.Telemetry_Locale" json:"locale,omitempty"`
	Version                *string               `protobuf:"bytes,4,req,name=version" json:"version,omitempty"`
	Platform               *Telemetry_Platform   `protobuf:"varint,5,req,name=platform,enum=patching.Telemetry_Platform" json:"platform,omitempty"`
	Os                     *string               `protobuf:"bytes,6,req,name=os" json:"os,omitempty"`
	ScreenUI               *Telemetry_ScreenUI   `protobuf:"varint,7,req,name=screen_u_i,enum=patching.Telemetry_ScreenUI" json:"screen_u_i,omitempty"`
	Store                  *Telemetry_Store      `protobuf:"varint,8,req,name=store,enum=patching.Telemetry_Store" json:"store,omitempty"`
	SessionId              *string               `protobuf:"bytes,9,req,name=session_id" json:"session_id,omitempty"`
	DeviceUniqueIdentifier *string               `protobuf:"bytes,10,req,name=device_unique_identifier" json:"device_unique_identifier,omitempty"`
	Event                  *uint64               `protobuf:"varint,11,req,name=event" json:"event,omitempty"`
	BnetRegion             *Telemetry_BnetRegion `protobuf:"varint,12,opt,name=bnet_region,enum=patching.Telemetry_BnetRegion,def=-1" json:"bnet_region,omitempty"`
	GameAccountId          *uint64               `protobuf:"varint,13,opt,name=game_account_id" json:"game_account_id,omitempty"`
	ErrorCode              *int64                `protobuf:"varint,14,opt,name=error_code" json:"error_code,omitempty"`
	Message                *string               `protobuf:"bytes,15,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized       []byte                `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}

const Default_Telemetry_BnetRegion Telemetry_BnetRegion = Telemetry_REGION_UNINITIALIZED

func (m *Telemetry) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Telemetry) GetLevel() Telemetry_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Telemetry_LEVEL_NONE
}

func (m *Telemetry) GetLocale() Telemetry_Locale {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return Telemetry_LOCALE_UNKNOWN
}

func (m *Telemetry) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *Telemetry) GetPlatform() Telemetry_Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return Telemetry_PLATFORM_UNKNOWN
}

func (m *Telemetry) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *Telemetry) GetScreenUI() Telemetry_ScreenUI {
	if m != nil && m.ScreenUI != nil {
		return *m.ScreenUI
	}
	return Telemetry_SCREENUI_UNKNOWN
}

func (m *Telemetry) GetStore() Telemetry_Store {
	if m != nil && m.Store != nil {
		return *m.Store
	}
	return Telemetry_STORE_BLIZZARD
}

func (m *Telemetry) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *Telemetry) GetDeviceUniqueIdentifier() string {
	if m != nil && m.DeviceUniqueIdentifier != nil {
		return *m.DeviceUniqueIdentifier
	}
	return ""
}

func (m *Telemetry) GetEvent() uint64 {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return 0
}

func (m *Telemetry) GetBnetRegion() Telemetry_BnetRegion {
	if m != nil && m.BnetRegion != nil {
		return *m.BnetRegion
	}
	return Default_Telemetry_BnetRegion
}

func (m *Telemetry) GetGameAccountId() uint64 {
	if m != nil && m.GameAccountId != nil {
		return *m.GameAccountId
	}
	return 0
}

func (m *Telemetry) GetErrorCode() int64 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *Telemetry) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("patching.DataOnlyPatching_Status", DataOnlyPatching_Status_name, DataOnlyPatching_Status_value)
	proto.RegisterEnum("patching.DataOnlyPatching_Locale", DataOnlyPatching_Locale_name, DataOnlyPatching_Locale_value)
	proto.RegisterEnum("patching.DataOnlyPatching_Platform", DataOnlyPatching_Platform_name, DataOnlyPatching_Platform_value)
	proto.RegisterEnum("patching.DataOnlyPatching_BnetRegion", DataOnlyPatching_BnetRegion_name, DataOnlyPatching_BnetRegion_value)
	proto.RegisterEnum("patching.Telemetry_Level", Telemetry_Level_name, Telemetry_Level_value)
	proto.RegisterEnum("patching.Telemetry_Locale", Telemetry_Locale_name, Telemetry_Locale_value)
	proto.RegisterEnum("patching.Telemetry_Platform", Telemetry_Platform_name, Telemetry_Platform_value)
	proto.RegisterEnum("patching.Telemetry_ScreenUI", Telemetry_ScreenUI_name, Telemetry_ScreenUI_value)
	proto.RegisterEnum("patching.Telemetry_Store", Telemetry_Store_name, Telemetry_Store_value)
	proto.RegisterEnum("patching.Telemetry_BnetRegion", Telemetry_BnetRegion_name, Telemetry_BnetRegion_value)
}
