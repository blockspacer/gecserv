// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type Color int8

const (
	ColorRed    Color = 0
	ColorGreen  Color = 1
	ColorBlue   Color = 2
	ColorYellow Color = 3
	ColorPink   Color = 4
	ColorGray   Color = 5
	ColorOrange Color = 6
)

var EnumNamesColor = map[Color]string{
	ColorRed:    "Red",
	ColorGreen:  "Green",
	ColorBlue:   "Blue",
	ColorYellow: "Yellow",
	ColorPink:   "Pink",
	ColorGray:   "Gray",
	ColorOrange: "Orange",
}

var EnumValuesColor = map[string]Color{
	"Red":    ColorRed,
	"Green":  ColorGreen,
	"Blue":   ColorBlue,
	"Yellow": ColorYellow,
	"Pink":   ColorPink,
	"Gray":   ColorGray,
	"Orange": ColorOrange,
}

func (v Color) String() string {
	if s, ok := EnumNamesColor[v]; ok {
		return s
	}
	return "Color(" + strconv.FormatInt(int64(v), 10) + ")"
}

type GameMessage byte

const (
	GameMessageNONE      GameMessage = 0
	GameMessageMapUpdate GameMessage = 1
)

var EnumNamesGameMessage = map[GameMessage]string{
	GameMessageNONE:      "NONE",
	GameMessageMapUpdate: "MapUpdate",
}

var EnumValuesGameMessage = map[string]GameMessage{
	"NONE":      GameMessageNONE,
	"MapUpdate": GameMessageMapUpdate,
}

func (v GameMessage) String() string {
	if s, ok := EnumNamesGameMessage[v]; ok {
		return s
	}
	return "GameMessage(" + strconv.FormatInt(int64(v), 10) + ")"
}

type GameMessageT struct {
	Type GameMessage
	Value interface{}
}

func (t *GameMessageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case GameMessageMapUpdate:
		return t.Value.(*MapUpdateT).Pack(builder)
	}
	return 0
}

func (rcv GameMessage) UnPack(table flatbuffers.Table) *GameMessageT {
	switch rcv {
	case GameMessageMapUpdate:
		x := MapUpdate{_tab: table}
		return &GameMessageT{ Type: GameMessageMapUpdate, Value: x.UnPack() }
	}
	return nil
}

type PlayerAction byte

const (
	PlayerActionW_DOWN  PlayerAction = 0
	PlayerActionW_UP    PlayerAction = 1
	PlayerActionA_DOWN  PlayerAction = 2
	PlayerActionA_UP    PlayerAction = 3
	PlayerActionS_DOWN  PlayerAction = 4
	PlayerActionS_UP    PlayerAction = 5
	PlayerActionD_DOWN  PlayerAction = 6
	PlayerActionD_UP    PlayerAction = 7
	PlayerActionM1_DOWN PlayerAction = 8
	PlayerActionM1_UP   PlayerAction = 9
	PlayerActionM2_DOWN PlayerAction = 10
	PlayerActionM2_UP   PlayerAction = 11
)

var EnumNamesPlayerAction = map[PlayerAction]string{
	PlayerActionW_DOWN:  "W_DOWN",
	PlayerActionW_UP:    "W_UP",
	PlayerActionA_DOWN:  "A_DOWN",
	PlayerActionA_UP:    "A_UP",
	PlayerActionS_DOWN:  "S_DOWN",
	PlayerActionS_UP:    "S_UP",
	PlayerActionD_DOWN:  "D_DOWN",
	PlayerActionD_UP:    "D_UP",
	PlayerActionM1_DOWN: "M1_DOWN",
	PlayerActionM1_UP:   "M1_UP",
	PlayerActionM2_DOWN: "M2_DOWN",
	PlayerActionM2_UP:   "M2_UP",
}

var EnumValuesPlayerAction = map[string]PlayerAction{
	"W_DOWN":  PlayerActionW_DOWN,
	"W_UP":    PlayerActionW_UP,
	"A_DOWN":  PlayerActionA_DOWN,
	"A_UP":    PlayerActionA_UP,
	"S_DOWN":  PlayerActionS_DOWN,
	"S_UP":    PlayerActionS_UP,
	"D_DOWN":  PlayerActionD_DOWN,
	"D_UP":    PlayerActionD_UP,
	"M1_DOWN": PlayerActionM1_DOWN,
	"M1_UP":   PlayerActionM1_UP,
	"M2_DOWN": PlayerActionM2_DOWN,
	"M2_UP":   PlayerActionM2_UP,
}

