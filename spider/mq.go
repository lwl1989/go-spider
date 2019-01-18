package spider

//docker run -d --hostname my-rabbit --name rmq -p 15672:15672 -p 5672:5672 -p 25672:25672 -e RABBITMQ_DEFAULT_USER=用户名 -e RABBITMQ_DEFAULT_PASS=密码 rabbitmq:3-management

//do not use mq in this project
type messageQueue struct {
	Type string `json:"type"`
} 
