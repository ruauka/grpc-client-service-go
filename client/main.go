package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"

	"client/internal/request"
	"client/pb"
)

// Execute - основная функция скрипта.
func Execute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(req)

	conn, err := grpc.Dial(
		":8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewExecutorClient(conn)

	payload := FmtToGRPC(req)

	resp, err := c.Execute(context.Background(), payload)
	if err != nil {
		errStatus, _ := status.FromError(err)
		log.Println(errStatus.Message())
		log.Println(errStatus.Code())
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck
}

func FmtToGRPC(req *request.Request) *pb.Request {
	var accountsSl []*pb.Account
	for _, v := range req.Accounts {
		account := &pb.Account{PaymentDateBalance: *v.PaymentDateBalance}
		accountsSl = append(accountsSl, account)
	}

	var loansSl []*pb.Loan
	for _, v := range req.Loans {
		loan := &pb.Loan{
			Payment:           v.Payment,
			PaymentDate:       timestamppb.New(v.PaymentDate),
			ActualPaymentDate: timestamppb.New(v.ActualPaymentDate),
		}
		loansSl = append(loansSl, loan)
	}

	return &pb.Request{
		Name:    req.Name,
		Surname: req.Surname,
		Account: accountsSl,
		Loan:    loansSl,
	}
}

func main() {
	router := httprouter.New()
	router.POST("/execute", Execute)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
