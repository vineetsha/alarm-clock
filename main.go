package main

import ()

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	//// Listen for SIGINT and SIGTERM signals and gracefully shutdown the server
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()
	//executors.InitThreadPool(10, 2)
	//mux := controllers.Register()
	//models.ConnectDB()
	//server := &http.Server{
	//	Addr:    "127.0.0.1:8080",
	//	Handler: mux,
	//}
	//
	//// Make server a goroutine so that it doesn't block the graceful shutdown handling below
	//go server.ListenAndServe()
	//
	//// Wait for SIGINT (Ctrl+C) or SIGTERM signal
	//<-ctx.Done()
	//
	//// CReate shutdown context as notification context got cancelled
	//shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer shutdownCancel()
	//if err := server.Shutdown(shutdownCtx); err != nil {
	//	fmt.Println("Graceful Shutdown error:", err)
	//	panic(err)
	//}
	//
	//// Wait for all in-flight requests to complete
	//controllers.CloseInFlightRequests()
	//
	//// Close DB connection
	//models.ShutDownDB()
	//
	//// Shutdown thread pool
	//executors.ShutdownThreadPool()
	//
	//os.Exit(0)

}
