namespace go echo

struct EchoRequest {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoRequest")

struct EchoResponse {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoResponse")

service TestService {
    i32 EchoInt(1: i32 req)
    byte EchoByte(1: byte req)
    EchoResponse Echo(1: EchoRequest req)
}
