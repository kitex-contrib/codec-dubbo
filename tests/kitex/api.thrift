namespace go echo

struct EchoRequest {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoRequest")

struct EchoResponse {
    1: required i32 int32,
}(JavaClassName="kitex.echo.EchoResponse")

struct EchoMultiBoolResponse {
    1: required bool baseResp,
    2: required list<bool> listResp,
    3: required map<bool, bool> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiBoolResponse")

struct EchoMultiByteResponse {
    1: required byte baseResp,
    2: required list<byte> listResp,
    3: required map<byte, byte> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiByteResponse")

struct EchoMultiInt16Response {
    1: required i16 baseResp,
    2: required list<i16> listResp,
    3: required map<i16, i16> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiInt16Response")

struct EchoMultiInt32Response {
    1: required i32 baseResp,
    2: required list<i32> listResp,
    3: required map<i32, i32> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiInt32Response")

struct EchoMultiInt64Response {
    1: required i64 baseResp,
    2: required list<i64> listResp,
    3: required map<i64, i64> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiInt64Response")

struct EchoMultiDoubleResponse {
    1: required double baseResp,
    2: required list<double> listResp,
    3: required map<double, double> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiDoubleResponse")

struct EchoMultiStringResponse {
    1: required string baseResp,
    2: required list<string> listResp,
    3: required map<string, string> mapResp,
}(JavaClassName="org.apache.dubbo.tests.api.EchoMultiStringResponse")

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
    list<bool> EchoBoolList(1: list<bool> req)
    list<byte> EchoByteList(1: list<byte> req)
    list<i16> EchoInt16List(1: list<i16> req)
    list<i32> EchoInt32List(1: list<i32> req)
    list<i64> EchoInt64List(1: list<i64> req)
    list<double> EchoDoubleList(1: list<double> req)
    list<string> EchoStringList(1: list<string> req)
    list<binary> EchoBinaryList(1: list<binary> req)

    // container map
    map<bool, bool> EchoBool2BoolMap(1: map<bool, bool> req)
    map<bool, byte> EchoBool2ByteMap(1: map<bool, byte> req)
    map<bool, i16> EchoBool2Int16Map(1: map<bool, i16> req)
    map<bool, i32> EchoBool2Int32Map(1: map<bool, i32> req)
    map<bool, i64> EchoBool2Int64Map(1: map<bool, i64> req)
    map<bool, double> EchoBool2DoubleMap(1: map<bool, double> req)
    map<bool, string> EchoBool2StringMap(1: map<bool, string> req)
    map<bool, binary> EchoBool2BinaryMap(1: map<bool, binary> req)
    map<bool, list<bool>> EchoBool2BoolListMap(1: map<bool, list<bool>> req)
    map<bool, list<byte>> EchoBool2ByteListMap(1: map<bool, list<byte>> req)
    map<bool, list<i16>> EchoBool2Int16ListMap(1: map<bool, list<i16>> req)
    map<bool, list<i32>> EchoBool2Int32ListMap(1: map<bool, list<i32>> req)
    map<bool, list<i64>> EchoBool2Int64ListMap(1: map<bool, list<i64>> req)
    map<bool, list<double>> EchoBool2DoubleListMap(1: map<bool, list<double>> req)
    map<bool, list<string>> EchoBool2StringListMap(1: map<bool, list<string>> req)
    map<bool, list<binary>> EchoBool2BinaryListMap(1: map<bool, list<binary>> req)

