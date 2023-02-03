When the app is deployed, it triggers a health check timeout by waiting 3s before responding when a request is made 10 minutes after startup. Looking at `cf events`, we can see the first crash event when the app crashes the first time. All subsequent crashes don’t show up (even though the app container is killed/restarted because it got unhealthy).

The app logs show subsequent crashes but output of cf events and event graph in Apps Manager UI does not reflect these crashes.
