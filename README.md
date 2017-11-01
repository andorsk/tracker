# Live Tracker 

Live tracker is a web application that will publically transmit your location to a web app. The initial use case is for henosisknot.com where I will let people follow me as I travel around the world. 

Dependencies: 
Go
Protobuf3.0

# TODO: 
Currently setting up the API to push a heartbeat to a server. Need to complete this first.
After, I need to set up a secure stream beween the endpoint and the android device that will be sending the location. 
This will require me to also build an android app which sends my location to an endpoint every 5 minutes. 

