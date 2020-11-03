// +build  opencontrail

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package opencontrail

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

func easyjson23c2e8cdDecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(in *jlexer.Lexer, out *rtMonitorRoute) {
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
		case "Operation":
			out.Operation = string(in.String())
		case "Family":
			out.Family = string(in.String())
		case "vrf_id":
			out.VrfId = int(in.Int())
		case "Prefix":
			out.Prefix = int(in.Int())
		case "Address":
			out.Address = string(in.String())
		case "nh_id":
			out.NhId = int(in.Int())
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
func easyjson23c2e8cdEncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(out *jwriter.Writer, in rtMonitorRoute) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Operation\":"
		out.RawString(prefix[1:])
		out.String(string(in.Operation))
	}
	{
		const prefix string = ",\"Family\":"
		out.RawString(prefix)
		out.String(string(in.Family))
	}
	{
		const prefix string = ",\"vrf_id\":"
		out.RawString(prefix)
		out.Int(int(in.VrfId))
	}
	{
		const prefix string = ",\"Prefix\":"
		out.RawString(prefix)
		out.Int(int(in.Prefix))
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"nh_id\":"
		out.RawString(prefix)
		out.Int(int(in.NhId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v rtMonitorRoute) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson23c2e8cdEncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v rtMonitorRoute) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson23c2e8cdEncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *rtMonitorRoute) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson23c2e8cdDecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *rtMonitorRoute) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson23c2e8cdDecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(l, v)
}