func (v PlayerAction) String() string {
	if s, ok := EnumNamesPlayerAction[v]; ok {
		return s
	}
	return "PlayerAction(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Vec2T struct {
	X float32
	Y float32
}

func (t *Vec2T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateVec2(builder, t.X, t.Y)
}
func (rcv *Vec2) UnPackTo(t *Vec2T) {
	t.X = rcv.X()
	t.Y = rcv.Y()
}

func (rcv *Vec2) UnPack() *Vec2T {
	if rcv == nil { return nil }
	t := &Vec2T{}
	rcv.UnPackTo(t)
	return t
}

type Vec2 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec2) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec2) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Vec2) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Vec2) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Vec2) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateVec2(builder *flatbuffers.Builder, x float32, y float32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
type Vec3T struct {
	X float32
	Y float32
	Z float32
}

func (t *Vec3T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateVec3(builder, t.X, t.Y, t.Z)
}
func (rcv *Vec3) UnPackTo(t *Vec3T) {
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Z = rcv.Z()
}

func (rcv *Vec3) UnPack() *Vec3T {
	if rcv == nil { return nil }
	t := &Vec3T{}
	rcv.UnPackTo(t)
	return t
}

type Vec3 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec3) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec3) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec3) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Vec3) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Vec3) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Vec3) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Vec3) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Vec3) MutateZ(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateVec3(builder *flatbuffers.Builder, x float32, y float32, z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependFloat32(z)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
type MapT struct {
	Name string
	GlobalX int32
	GlobalY int32
	MaxX int32
	MaxY int32
}

func (t *MapT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	MapStart(builder)
	MapAddName(builder, nameOffset)
	MapAddGlobalX(builder, t.GlobalX)
	MapAddGlobalY(builder, t.GlobalY)
	MapAddMaxX(builder, t.MaxX)
	MapAddMaxY(builder, t.MaxY)
	return MapEnd(builder)
}

func (rcv *Map) UnPackTo(t *MapT) {
	t.Name = string(rcv.Name())
	t.GlobalX = rcv.GlobalX()
	t.GlobalY = rcv.GlobalY()
	t.MaxX = rcv.MaxX()
	t.MaxY = rcv.MaxY()
}

func (rcv *Map) UnPack() *MapT {
	if rcv == nil { return nil }
	t := &MapT{}
	rcv.UnPackTo(t)
	return t
}

type Map struct {
	_tab flatbuffers.Table
}

func GetRootAsMap(buf []byte, offset flatbuffers.UOffsetT) *Map {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Map{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Map) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Map) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Map) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Map) GlobalX() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateGlobalX(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Map) GlobalY() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateGlobalY(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Map) MaxX() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateMaxX(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Map) MaxY() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateMaxY(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func MapStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MapAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func MapAddGlobalX(builder *flatbuffers.Builder, globalX int32) {
	builder.PrependInt32Slot(1, globalX, 0)
}
func MapAddGlobalY(builder *flatbuffers.Builder, globalY int32) {
	builder.PrependInt32Slot(2, globalY, 0)
}
func MapAddMaxX(builder *flatbuffers.Builder, maxX int32) {
	builder.PrependInt32Slot(3, maxX, 0)
}
func MapAddMaxY(builder *flatbuffers.Builder, maxY int32) {
	builder.PrependInt32Slot(4, maxY, 0)
}
func MapEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type PlayerT struct {
	Posx float32
	Posy float32
	Sid float64
	Col Color
	LastAck uint32
}

func (t *PlayerT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	PlayerStart(builder)
	PlayerAddPosx(builder, t.Posx)
	PlayerAddPosy(builder, t.Posy)
	PlayerAddSid(builder, t.Sid)
	PlayerAddCol(builder, t.Col)
	PlayerAddLastAck(builder, t.LastAck)
	return PlayerEnd(builder)
}

func (rcv *Player) UnPackTo(t *PlayerT) {
	t.Posx = rcv.Posx()
	t.Posy = rcv.Posy()
	t.Sid = rcv.Sid()
	t.Col = rcv.Col()
	t.LastAck = rcv.LastAck()
}

func (rcv *Player) UnPack() *PlayerT {
	if rcv == nil { return nil }
	t := &PlayerT{}
	rcv.UnPackTo(t)
	return t
}

type Player struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayer(buf []byte, offset flatbuffers.UOffsetT) *Player {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Player{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Player) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Player) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Player) Posx() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Player) MutatePosx(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *Player) Posy() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Player) MutatePosy(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *Player) Sid() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Player) MutateSid(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *Player) Col() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Player) MutateCol(n Color) bool {
	return rcv._tab.MutateInt8Slot(10, int8(n))
}

