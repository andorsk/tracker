#Model Folder

Contains not just all the structures but the interfaces for DB query. The controller willhandle the requests and then hit the model interface, which is responsible for actually picking up the information from the db. 


Basic Structure:


User

//Weakness of implemnetation: If there are too many points in the track then it will become an issue
HeartBeatTracks
    -> HeartBeats
    -> Users
    

