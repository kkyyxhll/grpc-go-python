# gRPC学习

`https://grpc.org.cn/docs/what-is-grpc/core-concepts/`  
gRPC基于定义服务的理念，指定可以远程调用的方法以及参数和返回类型(方法名,参数列表,返回值)  
gRPC使用Protocol Bufferrs作为接口定义语言(IDL)来描述服务接口和消息结构。  

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}  

## 四种服务方法

### 单向RPC

客户端向服务器发送单个请求并获得单个响应，如普通的函数调用一样。  
例如:  
rpc SayHello(HelloRequest) returns (HelloResponse);  

### 服务器流式RPC  

客户端向服务器发送请求并获得一个流来读取一系列响应消息。客户端从返回  
的流中读取，直到没有更多消息。gRPC 保证单个RPC调用内的消息顺序。  
例如:  
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);  

### 客户端流式RPC

客户端写入一系列消息并将其发送到服务器，同样使用提供的流。客户端完成写入消息后，  
等待服务器读取消息并返回其响应。同样，gRPC 保证单个 RPC 调用内的消息顺序。  
例如:  
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);  

### 双向流式RPC

双方使用读写流发送一系列消息。这两个流独立运行，因此客户端和服务器可以按照  
任何顺序读取和写入：例如，服务器可以在写入响应之前等待接收所有客户端消息，  
或者它可以交替读取消息然后写入消息，或者读取和写入的其他组合。每个流中的消息顺序保持不变。
例如:
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);  