func (rcv *Player) LastAck() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Player) MutateLastAck(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func PlayerStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func PlayerAddPosx(builder *flatbuffers.Builder, posx float32) {
	builder.PrependFloat32Slot(0, posx, 0.0)
}
func PlayerAddPosy(builder *flatbuffers.Builder, posy float32) {
	builder.PrependFloat32Slot(1, posy, 0.0)
}
func PlayerAddSid(builder *flatbuffers.Builder, sid float64) {
	builder.PrependFloat64Slot(2, sid, 0.0)
}
func PlayerAddCol(builder *flatbuffers.Builder, col Color) {
	builder.PrependInt8Slot(3, int8(col), 0)
}
func PlayerAddLastAck(builder *flatbuffers.Builder, lastAck uint32) {
	builder.PrependUint32Slot(4, lastAck, 0)
}
func PlayerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type MessageT struct {
	Data *GameMessageT
}

func (t *MessageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := t.Data.Pack(builder)
	
	MessageStart(builder)
	if t.Data != nil {
		MessageAddDataType(builder, t.Data.Type)
	}
	MessageAddData(builder, dataOffset)
	return MessageEnd(builder)
}

func (rcv *Message) UnPackTo(t *MessageT) {
	dataTable := flatbuffers.Table{}
	if rcv.Data(&dataTable) {
		t.Data = rcv.DataType().UnPack(dataTable)
	}
}

func (rcv *Message) UnPack() *MessageT {
	if rcv == nil { return nil }
	t := &MessageT{}
	rcv.UnPackTo(t)
	return t
}

type Message struct {
	_tab flatbuffers.Table
}

func GetRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Message{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Message) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Message) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Message) DataType() GameMessage {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return GameMessage(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Message) MutateDataType(n GameMessage) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Message) Data(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageAddDataType(builder *flatbuffers.Builder, dataType GameMessage) {
	builder.PrependByteSlot(0, byte(dataType), 0)
}
func MessageAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(data), 0)
}
func MessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type MapUpdateT struct {
	Seq uint32
	Logins []*PlayerT
	Logouts []float64
	Psyncs []*PlayerT
}

func (t *MapUpdateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	loginsOffset := flatbuffers.UOffsetT(0)
	if t.Logins != nil {
		loginsLength := len(t.Logins)
		loginsOffsets := make([]flatbuffers.UOffsetT, loginsLength)
		for j := 0; j < loginsLength; j++ {
			loginsOffsets[j] = t.Logins[j].Pack(builder)
		}
		MapUpdateStartLoginsVector(builder, loginsLength)
		for j := loginsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(loginsOffsets[j])
		}
		loginsOffset = builder.EndVector(loginsLength)
	}
	logoutsOffset := flatbuffers.UOffsetT(0)
	if t.Logouts != nil {
		logoutsLength := len(t.Logouts)
		MapUpdateStartLogoutsVector(builder, logoutsLength)
		for j := logoutsLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Logouts[j])
		}
		logoutsOffset = builder.EndVector(logoutsLength)
	}
	psyncsOffset := flatbuffers.UOffsetT(0)
	if t.Psyncs != nil {
		psyncsLength := len(t.Psyncs)
		psyncsOffsets := make([]flatbuffers.UOffsetT, psyncsLength)
		for j := 0; j < psyncsLength; j++ {
			psyncsOffsets[j] = t.Psyncs[j].Pack(builder)
		}
		MapUpdateStartPsyncsVector(builder, psyncsLength)
		for j := psyncsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(psyncsOffsets[j])
		}
		psyncsOffset = builder.EndVector(psyncsLength)
	}
	MapUpdateStart(builder)
	MapUpdateAddSeq(builder, t.Seq)
	MapUpdateAddLogins(builder, loginsOffset)
	MapUpdateAddLogouts(builder, logoutsOffset)
	MapUpdateAddPsyncs(builder, psyncsOffset)
	return MapUpdateEnd(builder)
}

