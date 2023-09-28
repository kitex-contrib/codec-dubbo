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

    // container map
    map<bool, bool> EchoBool2BoolMap(1: map<bool, bool>req)
    map<bool, byte> EchoBool2ByteMap(1: map<bool, byte>req)
    map<bool, i16> EchoBool2Int16Map(1: map<bool, i16>req)
    map<bool, i32> EchoBool2Int32Map(1: map<bool, i32>req)
    map<bool, i64> EchoBool2Int64Map(1: map<bool, i64>req)
    map<bool, double> EchoBool2DoubleMap(1: map<bool, double>req)
    map<bool, string> EchoBool2StringMap(1: map<bool, string>req)
    map<bool, binary> EchoBool2BinaryMap(1: map<bool, binary>req)
    map<bool, list<bool>> EchoBool2BoolListMap(1: map<bool, list<bool>>req)
    map<bool, list<byte>> EchoBool2ByteListMap(1: map<bool, list<byte>>req)
    map<bool, list<i16>> EchoBool2Int16ListMap(1: map<bool, list<i16>>req)
    map<bool, list<i32>> EchoBool2Int32ListMap(1: map<bool, list<i32>>req)
    map<bool, list<i64>> EchoBool2Int64ListMap(1: map<bool, list<i64>>req)
    map<bool, list<double>> EchoBool2DoubleListMap(1: map<bool, list<double>>req)
    map<bool, list<string>> EchoBool2StringListMap(1: map<bool, list<string>>req)
    map<bool, list<binary>> EchoBool2BinaryListMap(1: map<bool, list<binary>>req)
}(JavaClassName="org.apache.dubbo.tests.api.UserProvider")
