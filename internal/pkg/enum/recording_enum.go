package enum

type PostRecord uint

const (
	POST_RECORD_SECONDS_15 PostRecord = iota
	POST_RECORD_SECONDS_30
	POST_RECORD_MINUTE_1
)

func (pr PostRecord) Value() string {
	return []string{"15 Seconds", "30 Seconds", "1 Minute"}[pr]
}

type RecordingProfile uint

const (
	RECORDING_BASE RecordingProfile = iota
	RECORDING_MAIN
	RECORDING_HIGH
)

func (ep RecordingProfile) Value() string {
	return []string{"Base", "Main", "High"}[ep]
}

type MainBitRate uint

const (
	MAIN_BIT_RATE_1024 MainBitRate = iota
	MAIN_BIT_RATE_1536
	MAIN_BIT_RATE_2048
	MAIN_BIT_RATE_3072
	MAIN_BIT_RATE_4096
	MAIN_BIT_RATE_5120
	MAIN_BIT_RATE_6144
	MAIN_BIT_RATE_7168
	MAIN_BIT_RATE_8192
)

func (mbr MainBitRate) Value() int {
	return []int{1024, 1536, 2048, 3072, 4096, 5120, 6144, 7168, 8192}[mbr]
}

type SubBitRate uint

const (
	SUB_BIT_RATE_64 SubBitRate = iota
	SUB_BIT_RATE_128
	SUB_BIT_RATE_160
	SUB_BIT_RATE_192
	SUB_BIT_RATE_256
	SUB_BIT_RATE_384
	SUB_BIT_RATE_512
)

func (sbr SubBitRate) Value() int {
	return []int{64, 128, 160, 192, 256, 384, 512}[sbr]
}

type MainFrameRate uint

const (
	MAIN_FRAME_RATE_2 MainFrameRate = iota
	MAIN_FRAME_RATE_4
	MAIN_FRAME_RATE_6
	MAIN_FRAME_RATE_8
	MAIN_FRAME_RATE_10
	MAIN_FRAME_RATE_12
	MAIN_FRAME_RATE_15
	MAIN_FRAME_RATE_16
	MAIN_FRAME_RATE_18
	MAIN_FRAME_RATE_20
)

func (mfr MainFrameRate) Value() int {
	return []int{2, 4, 6, 8, 10, 12, 15, 16, 18, 20}[mfr]
}

type SubFrameRate uint

const (
	SUB_FRAME_RATE_4 SubFrameRate = iota
	SUB_FRAME_RATE_7
	SUB_FRAME_RATE_10
	SUB_FRAME_RATE_15
)

func (sfr SubFrameRate) Value() int {
	return []int{4, 7, 10, 15}[sfr]
}

type MainSize uint

const (
	MAIN_SIZE_3072_1728 MainSize = iota
	MAIN_SIZE_2592_1944
	MAIN_SIZE_2560_1440
	MAIN_SIZE_2048_1536
	MAIN_SIZE_2304_1296
)

func (ms MainSize) Value() string {
	return []string{"3072*1728", "2592*1944", "2560*1440", "2048*1536", "2304*1296"}[ms]
}

type SubSize uint

const (
	SUB_SIZE_640_360 SubSize = iota
)

func (ss SubSize) Value() string {
	return []string{"640*360"}[ss]
}
