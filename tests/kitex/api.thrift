namespace go echo

struct EchoRequest {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoRequest")

struct EchoResponse {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoResponse")

service TestService {
    i32 EchoInt(1: i32 req)

    // base types
    bool EchoBool(1: bool req)
    byte EchoByte(1: byte req)
    i16 EchoInt16(1: i16 req)
    i32 EchoInt32(1: i32 req)
    i64 EchoInt64(1: i64 req)
    double EchoDouble(1: double req)
    string EchoString(1: string req)

    // special types
    binary EchoBinary(1: binary req)

    // structs
    EchoResponse Echo(1: EchoRequest req)

    // container list
    list<bool> EchoBoolList(1: list<bool>req)
    list<byte> EchoByteList(1: list<byte>req)
    list<i16> EchoInt16List(1: list<i16>req)
    list<i32> EchoInt32List(1: list<i32>req)
    list<i64> EchoInt64List(1: list<i64>req)
    list<double> EchoDoubleList(1: list<double>req)
    list<string> EchoStringList(1: list<string>req)
    list<binary> EchoBinaryList(1: list<binary>req)
}
