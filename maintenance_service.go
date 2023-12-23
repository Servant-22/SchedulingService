package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Servant-22/SchedulingService/pb"
	"google.golang.org/grpc"
)

type maintenanceServer struct {
	pb.UnimplementedMaintenanceServiceServer
	db *sql.DB
}

func newServer(db *sql.DB) *maintenanceServer {
	return &maintenanceServer{
		db: db,
	}
}

func (s *maintenanceServer) ScheduleAppointment(ctx context.Context, req *pb.AppointmentRequest) (*pb.AppointmentResponse, error) {
	// Controleer of de gebruiker al een geplande afspraak heeft
	var existingAppointmentCount int
	err := s.db.QueryRow("SELECT COUNT(*) FROM appointments WHERE user_id = $1", req.UserId).Scan(&existingAppointmentCount)
	if err != nil {
		return nil, fmt.Errorf("database query error: %v", err)
	}
	if existingAppointmentCount > 0 {
		return nil, fmt.Errorf("gebruiker %s heeft al een geplande afspraak", req.UserId)
	}

	// Plan de afspraak in
	preferredTime, _ := time.Parse(time.RFC3339, req.PreferredTime)
	confirmationID := fmt.Sprintf("CONF-%d", time.Now().Unix()) // Eenvoudige generatie van een bevestigings-ID

	_, err = s.db.Exec("INSERT INTO appointments (user_id, task_description, preferred_time, confirmation_id) VALUES ($1, $2, $3, $4)",
		req.UserId, req.TaskDescription, preferredTime, confirmationID)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule appointment: %v", err)
	}

	return &pb.AppointmentResponse{
		ConfirmationId: confirmationID,
		ScheduledTime:  req.PreferredTime,
	}, nil
}

func main() {
	db := connectDB()
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMaintenanceServiceServer(grpcServer, newServer(db))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
