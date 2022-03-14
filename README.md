# discord-chat-bot

Add discord notifications to your Kubernetes deployment.

This discord chat bot deploys a pod in your kubernetes cluster.
The pod will listen for any internal REST or gRPC requests, which will include a message payload.
This messsage payload will be sent to the desired discord channel, making it easy to add discord notifications to your cluster.
