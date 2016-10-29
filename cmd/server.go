package main

import (
       "log"
       "net"

       "golang.org/x/net/context"
       "google.golang.org/grpc"
       pb "github.com/perigee/terrant/pb"
       "github.com/hashicorp/terraform/config"
       "github.com/hashicorp/terraform/terraform"
)

const (
      port = ":50051"
)



type server struct {}



func (s *server) Destroy(ctx context.Context, in *pb.ApplyReq) (*pb.ApplyResp, error) {

     
     
     return &pb.ApplyResp{ProjectId: "Applying " + in.ProjectId}, nil
}


func (s *server) Apply(ctx context.Context, in *pb.ApplyReq) (*pb.ApplyResp, error) {

   
     
     return &pb.ApplyResp{ProjectId: "Applying " + in.ProjectId}, nil
}



func main() {
     lis, err := net.Listen("tcp", port)
     if err != nil {
     	log.Fatalf("failed to listen: %v", err)
     }

     s := grpc.NewServer()
     pb.RegisterProvisionServer(s, &server{})
     if err := s.Serve(lis); err != nil {
     	log.Fatalf("failed to serve: %v", err)
     }
}