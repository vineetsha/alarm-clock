package executors

import "sync"

type Job func()

func Worker(jobs <-chan Job, wg *sync.WaitGroup) {
	for job := range jobs {
		// Run the job, runnable only
		job()
		wg.Done()
	}
}

var threadPoolInstance *ThreadPool
var isInitialized bool = false

var once sync.Once

type ThreadPool struct {
	Jobs           chan Job
	threadpoolSize int
	workersWg      sync.WaitGroup
	workers        int
}

func InitThreadPool(threadpoolSize int, workers int) {
	once.Do(func() {
		threadPoolInstance = &ThreadPool{
			Jobs:           make(chan Job, threadpoolSize),
			threadpoolSize: threadpoolSize,
			workers:        workers,
			workersWg:      sync.WaitGroup{},
		}
		for i := 0; i < workers; i++ {
			threadPoolInstance.workersWg.Add(1)
			go Worker(threadPoolInstance.Jobs, &threadPoolInstance.workersWg)
		}
		isInitialized = true
	})
}

func SubmitJob(job Job) {
	if isInitialized {
		threadPoolInstance.workersWg.Add(1)
		threadPoolInstance.Jobs <- job
	} else {
		panic("Thread pool not initialized")
	}

}

func ShutdownThreadPool() {
	if isInitialized {
		close(threadPoolInstance.Jobs)
		// Wait for already running jobs
		threadPoolInstance.workersWg.Wait()
	} else {
		panic("Thread pool not initialized")
	}

}
