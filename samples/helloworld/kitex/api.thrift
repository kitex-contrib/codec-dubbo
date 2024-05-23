namespace go hello

enum KitexEnum {
    ONE,
    TWO,
    THREE,
    FOUR,
    FIVE,
}(JavaClassName="org.cloudwego.kitex.samples.enumeration.KitexEnum")

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

service GreetEnumService {
    KitexEnum GreetEnum(1: KitexEnum req)
}