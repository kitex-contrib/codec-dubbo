namespace go echo

struct EchoRequest {
    1: required i32 int32,
}

struct EchoResponse {
    1: required i32 int32,
}

service TestService {
    i32 EchoInt(1: i32 req)
    EchoResponse Echo(1: EchoRequest req)
}
