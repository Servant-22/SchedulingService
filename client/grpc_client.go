package client

import (
	"context"
	"log"
	"time"

	pb "github.com/Servant-22/SchedulingService/pb"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	client pb.MaintenanceServiceClient
}

func NewGrpcClient(address string) *GrpcClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &GrpcClient{
		client: pb.NewMaintenanceServiceClient(conn),
	}
}

func (g *GrpcClient) ScheduleAppointment(userId, taskDescription, preferredTime string) (*pb.AppointmentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.AppointmentRequest{
		UserId:          userId,
		TaskDescription: taskDescription,
		PreferredTime:   preferredTime,
	}

	return g.client.ScheduleAppointment(ctx, request)
}
