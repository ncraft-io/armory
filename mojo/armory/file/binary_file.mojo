
/// 二进制文件，包括文件的元信息及内容
type BinaryFile {
    name: String @1 //< 文件名
    mine_type: String @2 //< 文件的mine-type
    url: Url @3 //< 文件的url
    size: Int64 @4 //< 文件大小，单位：Byte

    content: Bytes @20 //< 文件内容

    create_time: Timestamp @100 //< 文件创建时间
}
