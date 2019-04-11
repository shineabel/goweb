package main

import (
	"github.com/tendermint/go-amino"
	"fmt"
)

func main() {

	var cdc = amino.NewCodec()
	cdc.RegisterInterface((*Message)(nil),nil)
	cdc.RegisterConcrete((&bcMessage{}),"bcMessage",nil)
	cdc.RegisterConcrete((&bcResponse{}),"bcResponse",nil)
	cdc.RegisterConcrete((&bcStatus{}),"bcStatus",nil)


	var bm = &bcMessage{
		Height:100,
		Message:"ABC",
	}
	var msg = bm

	bz , _ := cdc.MarshalBinaryLengthPrefixed(msg)


	fmt.Printf("Encode:%X",bz)

	var msg2 bcMessage
	cdc.UnmarshalBinaryLengthPrefixed(bz,&msg2)

	fmt.Println(msg2)


}

type Message interface {

}

type bcMessage struct {
	Height int
	Message string
}

type bcResponse struct {
	Status int
	Message string
}

type bcStatus struct {
	Peers int
}