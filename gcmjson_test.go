package gcmjson

import (
	"encoding/json"
	"math/rand"
	"testing"
)

const (
	HEX_LETTERS_DIGITS = "abcdef0123456789"

	GCM_RESPONSE_MOCK = `{
	"multicast_id": 7881985661182975128,
	"success": 1,"failure": 0,"canonical_ids": 0,
	"results": [
		{
			"message_id": "0:14691106133768229805a54495"
		}
	]
}`
)

func genRandomString(keylen int) string {
	l := len(HEX_LETTERS_DIGITS)
	b := make([]byte, keylen)
	for i := 0; i < keylen; i++ {
		b[i] = HEX_LETTERS_DIGITS[rand.Intn(l)]
	}
	return string(b)
}

/* Encoding, standard json library */

func BenchmarkEncodeJson(b *testing.B) {
	to := genRandomString(160)
	sub_1 := genRandomString(40)
	sub_2 := genRandomString(40)
	sub_3 := genRandomString(40)

	for i := 0; i < b.N; i++ {
		data := GcmDataSubList{SubList: make([]GcmDataSubItem, 0, 0)}
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_1, ChangeTs: 1})
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_2, ChangeTs: 2})
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_3, ChangeTs: 3})

		packet := GcmPacket{To: to, CollapseKey: "cs", Priority: "high", RestrictedPackage: "org.kman.AquaMail"}

		buf, err := json.Marshal(&packet)

		if buf == nil || err != nil {
			b.Fatal("Error encoding json: %s, %s", buf, err)
		}
	}
}

/* Encoding, EasyJSON */

func BenchmarkEncodeEasyJson(b *testing.B) {
	to := genRandomString(160)
	sub_1 := genRandomString(40)
	sub_2 := genRandomString(40)
	sub_3 := genRandomString(40)

	for i := 0; i < b.N; i++ {
		data := GcmDataSubList{SubList: make([]GcmDataSubItem, 0, 0)}
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_1, ChangeTs: 1})
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_2, ChangeTs: 2})
		data.SubList = append(data.SubList, GcmDataSubItem{SubId: sub_3, ChangeTs: 3})

		packet := GcmPacket{To: to, CollapseKey: "cs", Priority: "high", RestrictedPackage: "org.kman.AquaMail"}

		buf, err := packet.MarshalJSON()

		if buf == nil || err != nil {
			b.Fatal("Error encoding json: %s, %s", buf, err)
		}
	}
}

/* Decoding, standard json library */

func BenchmarkDecodeJson(b *testing.B) {

	responseBytes := []byte(GCM_RESPONSE_MOCK)

	for i := 0; i < b.N; i++ {
		var resp *GcmResponse
		err := json.Unmarshal(responseBytes, &resp)

		if resp == nil || err != nil {
			b.Fatal("Error decoding json: %s, %s", resp, err)
		}
	}
}

/* Decoding, EasyJSON */

func BenchmarkDecodeEasyJson(b *testing.B) {

	responseBytes := []byte(GCM_RESPONSE_MOCK)

	for i := 0; i < b.N; i++ {
		var resp GcmResponse
		err := resp.UnmarshalJSON(responseBytes)

		if err != nil {
			b.Fatal("Error decoding json: %s, %s", resp, err)
		}
	}
}
