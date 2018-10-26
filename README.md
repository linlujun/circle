# circle
## go restful api server

### POST    "/user"                               CreateUser
  
### POST    "/comment/topic/:topicid"             AddComment
  
### POST    "/user/:username"                     Login
  
### POST    "/info/user/:username"                SetUserInfo
  
### POST    "/info/password/user/:username"       SetUserPwd
  

### POST    "/circle/user/:username"              JoninACircle

### DELETE  "/user/:username"                    DelUser
  
### GET     "/user/:username"                       GetUserInfo

### POST    "/topic"                               AddNewTopic
  
### GET     "/topic/:topicid"                       GetTopicInfo
    
### PUT     "/topic/:topicid"                       ModifyTopic
 
### DELETE  "/topic/:topicid/:circleid"         DelTopic

### GET     "/topics/circle/:cid"                   ListCircleTopics
  
### GET     "/topics/user/:uname"                   ListUserTopics
 
### GET     "/circles"                              ListCircles
  
### GET     "/comments/:tid"                        ListComments
	
### POST    "/upload/pics"                        UploadHandler

### POST    "/admin/circle"                       CreateCircle
  
### DELETE  "/admin/circle/:cid"                DelCircle
  
### PUT     "/admin/circle/:cid"                    SetCircleDesc
