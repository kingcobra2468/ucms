# **UCMS**
UCMS (universal client message service), is a microservice that handles 
the propogation of client notifications onto a topic on FCM.

## **Config**

### **Firebase**
To setup UCMS, one must first setup FCM service via Firebase. Create a new Firebase project,
and download the service account file as needed by the firebase andmin SDK. Then, the path
to this JSON file must be passed to the **GOOGLE_APPLICATION_CREDENTIALS** environment variable.


### **Environemnt Setup**
The following environment variables are configurable when launching UCMS:
- **UCMS_HOSTNAME (string)=** hostname for ucms (default "127.0.0.1")
- **UCMS_PORT (int)=** port for ucms (default 8080)
- **UCMS_FCM_TOPIC (string)=** fcm topic for registration token subscription (default "un")

## **Topic**
UCMS functions by attaching all devices to topic specified by the **--topic** flag. By default,
the topic will be "un" which stands for universal notification. Regardless, when sending
notifications, ensure that notifications are being sent to the correct topic.


## **Endpoints**
The following REST endpoints are available:
- `/notification` **[POST]** Push an event as a notification event to the FCM topic. Request schema:
  ```json
  {
      "title": "title of the notification (string)",
      "body": "body of the notification (string)",
      "service_name": "name of the client which is sending this notification (string)",
      "image": "optional field for image url of an image within the notification. (string)",
  }
  ```

## Setup
1. Install Golang(1.16) onto the machine.
2. Setup Firebase and Redis as specified [here](#config).
3. Build the application with `go build`.
4. Launch the application with the appropriate flags via the `ucms` binary.