func (rcv *MapUpdate) UnPackTo(t *MapUpdateT) {
	t.Seq = rcv.Seq()
	loginsLength := rcv.LoginsLength()
	t.Logins = make([]*PlayerT, loginsLength)
	for j := 0; j < loginsLength; j++ {
		x := Player{}
		rcv.Logins(&x, j)
		t.Logins[j] = x.UnPack()
	}
	logoutsLength := rcv.LogoutsLength()
	t.Logouts = make([]float64, logoutsLength)
	for j := 0; j < logoutsLength; j++ {
		t.Logouts[j] = rcv.Logouts(j)
	}
	psyncsLength := rcv.PsyncsLength()
	t.Psyncs = make([]*PlayerT, psyncsLength)
	for j := 0; j < psyncsLength; j++ {
		x := Player{}
		rcv.Psyncs(&x, j)
		t.Psyncs[j] = x.UnPack()
	}
}

func (rcv *MapUpdate) UnPack() *MapUpdateT {
	if rcv == nil { return nil }
	t := &MapUpdateT{}
	rcv.UnPackTo(t)
	return t
}

type MapUpdate struct {
	_tab flatbuffers.Table
}

func GetRootAsMapUpdate(buf []byte, offset flatbuffers.UOffsetT) *MapUpdate {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MapUpdate{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MapUpdate) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MapUpdate) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MapUpdate) Seq() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MapUpdate) MutateSeq(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *MapUpdate) Logins(obj *Player, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MapUpdate) LoginsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MapUpdate) Logouts(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *MapUpdate) LogoutsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MapUpdate) MutateLogouts(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *MapUpdate) Psyncs(obj *Player, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MapUpdate) PsyncsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MapUpdateStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func MapUpdateAddSeq(builder *flatbuffers.Builder, seq uint32) {
	builder.PrependUint32Slot(0, seq, 0)
}
func MapUpdateAddLogins(builder *flatbuffers.Builder, logins flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(logins), 0)
}
func MapUpdateStartLoginsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MapUpdateAddLogouts(builder *flatbuffers.Builder, logouts flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(logouts), 0)
}
func MapUpdateStartLogoutsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func MapUpdateAddPsyncs(builder *flatbuffers.Builder, psyncs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(psyncs), 0)
}
func MapUpdateStartPsyncsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MapUpdateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type PlayerInputT struct {
	Seq uint32
	Actions []PlayerAction
}

func (t *PlayerInputT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	actionsOffset := flatbuffers.UOffsetT(0)
	if t.Actions != nil {
		actionsLength := len(t.Actions)
		PlayerInputStartActionsVector(builder, actionsLength)
		for j := actionsLength - 1; j >= 0; j-- {
			builder.PrependByte(byte(t.Actions[j]))
		}
		actionsOffset = builder.EndVector(actionsLength)
	}
	PlayerInputStart(builder)
	PlayerInputAddSeq(builder, t.Seq)
	PlayerInputAddActions(builder, actionsOffset)
	return PlayerInputEnd(builder)
}

func (rcv *PlayerInput) UnPackTo(t *PlayerInputT) {
	t.Seq = rcv.Seq()
	actionsLength := rcv.ActionsLength()
	t.Actions = make([]PlayerAction, actionsLength)
	for j := 0; j < actionsLength; j++ {
		t.Actions[j] = rcv.Actions(j)
	}
}

func (rcv *PlayerInput) UnPack() *PlayerInputT {
	if rcv == nil { return nil }
	t := &PlayerInputT{}
	rcv.UnPackTo(t)
	return t
}

type PlayerInput struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerInput(buf []byte, offset flatbuffers.UOffsetT) *PlayerInput {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerInput{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerInput) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerInput) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerInput) Seq() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInput) MutateSeq(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *PlayerInput) Actions(j int) PlayerAction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return PlayerAction(rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *PlayerInput) ActionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PlayerInput) ActionsBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PlayerInput) MutateActions(j int, n PlayerAction) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), byte(n))
	}
	return false
}

func PlayerInputStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PlayerInputAddSeq(builder *flatbuffers.Builder, seq uint32) {
	builder.PrependUint32Slot(0, seq, 0)
}
func PlayerInputAddActions(builder *flatbuffers.Builder, actions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(actions), 0)
}
func PlayerInputStartActionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PlayerInputEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