    EchoMultiBoolResponse EchoMultiBool(1: bool baseReq, 2: list<bool> listReq, 3: map<bool, bool> mapReq)
    EchoMultiByteResponse EchoMultiByte(1: byte baseReq, 2: list<byte> listReq, 3: map<byte, byte> mapReq)
    EchoMultiInt16Response EchoMultiInt16(1: i16 baseReq, 2: list<i16> listReq, 3: map<i16, i16> mapReq)
    EchoMultiInt32Response EchoMultiInt32(1: i32 baseReq, 2: list<i32> listReq, 3: map<i32, i32> mapReq)
    EchoMultiInt64Response EchoMultiInt64(1: i64 baseReq, 2: list<i64> listReq, 3: map<i64, i64> mapReq)
    EchoMultiDoubleResponse EchoMultiDouble(1: double baseReq, 2: list<double> listReq, 3: map<double, double> mapReq)
    EchoMultiStringResponse EchoMultiString(1: string baseReq, 2: list<string> listReq, 3: map<string, string> mapReq)

    // method annotation
    bool EchoBaseBool(1: bool req) (hessian.argsType="boolean")
    byte EchoBaseByte(1: byte req) (hessian.argsType="byte")
    i16 EchoBaseInt16(1: i16 req) (hessian.argsType="short")
    i32 EchoBaseInt32(1: i32 req) (hessian.argsType="int")
    i64 EchoBaseInt64(1: i64 req) (hessian.argsType="long")
    double EchoBaseDouble(1: double req) (hessian.argsType="double")
    list<bool> EchoBaseBoolList(1: list<bool> req) (hessian.argsType="boolean[]")
    list<byte> EchoBaseByteList(1: list<byte> req) (hessian.argsType="byte[]")
    list<i16> EchoBaseInt16List(1: list<i16> req) (hessian.argsType="short[]")
    list<i32> EchoBaseInt32List(1: list<i32> req) (hessian.argsType="int[]")
    list<i64> EchoBaseInt64List(1: list<i64> req) (hessian.argsType="long[]")
    list<double> EchoBaseDoubleList(1: list<double> req) (hessian.argsType="double[]")
    map<bool, bool> EchoBool2BoolBaseMap(1: map<bool, bool> req) (hessian.argsType="java.util.HashMap")
    map<bool, byte> EchoBool2ByteBaseMap(1: map<bool, byte> req) (hessian.argsType="java.util.HashMap")
    map<bool, i16> EchoBool2Int16BaseMap(1: map<bool, i16> req) (hessian.argsType="java.util.HashMap")
    map<bool, i32> EchoBool2Int32BaseMap(1: map<bool, i32> req) (hessian.argsType="java.util.HashMap")
    map<bool, i64> EchoBool2Int64BaseMap(1: map<bool, i64> req) (hessian.argsType="java.util.HashMap")
    map<bool, double> EchoBool2DoubleBaseMap(1: map<bool, double> req) (hessian.argsType="java.util.HashMap")
    EchoMultiBoolResponse EchoMultiBaseBool(1: bool baseReq, 2: list<bool> listReq, 3: map<bool, bool> mapReq) (hessian.argsType="boolean,boolean[],java.util.HashMap")
    EchoMultiByteResponse EchoMultiBaseByte(1: byte baseReq, 2: list<byte> listReq, 3: map<byte, byte> mapReq) (hessian.argsType="byte,byte[],java.util.HashMap")
    EchoMultiInt16Response EchoMultiBaseInt16(1: i16 baseReq, 2: list<i16> listReq, 3: map<i16, i16> mapReq) (hessian.argsType="short,short[],java.util.HashMap")
    EchoMultiInt32Response EchoMultiBaseInt32(1: i32 baseReq, 2: list<i32> listReq, 3: map<i32, i32> mapReq) (hessian.argsType="int,int[],java.util.HashMap")
    EchoMultiInt64Response EchoMultiBaseInt64(1: i64 baseReq, 2: list<i64> listReq, 3: map<i64, i64> mapReq) (hessian.argsType="long,long[],java.util.HashMap")
    EchoMultiDoubleResponse EchoMultiBaseDouble(1: double baseReq, 2: list<double> listReq, 3: map<double, double> mapReq) (hessian.argsType="double,double[],java.util.HashMap")
}(JavaClassName="org.apache.dubbo.tests.api.UserProvider")
