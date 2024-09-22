package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"

	hello "github.com/stuttgart-things/cameo/hello" // Import the generated protobuf package

	"google.golang.org/grpc"
)

// gRPC server implementation
type server struct {
	hello.UnimplementedGreeterServer
}

// SayHello implements the gRPC SayHello method
func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// PageData is used to pass data into the template
type PageData struct {
	Title string
}

// renderForm renders the form page
func renderForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := PageData{Title: "HTMX Form Example with gRPC"}
	tmpl.Execute(w, data)
}

// handleSubmit handles the form submission
func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get form data
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Return a response to update the page dynamically
	response := fmt.Sprintf("<div><p><strong>Name:</strong> %s</p><p><strong>Email:</strong> %s</p></div>", name, email)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, response)
}

func main() {
	// HTTP server part
	http.HandleFunc("/", renderForm)
	http.HandleFunc("/submit", handleSubmit)

	// gRPC server part
	grpcServer := grpc.NewServer()
	hello.RegisterGreeterServer(grpcServer, &server{})

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Start HTTP server
	fmt.Println("HTTP server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
