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

type MessageType int8

const (
	MessageTypelogin  MessageType = 0
	MessageTypelogout MessageType = 1
	MessageTypesync   MessageType = 2
)

var EnumNamesMessageType = map[MessageType]string{
	MessageTypelogin:  "login",
	MessageTypelogout: "logout",
	MessageTypesync:   "sync",
}

var EnumValuesMessageType = map[string]MessageType{
	"login":  MessageTypelogin,
	"logout": MessageTypelogout,
	"sync":   MessageTypesync,
}

func (v MessageType) String() string {
	if s, ok := EnumNamesMessageType[v]; ok {
		return s
	}
	return "MessageType(" + strconv.FormatInt(int64(v), 10) + ")"
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
type PlayerT struct {
	Pos *Vec2T
	Uuid string
	Col Color
}

func (t *PlayerT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	uuidOffset := builder.CreateString(t.Uuid)
	PlayerStart(builder)
	posOffset := t.Pos.Pack(builder)
	PlayerAddPos(builder, posOffset)
	PlayerAddUuid(builder, uuidOffset)
	PlayerAddCol(builder, t.Col)
	return PlayerEnd(builder)
}

func (rcv *Player) UnPackTo(t *PlayerT) {
	t.Pos = rcv.Pos(nil).UnPack()
	t.Uuid = string(rcv.Uuid())
	t.Col = rcv.Col()
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

func (rcv *Player) Pos(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec2)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Player) Uuid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Player) Col() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Player) MutateCol(n Color) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func PlayerStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PlayerAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0)
}
func PlayerAddUuid(builder *flatbuffers.Builder, uuid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uuid), 0)
}
func PlayerAddCol(builder *flatbuffers.Builder, col Color) {
	builder.PrependInt8Slot(2, int8(col), 0)
}
func PlayerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type MessageT struct {
	Type MessageType
	Player *PlayerT
}

func (t *MessageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	playerOffset := t.Player.Pack(builder)
	MessageStart(builder)
	MessageAddType(builder, t.Type)
	MessageAddPlayer(builder, playerOffset)
	return MessageEnd(builder)
}

func (rcv *Message) UnPackTo(t *MessageT) {
	t.Type = rcv.Type()
	t.Player = rcv.Player(nil).UnPack()
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

func (rcv *Message) Type() MessageType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return MessageType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Message) MutateType(n MessageType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Message) Player(obj *Player) *Player {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Player)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageAddType(builder *flatbuffers.Builder, type_ MessageType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func MessageAddPlayer(builder *flatbuffers.Builder, player flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(player), 0)
}
func MessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
