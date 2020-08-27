package main

import (
	"testing"
)

func BenchmarkStandartMarshal(b *testing.B) {
	for i:=0; i<b.N; i++ {
		U := InitUsersStruct(0)
		_, _ = MarshalStandart(U)
	}
}

func BenchmarkStandartUnmarshal(b *testing.B) {
	for i:=0; i<b.N; i++ {
		xmlSrc := []byte("<USERS_PACK><id_organization>68</id_organization></USERS_PACK>")
		_, _ = UnmarshalStandart(xmlSrc)
	}
}

func BenchmarkStreamMarshal(b *testing.B) {
	for i:=0; i<b.N; i++ {
		U := InitUsersStruct(0)
		_, _ = MarshalStream(U)
	}
}

func BenchmarkStreamUnmarshal(b *testing.B) {
	for i:=0; i<b.N; i++ {
		xmlSrc := []byte("<USERS_PACK><id_organization>68</id_organization></USERS_PACK>")
		_, _ = UnmarshalStream(xmlSrc)
	}
}

func BenchmarkLibxml2Marshal(b *testing.B) {
	for i:=0; i<b.N; i++ {
		U := InitUsersStruct(0)
		_, _ = MarshalLibXML2_tester(U)
	}
}
