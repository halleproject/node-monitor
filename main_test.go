package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"testing"
// 	"time"
// )
//
// func TestValidatorsByHeight(t *testing.T) {
//
// 	height := 79895
//
// 	cfg := NodeConfig{
// 		RPCNode:           "tcp://192.168.3.200:26657",
// 		APIServerEndpoint: "http://192.168.3.100:8545",
// 	}
//
// 	cli := NewClient(&cfg)
//
// 	valSet, err := cli.ValidatorSet(int64(height))
// 	if err != nil {
// 		t.Errorf("validatorset  error: %v \n", err.Error())
// 	}
//
// 	jsonValSetAsJson, err := json.MarshalIndent(valSet, "", "\t")
// 	if err != nil {
// 		t.Errorf("validator json marshal err:  %v \n", err.Error())
// 	}
//
// 	t.Logf("height:  %v  \n   %s \n", height, jsonValSetAsJson)
//
// }
//
// //Validators()
//
// func TestValidatorsOld(t *testing.T) {
//
// 	cfg := NodeConfig{
// 		RPCNode:           "tcp://192.168.3.200:26657",
// 		APIServerEndpoint: "http://192.168.3.100:8545",
// 	}
//
// 	cli := NewClient(&cfg)
//
// 	valSet, err := cli.Validators()
// 	if err != nil {
// 		t.Errorf("validatorset  error: %v \n", err.Error())
// 	}
//
// 	jsonValSetAsJson, err := json.MarshalIndent(valSet, "", "\t")
// 	if err != nil {
// 		t.Errorf("validator json marshal err:  %v \n", err.Error())
// 	}
//
// 	t.Logf("height: %s \n", jsonValSetAsJson)
//
// }
//
// func TestAlarm(t *testing.T) {
//
// 	ticker1 := time.NewTicker(1 * time.Second)
// 	i := 1
// 	for c := range ticker1.C {
// 		i++
// 		fmt.Println(c.Format("2006/01/02 15:04:05.999999999"))
// 		if i > 5 {
// 			ticker1.Stop()
// 			break
// 		}
// 	}
// 	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"), " 1 Finished.")
//
// 	i = 1
// 	ticker2 := time.AfterFunc(1*time.Second, func() {
// 		i++
// 		fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"))
// 	})
//
// 	for {
// 		select {
// 		case <-ticker2.C:
// 			fmt.Println("nsmei")
// 		case <-time.After(3 * time.Second):
// 			if i <= 5 {
// 				ticker2.Reset(1 * time.Second)
// 				continue
// 			}
// 			goto BRK
// 		}
// 	BRK:
// 		ticker2.Stop()
// 		break
// 	}
// 	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"), " 2 Finished.")
// }
