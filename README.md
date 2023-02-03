When the app is deployed, it triggers a health check timeout by waiting 3s before responding when a request is made 10 minutes after startup. Looking at `cf events`, we can see the first crash event when the app crashes the first time. All subsequent crashes don’t show up (even though the app container is killed/restarted because it got unhealthy).

The app logs show subsequent crashes but output of cf events and event graph in Apps Manager UI does not reflect these crashes.


Reproduction Steps

1. Deploy the sample app.
2. Wait 10 minutes for bosh heath check to kick in.
3. The app failed bosh health check with line: \[HEALTH/0] ERR Failed to make HTTP request to '/' on port 8080: timed out after 1.00 seconds. The app container can be seen to be stopped, destroyed, created and started anew. The Apps Manager UI and CF CLI both report two events "App process crashed" and "App crashed".
4. Wait another 10 minutes.
5. Once again in the app logs we can see that the app failed bosh health check and the app container can be seen to be stopped, destroyed, created and started anew. This time there is no new events being recorded.
