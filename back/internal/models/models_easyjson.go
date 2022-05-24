// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeHurrlesInternalModels(in *jlexer.Lexer, out *UserCredentials) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels(out *jwriter.Writer, in UserCredentials) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Email != "" {
		const prefix string = ",\"email\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCredentials) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCredentials) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCredentials) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCredentials) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels1(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
		case "number":
			out.Number = string(in.String())
		case "isAdmin":
			out.IsAdmin = bool(in.Bool())
		case "isRestaurant":
			out.IsRestaurant = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if in.FullName != "" {
		const prefix string = ",\"fullName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FullName))
	}
	if in.Number != "" {
		const prefix string = ",\"number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Number))
	}
	if in.IsAdmin {
		const prefix string = ",\"isAdmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsAdmin))
	}
	if in.IsRestaurant {
		const prefix string = ",\"isRestaurant\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsRestaurant))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels1(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels2(in *jlexer.Lexer, out *Session) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Cookie":
			out.Cookie = string(in.String())
		case "UserID":
			out.UserID = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels2(out *jwriter.Writer, in Session) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Cookie\":"
		out.RawString(prefix[1:])
		out.String(string(in.Cookie))
	}
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Session) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Session) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Session) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Session) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels2(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels3(in *jlexer.Lexer, out *RestaurantList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RestaurantList, 0, 0)
			} else {
				*out = RestaurantList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Restaurant
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels3(out *jwriter.Writer, in RestaurantList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels3(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels4(in *jlexer.Lexer, out *Restaurant) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "metro":
			out.Metro = string(in.String())
		case "number":
			out.Number = string(in.String())
		case "openTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OpenTime).UnmarshalJSON(data))
			}
		case "closeTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CloseTime).UnmarshalJSON(data))
			}
		case "kitchen":
			out.Kitchen = string(in.String())
		case "img":
			out.Img = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels4(out *jwriter.Writer, in Restaurant) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Address != "" {
		const prefix string = ",\"address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	if in.Metro != "" {
		const prefix string = ",\"metro\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Metro))
	}
	if in.Number != "" {
		const prefix string = ",\"number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Number))
	}
	if true {
		const prefix string = ",\"openTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.OpenTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"closeTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CloseTime).MarshalJSON())
	}
	if in.Kitchen != "" {
		const prefix string = ",\"kitchen\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kitchen))
	}
	if in.Img != "" {
		const prefix string = ",\"img\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Img))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Restaurant) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Restaurant) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Restaurant) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Restaurant) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels4(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels5(in *jlexer.Lexer, out *PlaceParameters) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "restaurantId":
			out.RestaurantId = uint64(in.Uint64())
		case "time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		case "floor":
			out.Floor = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels5(out *jwriter.Writer, in PlaceParameters) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RestaurantId != 0 {
		const prefix string = ",\"restaurantId\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.RestaurantId))
	}
	if true {
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Time).MarshalJSON())
	}
	if in.Floor != 0 {
		const prefix string = ",\"floor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Floor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlaceParameters) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlaceParameters) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlaceParameters) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlaceParameters) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels5(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels6(in *jlexer.Lexer, out *PlaceList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PlaceList, 0, 0)
			} else {
				*out = PlaceList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 Place
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels6(out *jwriter.Writer, in PlaceList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PlaceList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlaceList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlaceList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlaceList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels6(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels7(in *jlexer.Lexer, out *Place) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "restaurantId":
			out.RestaurantId = uint64(in.Uint64())
		case "capacity":
			out.Capacity = int(in.Int())
		case "number":
			out.Number = int(in.Int())
		case "leftTop":
			out.LeftTop = int(in.Int())
		case "rightBottom":
			out.RightBottom = int(in.Int())
		case "floor":
			out.Floor = int(in.Int())
		case "width":
			out.Width = int(in.Int())
		case "height":
			out.Height = int(in.Int())
		case "isBooked":
			out.IsBooked = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels7(out *jwriter.Writer, in Place) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.RestaurantId != 0 {
		const prefix string = ",\"restaurantId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.RestaurantId))
	}
	if in.Capacity != 0 {
		const prefix string = ",\"capacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Capacity))
	}
	if in.Number != 0 {
		const prefix string = ",\"number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Number))
	}
	{
		const prefix string = ",\"leftTop\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.LeftTop))
	}
	{
		const prefix string = ",\"rightBottom\":"
		out.RawString(prefix)
		out.Int(int(in.RightBottom))
	}
	if in.Floor != 0 {
		const prefix string = ",\"floor\":"
		out.RawString(prefix)
		out.Int(int(in.Floor))
	}
	if in.Width != 0 {
		const prefix string = ",\"width\":"
		out.RawString(prefix)
		out.Int(int(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Int(int(in.Height))
	}
	{
		const prefix string = ",\"isBooked\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsBooked))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Place) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Place) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Place) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Place) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels7(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels8(in *jlexer.Lexer, out *Payment) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "orderId":
			out.OrderId = uint64(in.Uint64())
		case "totalCost":
			out.TotalCost = int32(in.Int32())
		case "format":
			out.Format = string(in.String())
		case "status":
			out.Status = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels8(out *jwriter.Writer, in Payment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.OrderId != 0 {
		const prefix string = ",\"orderId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.OrderId))
	}
	if in.TotalCost != 0 {
		const prefix string = ",\"totalCost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TotalCost))
	}
	if in.Format != "" {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Format))
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Payment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Payment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Payment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Payment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels8(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels9(in *jlexer.Lexer, out *OrderList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(OrderList, 0, 0)
			} else {
				*out = OrderList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 Order
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels9(out *jwriter.Writer, in OrderList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v OrderList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels9(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels10(in *jlexer.Lexer, out *Order) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "userId":
			out.UserId = uint64(in.Uint64())
		case "placeId":
			out.PlaceId = uint64(in.Uint64())
		case "restaurantTitle":
			out.RestaurantTitle = string(in.String())
		case "restaurantAddress":
			out.RestaurantAddress = string(in.String())
		case "restaurantMetro":
			out.RestaurantMetro = string(in.String())
		case "placeNumber":
			out.PlaceNumber = int32(in.Int32())
		case "placeCapacity":
			out.PlaceCapacity = int32(in.Int32())
		case "dishesIds":
			if in.IsNull() {
				in.Skip()
				out.DishesIds = nil
			} else {
				in.Delim('[')
				if out.DishesIds == nil {
					if !in.IsDelim(']') {
						out.DishesIds = make([]int32, 0, 16)
					} else {
						out.DishesIds = []int32{}
					}
				} else {
					out.DishesIds = (out.DishesIds)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int32
					v10 = int32(in.Int32())
					out.DishesIds = append(out.DishesIds, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dishesTitles":
			if in.IsNull() {
				in.Skip()
				out.DishesTitles = nil
			} else {
				in.Delim('[')
				if out.DishesTitles == nil {
					if !in.IsDelim(']') {
						out.DishesTitles = make([]string, 0, 4)
					} else {
						out.DishesTitles = []string{}
					}
				} else {
					out.DishesTitles = (out.DishesTitles)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.DishesTitles = append(out.DishesTitles, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dishesPrices":
			if in.IsNull() {
				in.Skip()
				out.DishesPrices = nil
			} else {
				in.Delim('[')
				if out.DishesPrices == nil {
					if !in.IsDelim(']') {
						out.DishesPrices = make([]int32, 0, 16)
					} else {
						out.DishesPrices = []int32{}
					}
				} else {
					out.DishesPrices = (out.DishesPrices)[:0]
				}
				for !in.IsDelim(']') {
					var v12 int32
					v12 = int32(in.Int32())
					out.DishesPrices = append(out.DishesPrices, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dishesCounts":
			if in.IsNull() {
				in.Skip()
				out.DishesCounts = nil
			} else {
				in.Delim('[')
				if out.DishesCounts == nil {
					if !in.IsDelim(']') {
						out.DishesCounts = make([]int32, 0, 16)
					} else {
						out.DishesCounts = []int32{}
					}
				} else {
					out.DishesCounts = (out.DishesCounts)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int32
					v13 = int32(in.Int32())
					out.DishesCounts = append(out.DishesCounts, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "startTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "endTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndTime).UnmarshalJSON(data))
			}
		case "cost":
			out.Cost = int32(in.Int32())
		case "createdTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedTime).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels10(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.UserId != 0 {
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.UserId))
	}
	if in.PlaceId != 0 {
		const prefix string = ",\"placeId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.PlaceId))
	}
	if in.RestaurantTitle != "" {
		const prefix string = ",\"restaurantTitle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RestaurantTitle))
	}
	if in.RestaurantAddress != "" {
		const prefix string = ",\"restaurantAddress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RestaurantAddress))
	}
	if in.RestaurantMetro != "" {
		const prefix string = ",\"restaurantMetro\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RestaurantMetro))
	}
	if in.PlaceNumber != 0 {
		const prefix string = ",\"placeNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlaceNumber))
	}
	if in.PlaceCapacity != 0 {
		const prefix string = ",\"placeCapacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlaceCapacity))
	}
	if len(in.DishesIds) != 0 {
		const prefix string = ",\"dishesIds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.DishesIds {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v15))
			}
			out.RawByte(']')
		}
	}
	if len(in.DishesTitles) != 0 {
		const prefix string = ",\"dishesTitles\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.DishesTitles {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.DishesPrices) != 0 {
		const prefix string = ",\"dishesPrices\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.DishesPrices {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.DishesCounts) != 0 {
		const prefix string = ",\"dishesCounts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.DishesCounts {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v21))
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"startTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.StartTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"endTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.EndTime).MarshalJSON())
	}
	{
		const prefix string = ",\"cost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Cost))
	}
	if true {
		const prefix string = ",\"createdTime\":"
		out.RawString(prefix)
		out.Raw((in.CreatedTime).MarshalJSON())
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels10(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels11(in *jlexer.Lexer, out *DishOrder) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "number":
			out.Number = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels11(out *jwriter.Writer, in DishOrder) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Number))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DishOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DishOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DishOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DishOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels11(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels12(in *jlexer.Lexer, out *DishList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DishList, 0, 1)
			} else {
				*out = DishList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v22 Dish
			(v22).UnmarshalEasyJSON(in)
			*out = append(*out, v22)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels12(out *jwriter.Writer, in DishList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v23, v24 := range in {
			if v23 > 0 {
				out.RawByte(',')
			}
			(v24).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v DishList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DishList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DishList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DishList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels12(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels13(in *jlexer.Lexer, out *Dish) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint64(in.Uint64())
		case "restaurantId":
			out.RestaurantId = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "price":
			out.Price = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels13(out *jwriter.Writer, in Dish) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.RestaurantId != 0 {
		const prefix string = ",\"restaurantId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.RestaurantId))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Price))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Dish) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Dish) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Dish) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Dish) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels13(l, v)
}
func easyjsonD2b7633eDecodeHurrlesInternalModels14(in *jlexer.Lexer, out *Base) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "isError":
			out.IsError = bool(in.Bool())
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeHurrlesInternalModels14(out *jwriter.Writer, in Base) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsError {
		const prefix string = ",\"isError\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsError))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Base) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHurrlesInternalModels14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Base) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHurrlesInternalModels14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Base) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHurrlesInternalModels14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Base) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHurrlesInternalModels14(l, v)
}
