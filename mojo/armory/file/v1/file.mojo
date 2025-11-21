/// 文件服务
interface File {
    /// 查询服务范围内的所有数据库及表单
    @http.get("/armory/file/v1/files/{name}")
    get_file(name: String @1) //< 文件名称
            -> BinaryFile

    /// 上传文件，采用http form的形式进行上传，返回文件元信息
    @http.post("/armory/file/v1/files")
    create_file(file: BinaryFile @1) //< 需要上传的文件
            -> BinaryFile

    /// 批量上传文件，采用http form的形式进行上传，返回文件元信息
    @http.post("/armory/file/v1/files:batch")
    batch_create_file(files: [BinaryFile] @1) //< 需要上传的文件集
            -> [BinaryFile]
}
