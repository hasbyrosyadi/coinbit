package usecase

import (
	pb "coinbit-test/gen/proto"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"gopkg.in/redis.v5"
)

type DepositServiceServer struct {
	pb.UnimplementedDepositServiceServer
}

const Deposit = "deposit-wallet "

type StoreData struct {
	Amount    float32
	CreatedAt time.Time
}

func (s *DepositServiceServer) Deposit(ctx context.Context, req *pb.RequestDeposit) (*pb.ResponsePostDeposit, error) {

	// psudocode
	// 1. get input from req.
	// 2. publish to kafka for set into redis with key walletID and value is amount, can use viewTable by goka.
	//    but I still learn how to use that and implement it.
	// 3. return success if not found error.

	strWallet := strconv.Itoa(int(req.Request.WalletID))

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	var storeData StoreData

	storeData.Amount = req.Request.Amount
	storeData.CreatedAt = time.Now()
	data, _ := json.Marshal(storeData)
	err := client.RPush(Deposit+strWallet, data).Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &pb.ResponsePostDeposit{
		Result: "success",
	}, nil
}

func (s *DepositServiceServer) GetDeposit(ctx context.Context, req *pb.DepositRequest) (*pb.ResponseGetDeposit, error) {

	// psudocode
	// 1. get input from req
	// 2. get value from redis with key walletID
	// 3. get time from all index
	// 4. calculate range date and compare within a single 2 minutes window

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	strWallet := strconv.Itoa(int(req.WalletID))

	walletAmount, err := client.LRange(Deposit+strWallet, 0, math.MaxInt64).Result()
	if err != nil {
		log.Fatal(err)
	}
	var storeData []StoreData

	// insert into struct model
	for i := 0; i < len(walletAmount); i++ {
		data := StoreData{}
		json.Unmarshal([]byte(walletAmount[i]), &data)

		storeData = append(storeData, data)
	}

	totalBalance := float32(0.0)
	rollingPeriod := float64(120) // seconds
	var initDeposit, lastDeposit time.Time

	for i := 0; i < len(storeData); i++ {
		fmt.Println(storeData[i])
		totalBalance += storeData[i].Amount
		initDeposit = storeData[0].CreatedAt
		lastDeposit = storeData[len(storeData)-1].CreatedAt
	}

	var threshold bool
	rangeSecond := lastDeposit.Sub(initDeposit).Abs().Seconds()
	if (rollingPeriod > rangeSecond) && (totalBalance >= 10000) {
		threshold = true
	} else {
		threshold = false
	}

	return &pb.ResponseGetDeposit{
		WalletID:       req.WalletID,
		Balance:        totalBalance,
		AboveThreshold: threshold,
	}, nil
}
