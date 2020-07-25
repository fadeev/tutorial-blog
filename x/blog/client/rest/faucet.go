package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/tendermint/tendermint/libs/bech32"
)

type claimReq struct {
	Address string
}

func faucetHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var claim claimReq
		decoder := json.NewDecoder(r.Body)
		decoderErr := decoder.Decode(&claim)
		if decoderErr != nil {
			log.Println(decoderErr)
		}
		// make sure address is bech32
		readableAddress, decodedAddress, decodeErr := bech32.DecodeAndConvert(claim.Address)
		if decodeErr != nil {
			log.Println(decodeErr)
		}
		// re-encode the address in bech32
		encodedAddress, encodeErr := bech32.ConvertAndEncode(readableAddress, decodedAddress)
		if encodeErr != nil {
			log.Println(encodeErr)
		}
		cmd := exec.Command("blogcli", "tx", "send", "me", encodedAddress, "1token", "-y")
		_, err := cmd.Output()
		if err != nil {
			log.Println(fmt.Sprintf("%s", err))
		}
		rest.PostProcessResponse(w, cliCtx, encodedAddress)
	}
}
