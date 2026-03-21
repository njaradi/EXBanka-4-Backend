package main

import (
	"log"
	"net"

	loandb "github.com/RAF-SI-2025/EXBanka-4-Backend/services/loan-service/db"
	"github.com/RAF-SI-2025/EXBanka-4-Backend/services/loan-service/handlers"
	pb "github.com/RAF-SI-2025/EXBanka-4-Backend/shared/pb/loan"
	"google.golang.org/grpc"
)

const (
	loanDBDSN     = "postgres://loan_user:loan_pass@localhost:5439/loan_db?sslmode=disable"
	accountDBDSN  = "postgres://account_user:account_pass@localhost:5436/account_db?sslmode=disable"
	exchangeDBDSN = "postgres://exchange_user:exchange_pass@localhost:5438/exchange_db?sslmode=disable"
	grpcPort      = ":50058"
)

func main() {
	loanDB, err := loandb.Connect(loanDBDSN)
	if err != nil {
		log.Fatalf("failed to connect to loan_db: %v", err)
	}
	defer loanDB.Close()

	accountDB, err := loandb.Connect(accountDBDSN)
	if err != nil {
		log.Fatalf("failed to connect to account_db: %v", err)
	}
	defer accountDB.Close()

	exchangeDB, err := loandb.Connect(exchangeDBDSN)
	if err != nil {
		log.Fatalf("failed to connect to exchange_db: %v", err)
	}
	defer exchangeDB.Close()

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcPort, err)
	}

	srv := grpc.NewServer()
	loanServer := &handlers.LoanServer{
		DB:         loanDB,
		AccountDB:  accountDB,
		ExchangeDB: exchangeDB,
	}
	pb.RegisterLoanServiceServer(srv, loanServer)

	// Start cron jobs
	loanServer.StartCronJobs()

	log.Printf("loan-service gRPC server listening on %s", grpcPort)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}
