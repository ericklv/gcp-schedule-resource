# API for start/stop/restart/resize Cloud SQL instance
This API allows to execute GCloud SDK commands (gcloud) to change the status of a Cloud SQL instance.

Using goroutines the command is executed in the background, some commands can take up to 5min to complete.

---
**NOTE**

If you want to use the API synchronously you can use `api-sync` branch

---

# Requirements
To run needs golang >= 1.22 and Gcloud SDK. Your account must have Cloud SQL Admin permissions.

To run in local use (default port is 5432): 
```bash
go run main.go
```

# API
- `/health` : check if API is enabled.
- `/:action/:inst_name`: Has 2 params
  - action: can be "start", "stop" , "restart"
  - inst_name: cloud sql instance name.
- `/:action/:inst_name/:resize`: Has 3 params.
  - action: set "resize"
  - inst_name: cloud sql instance name. 
  - resize: can be "up" or "down".
  - Configure `resize.json` with valid machine-types for your instance. 
  `up` for work hours and `down` for non work hours. Example:
  ```json
    {
    "machines": [
        {
            "name": "awesome-machine",
            "up": "db-n1-standard-32",
            "down": "db-n1-standard-2"
        }
    ]}
  ```
  To see list valid machines types and specs run:
  ```bash
    gcloud alpha sql tiers list
  ```

# Create a Cloud Run
1. Configure a google account service in GCP and add Cloud SQL Admin permissions.
2. Generate a key file as JSON [Ref](https://cloud.google.com/sdk/gcloud/reference/auth/activate-service-account). 
3. Update Dockerfile with this values.  
   - `key-file` is your JSON  
   - `project` is project-id of instances.
   - `activate-service-account` is your google account service.
4. Generate image with (see logs if your account service has any permissions problem):
```bash
docker build -t {your_name}/{image_name} -f Dockerfile . --progress plain --no-cache
```
5. Upload image to Docker Hub or Artifact Registry in GCP
6. Make a cloud run with docker image o create a cloudbuild to generate image and deploy. 
  - Remmember enable `CPU is always allocated` in cloud run configuration if use `main` branch (for goroutines works correctly). 
7. Use Cloud Scheduler to consume the service.
 
Good luck, have fun.