#[allow(dead_code)]

extern "C" {
    fn new_client() -> i32;
    fn new_client_with_port(port: i32) -> i32;
    fn invoke_method_with_content(
        app_id: *const u8,
        app_id_len: i32,
        method: *const u8,
        method_len: i32,
        verb: *const u8,
        verb_len: i32,
        content_type: *const u8,
        content_type_len: i32,
        content: *const u8,
        content_len: i32,
    ) -> i32;
    fn write_mem(pointer: *const u8);
}

#[no_mangle]
pub unsafe extern "C" fn run_dapr() -> i32 {
    let _res = new_client();

    let app_id = "dapr-demo".to_string();
    let api_ptr = app_id.as_bytes().as_ptr();
    let api_len = app_id.len() as i32;

    let method = "/api/hello".to_string();
    let method_ptr = method.as_bytes().as_ptr();
    let method_len = method.len() as i32;

    let verb = "post".to_string();
    let verb_ptr = verb.as_bytes().as_ptr();
    let verb_len = verb.len() as i32;

    let content_ty = "text".to_string();
    let content_ty_ptr = content_ty.as_bytes().as_ptr();
    let content_ty_len = content_ty.len() as i32;

    let content = "hello".to_string();
    let content_ptr = content.as_bytes().as_ptr();
    let content_len = content.len() as i32;

    let res_len = invoke_method_with_content(
        api_ptr,
        api_len,
        method_ptr,
        method_len,
        verb_ptr,
        verb_len,
        content_ty_ptr,
        content_ty_len,
        content_ptr,
        content_len,
    );

    // malloc memory
    let mut buffer = Vec::with_capacity(res_len as usize);
    let pointer = buffer.as_mut_ptr();

    // call host function to write source code to the memory
    write_mem(pointer);

    // find occurrences from source code
    buffer.set_len(res_len as usize);
    let content = String::from_utf8(buffer).unwrap();
    println!("{}", content);
    return content.len() as i32;
}
