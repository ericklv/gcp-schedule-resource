# API for start/stop Cloud SQL instance
This API use gcloud in background to change status in cloud sql instances.

To run needs golang >= 1.22.5 and Gcloud SDK. Your account must have Cloud SQL Admin permissions.

To run in local use (default port is 5432): 
```bash
go run main.go
```
For docker, configure a google account service in GCP, then add Cloud SQL Admin permissions.
- Generate a key file as JSON [Ref](https://cloud.google.com/sdk/gcloud/reference/auth/activate-service-account). To generate a docker image update Dockerfile with this values.
- Generate image with (see logs if your account service has any permissions problem):
```bash
docker build -t {your_name}/{image_name} -f Dockerfile . --progress plain --no-cache
```

Make a cloud run o cloud function gen2 with docker image, call this service with Cloud Scheduler
 
Good luck, have fun.