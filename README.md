# worker-pool

A video-encoding pipeline in Go built on the classic **dispatcher / worker-pool
concurrency pattern** — bounded parallelism over a stream of jobs using nothing but
goroutines and channels.

## How it works

```
jobQueue (chan VideoProcessingJob)
        │
   VideoDispatcher.Run()
        │  spawns N workers, each registers its own
        │  job channel into the shared WorkerPool
        ▼
WorkerPool (chan chan VideoProcessingJob)
        │  dispatcher pulls a free worker's channel,
        │  hands it the next job
        ▼
videoWorker.processVideoJob()
        │  encodes via the pluggable Processor (ffmpeg-backed encoders)
        ▼
NotifyChan (per-job ProcessingMessage: success/failure + output path)
```

- **`streamer/pool.go`** — dispatcher + worker lifecycle (`WorkerPool chan chan Job` pattern)
- **`streamer/streamer.go`** — job/video/result types, pluggable `Processor` interface
- **`streamer/encoders.go`** — concrete encoders (MP4 and friends)
- **`app/main.go`** — example driver: queues videos (including a bad input to show failure handling) and collects results from the notify channel

## Why this pattern

Encoding is CPU/IO-heavy; unbounded goroutines would thrash. The pool gives:

- fixed worker count (bounded resource usage)
- backpressure via the job queue
- per-job result delivery without shared state — results flow back on each job's own channel

## Run it

```bash
cd app
go run .
```

Requires ffmpeg on PATH for the real encoders.
