package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/superoo7/statbot2/proto"
	"google.golang.org/grpc"
)

type CoinChart struct {
	FileName  string
	Timestamp int64
	Key       string
}

func GetChart(coin string) *CoinChart {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewCoinPriceChartClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.PriceChart(ctx, &pb.CoinInfo{Coin: coin})
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}
	log.Printf("[%d] %s %s", r.Timestamp, r.FileName, r.Key)
	return &CoinChart{
		Timestamp: r.Timestamp,
		FileName:  r.FileName,
		Key:       r.Key,
	}
}

func GetDailyChart() *CoinChart {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewCoinPriceChartClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	r, err := c.DailyChart(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}
	log.Printf("[%d] %s %s", r.Timestamp, r.FileName, r.Key)
	return &CoinChart{
		Timestamp: r.Timestamp,
		FileName:  r.FileName,
		Key:       r.Key,
	}
}
