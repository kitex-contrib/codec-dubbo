namespace go hello

struct GreetRequest {
    1: required string req,
}(JavaClassName="org.cloudwego.kitex.samples.api.GreetRequest")

struct GreetResponse {
    1: required string resp,
}(JavaClassName="org.cloudwego.kitex.samples.api.GreetResponse")

service GreetService {
    string Greet(1: string req)
    GreetResponse GreetWithStruct(1: GreetRequest req)
}
